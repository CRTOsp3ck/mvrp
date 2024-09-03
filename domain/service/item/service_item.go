package item

import (
	"context"
	"mvrp/data/model/item"
	"mvrp/domain/dto"
)

// LIST ITEM
type ListItemRequest struct {
	Ctx context.Context
}

func NewListItemRequest(ctx context.Context) *ListItemRequest {
	return &ListItemRequest{Ctx: ctx}
}

type ListItemResponse struct {
	Payload item.ItemSlice
}

func NewListItemResponse(payload item.ItemSlice) *ListItemResponse {
	return &ListItemResponse{Payload: payload}
}

func (s *ItemService) ListItem(req *ListItemRequest) (*ListItemResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Item.ListAllItems(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListItemResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH ITEM
type SearchItemRequest struct {
	Ctx     context.Context
	Payload dto.SearchItemDTO
}

func NewSearchItemRequest(ctx context.Context, payload dto.SearchItemDTO) *SearchItemRequest {
	return &SearchItemRequest{Ctx: ctx, Payload: payload}
}

type SearchItemResponse struct {
	Payload item.ItemSlice
}

func NewSearchItemResponse(payload item.ItemSlice) *SearchItemResponse {
	return &SearchItemResponse{Payload: payload}
}
func (s *ItemService) SearchItem(req *SearchItemRequest) (*SearchItemResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Item.SearchItems(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchItemResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET ITEM
type GetItemRequest struct {
	Ctx context.Context
	ID  int
}

func NewGetItemRequest(ctx context.Context, id int) *GetItemRequest {
	return &GetItemRequest{Ctx: ctx, ID: id}
}

type GetItemResponse struct {
	Payload item.Item
}

func NewGetItemResponse(payload item.Item) *GetItemResponse {
	return &GetItemResponse{Payload: payload}
}

func (s *ItemService) GetItem(req *GetItemRequest) (*GetItemResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := GetItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE ITEM
type CreateItemRequest struct {
	Ctx     context.Context
	Payload dto.CreateItemDTO
}

func NewCreateItemRequest(ctx context.Context, payload dto.CreateItemDTO) *CreateItemRequest {
	return &CreateItemRequest{Ctx: ctx, Payload: payload}
}

type CreateItemResponse struct {
	Payload item.Item
}

func NewCreateItemResponse(payload item.Item) *CreateItemResponse {
	return &CreateItemResponse{Payload: payload}
}

func (s *ItemService) CreateItem(req *CreateItemRequest) (*CreateItemResponse, error) {
	/*
		1. Create Item
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create item
	err = s.Repo.Item.CreateItem(req.Ctx, tx, &req.Payload.Item)
	if err != nil {
		return nil, err
	}

	// get item
	res, err := s.Repo.Item.GetItemByID(req.Ctx, tx, req.Payload.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE ITEM
type UpdateItemRequest struct {
	Ctx     context.Context
	Payload dto.UpdateItemDTO
}

func NewUpdateItemRequest(ctx context.Context, payload dto.UpdateItemDTO) *UpdateItemRequest {
	return &UpdateItemRequest{Ctx: ctx, Payload: payload}
}

type UpdateItemResponse struct {
	Payload item.Item
}

func NewUpdateItemResponse(payload item.Item) *UpdateItemResponse {
	return &UpdateItemResponse{Payload: payload}
}

func (s *ItemService) UpdateItem(req *UpdateItemRequest) (*UpdateItemResponse, error) {
	/*
		1. Update Item
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateItemResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE ITEM
type DeleteItemRequest struct {
	Ctx context.Context
	ID  int
}

func NewDeleteItemRequest(ctx context.Context, id int) *DeleteItemRequest {
	return &DeleteItemRequest{Ctx: ctx, ID: id}
}

type DeleteItemResponse struct {
	Payload bool
}

func NewDeleteItemResponse(payload bool) *DeleteItemResponse {
	return &DeleteItemResponse{Payload: payload}
}

func (s *ItemService) DeleteItem(req *DeleteItemRequest) (*DeleteItemResponse, error) {
	/*
		1. Delete Item
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteItemResponse{
		Payload: true,
	}
	return &resp, nil
}
