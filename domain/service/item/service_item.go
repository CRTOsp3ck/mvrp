package item

import (
	"context"
	"mvrp/data/model/item"
	"mvrp/data/repo"
	"mvrp/domain/dto"
	"mvrp/errors"
)

// LIST ITEM
type ListItemRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	Type   *item.ItemType
}

func (s *ItemService) NewListItemRequest(ctx context.Context) *ListItemRequest {
	return &ListItemRequest{Ctx: ctx}
}

func (s *ItemService) NewListItemByTypeRequest(ctx context.Context, itemType *item.ItemType) *ListItemRequest {
	return &ListItemRequest{Ctx: ctx, Type: itemType}
}

type ListItemResponse struct {
	Payload item.ItemSlice `json:"payload"`
}

func (s *ItemService) NewListItemResponse(payload item.ItemSlice) *ListItemResponse {
	return &ListItemResponse{Payload: payload}
}

func (s *ItemService) ListItem(req *ListItemRequest) (*ListItemResponse, error) {
	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	res, err := s.Repo.Item.ListAllItems(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListItemResponse{
		Payload: res,
	}
	return &resp, nil
}

func (s *ItemService) ListItemByType(req *ListItemRequest) (*ListItemResponse, error) {
	if req.Type == nil {
		return nil, errors.WrapError(errors.ErrTypeMissingField, "Item type is required")
	}

	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	res, err := s.Repo.Item.ListItemsByType(req.Ctx, tx, req.Type.String())
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListItemResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH ITEM
type SearchItemRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchItemDTO
}

func (s *ItemService) NewSearchItemRequest(ctx context.Context, payload dto.SearchItemDTO) *SearchItemRequest {
	return &SearchItemRequest{Ctx: ctx, Payload: payload}
}

type SearchItemResponse struct {
	Payload    item.ItemSlice    `json:"payload"`
	Pagination dto.PaginationDTO `json:"pagination"`
}

func (s *ItemService) NewSearchItemResponse(payload item.ItemSlice) *SearchItemResponse {
	return &SearchItemResponse{Payload: payload}
}

func (s *ItemService) SearchItem(req *SearchItemRequest) (*SearchItemResponse, error) {
	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	res, totalCount, err := s.Repo.Item.SearchItems(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	pd := dto.PaginationDTO{
		TotalItems:   totalCount,
		ItemsPerPage: req.Payload.ItemsPerPage,
		Page:         req.Payload.Page,
		SortBy:       req.Payload.SortBy,
		OrderBy:      req.Payload.OrderBy,
	}
	resp := SearchItemResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET ITEM
type GetItemRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *ItemService) NewGetItemRequest(ctx context.Context, id int) *GetItemRequest {
	return &GetItemRequest{Ctx: ctx, ID: id}
}

type GetItemResponse struct {
	Payload item.Item `json:"payload"`
}

func (s *ItemService) NewGetItemResponse(payload item.Item) *GetItemResponse {
	return &GetItemResponse{Payload: payload}
}

func (s *ItemService) GetItem(req *GetItemRequest) (*GetItemResponse, error) {
	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	res, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE ITEM
type CreateItemRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.CreateItemDTO
}

func (s *ItemService) NewCreateItemRequest(ctx context.Context, payload dto.CreateItemDTO) *CreateItemRequest {
	return &CreateItemRequest{Ctx: ctx, Payload: payload}
}

type CreateItemResponse struct {
	Payload item.Item `json:"payload"`
}

func (s *ItemService) NewCreateItemResponse(payload item.Item) *CreateItemResponse {
	return &CreateItemResponse{Payload: payload}
}

func (s *ItemService) CreateItem(req *CreateItemRequest) (*CreateItemResponse, error) {
	/*
		1. Create Item
	*/

	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	// create item
	nextID, err := s.Repo.Item.GetNextEntryItemID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.Item.ID = nextID
	err = s.Repo.Item.CreateItem(req.Ctx, tx, &req.Payload.Item)
	if err != nil {
		return nil, err
	}

	// get item
	res, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.Payload.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := CreateItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE ITEM
type UpdateItemRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.UpdateItemDTO
}

func (s *ItemService) NewUpdateItemRequest(ctx context.Context, payload dto.UpdateItemDTO) *UpdateItemRequest {
	return &UpdateItemRequest{Ctx: ctx, Payload: payload}
}

type UpdateItemResponse struct {
	Payload item.Item `json:"payload"`
}

func (s *ItemService) NewUpdateItemResponse(payload item.Item) *UpdateItemResponse {
	return &UpdateItemResponse{Payload: payload}
}

func (s *ItemService) UpdateItem(req *UpdateItemRequest) (*UpdateItemResponse, error) {
	/*
		1. Update Item
	*/

	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	// update item
	err = s.Repo.Item.UpdateItem(req.Ctx, tx, &req.Payload.Item)
	if err != nil {
		return nil, err
	}

	// get updated item
	res, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.Payload.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := UpdateItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE ITEM
type DeleteItemRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *ItemService) NewDeleteItemRequest(ctx context.Context, id int) *DeleteItemRequest {
	return &DeleteItemRequest{Ctx: ctx, ID: id}
}

type DeleteItemResponse struct {
	Payload bool `json:"payload"`
}

func (s *ItemService) NewDeleteItemResponse(payload bool) *DeleteItemResponse {
	return &DeleteItemResponse{Payload: payload}
}

func (s *ItemService) DeleteItem(req *DeleteItemRequest) (*DeleteItemResponse, error) {
	/*
		1. Delete Item
	*/

	rtx := req.RepoTx
	var err error
	if rtx == nil {
		rtx, err = s.Repo.BeginRepoTx(req.Ctx)
		if err != nil {
			return nil, err
		}
		defer rtx.Tx.Rollback()
	}
	tx := rtx.Tx

	// get item
	item, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete item
	err = s.Repo.Item.DeleteItem(req.Ctx, tx, item)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := DeleteItemResponse{
		Payload: true,
	}
	return &resp, nil
}
