package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/data/repo"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/errors"
	"mvrp/util"
	"time"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST INVENTORY
type ListInventoryRequest struct {
	Ctx      context.Context
	RepoTx   *repo.RepoTx
	ItemType *string
}

func (s *InventoryService) NewListInventoryRequest(ctx context.Context) *ListInventoryRequest {
	return &ListInventoryRequest{
		Ctx: ctx,
	}
}

func (s *InventoryService) NewListInventoryByItemTypeRequest(ctx context.Context, itemType string) *ListInventoryRequest {
	return &ListInventoryRequest{
		Ctx:      ctx,
		ItemType: &itemType,
	}
}

type ListInventoryResponse struct {
	Payload inventory.InventorySlice `json:"payload"`
}

func (s *InventoryService) NewListInventoryResponse(payload inventory.InventorySlice) *ListInventoryResponse {
	return &ListInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListInventory(req *ListInventoryRequest) (*ListInventoryResponse, error) {
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

	res, err := s.Repo.Inventory.ListAllInventories(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListInventoryResponse{
		Payload: res,
	}
	return &resp, nil
}

func (s *InventoryService) ListInventoryByItemType(req *ListInventoryRequest) (*ListInventoryResponse, error) {
	if req.ItemType == nil {
		return nil, errors.WrapError(errors.ErrTypeMissingField, "ItemType is required")
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

	res, err := s.Repo.Inventory.ListInventoriesByItemType(req.Ctx, tx, *req.ItemType)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListInventoryResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH INVENTORY
type SearchInventoryRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchInventoryDTO
}

func (s *InventoryService) NewSearchInventoryRequest(ctx context.Context, payload dto.SearchInventoryDTO) *SearchInventoryRequest {
	return &SearchInventoryRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchInventoryResponse struct {
	Payload    inventory.InventorySlice `json:"payload"`
	Pagination dto.PaginationDTO        `json:"pagination"`
}

func (s *InventoryService) NewSearchInventoryResponse(payload inventory.InventorySlice) *SearchInventoryResponse {
	return &SearchInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchInventory(req *SearchInventoryRequest) (*SearchInventoryResponse, error) {
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

	res, totalCount, err := s.Repo.Inventory.SearchInventories(req.Ctx, tx, req.Payload)
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
	resp := SearchInventoryResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET INVENTORY
type GetInventoryRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *InventoryService) NewGetInventoryRequest(ctx context.Context, id int) *GetInventoryRequest {
	return &GetInventoryRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetInventoryResponse struct {
	Payload inventory.Inventory `json:"payload"`
}

func (s *InventoryService) NewGetInventoryResponse(payload inventory.Inventory) *GetInventoryResponse {
	return &GetInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetInventory(req *GetInventoryRequest) (*GetInventoryResponse, error) {
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

	res, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetInventoryResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE INVENTORY
type CreateInventoryRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.CreateInventoryDTO
}

func (s *InventoryService) NewCreateInventoryRequest(ctx context.Context, payload dto.CreateInventoryDTO) *CreateInventoryRequest {
	return &CreateInventoryRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateInventoryResponse struct {
	Payload inventory.Inventory `json:"payload"`
}

func (s *InventoryService) NewCreateInventoryResponse(payload inventory.Inventory) *CreateInventoryResponse {
	return &CreateInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) CreateInventory(req *CreateInventoryRequest) (*CreateInventoryResponse, error) {
	/*
		1. Create Inventory
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

	// process inventory
	err = proc.ProcessInventoryAmounts(&req.Payload.Inventory)
	if err != nil {
		return nil, err
	}

	// create inventory
	nextID, err := s.Repo.Inventory.GetNextEntryInventoryID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.Inventory.ID = nextID
	if req.Payload.Inventory.InventoryNumber == "" {
		req.Payload.Inventory.InventoryNumber = util.Util.Str.ToString(nextID)
	}
	// process generated amounts
	err = proc.ProcessInventoryAmounts(&req.Payload.Inventory)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.CreateInventory(req.Ctx, tx, &req.Payload.Inventory)
	if err != nil {
		return nil, err
	}

	// create inventory transaction if quantity is not zero
	if req.Payload.Inventory.QuantityAvailable.Big.Cmp(decimal.New(0, 2)) != 0 {
		nextID, err := s.Repo.Inventory.GetNextEntryInventoryTransactionID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		invTx := &inventory.InventoryTransaction{
			ID:              nextID,
			InventoryID:     null.IntFrom(req.Payload.Inventory.ID),
			TransactionType: inventory.InventoryTransactionTypeInitialStock,
			Quantity:        types.NewDecimal(req.Payload.Inventory.QuantityAvailable.Big),
			Reason:          null.StringFrom("Initial Stock"),
			TransactionDate: null.TimeFrom(time.Now().UTC()),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}
	}

	// get created inventory
	inventory, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := CreateInventoryResponse{
		Payload: *inventory,
	}

	return &resp, nil
}

// UPDATE INVENTORY
type UpdateInventoryRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.UpdateInventoryDTO
}

func (s *InventoryService) NewUpdateInventoryRequest(ctx context.Context, payload dto.UpdateInventoryDTO) *UpdateInventoryRequest {
	return &UpdateInventoryRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateInventoryResponse struct {
	Payload inventory.Inventory `json:"payload"`
}

func (s *InventoryService) NewUpdateInventoryResponse(payload inventory.Inventory) *UpdateInventoryResponse {
	return &UpdateInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) UpdateInventory(req *UpdateInventoryRequest) (*UpdateInventoryResponse, error) {
	/*
		1. Update Inventory
		2. Create Inventory Transaction
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

	// get current inventory amount
	currInventory, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}
	amountOffset := types.NewNullDecimal(decimal.New(0, 2))
	amountOffset.Sub(req.Payload.Inventory.QuantityAvailable.Big, currInventory.QuantityAvailable.Big)

	// process generated amounts
	err = proc.ProcessInventoryAmounts(&req.Payload.Inventory)
	if err != nil {
		return nil, err
	}
	// update inventory
	err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, &req.Payload.Inventory)
	if err != nil {
		return nil, err
	}

	// create inventory transaction
	quantityChanged := amountOffset.Big.Cmp(decimal.New(0, 2)) != 0
	if quantityChanged {
		nextID, err := s.Repo.Inventory.GetNextEntryInventoryTransactionID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		invTx := &inventory.InventoryTransaction{
			ID:              nextID,
			InventoryID:     null.IntFrom(req.Payload.Inventory.ID),
			TransactionType: inventory.InventoryTransactionTypeGeneralAdjustment,
			Quantity:        types.NewDecimal(amountOffset.Big),
			Reason:          null.StringFrom("General Adjustment"),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}
	}

	// get updated inventory
	inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := UpdateInventoryResponse{
		Payload: *inv,
	}

	return &resp, nil
}

// DELETE INVENTORY
type DeleteInventoryRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *InventoryService) NewDeleteInventoryRequest(ctx context.Context, id int) *DeleteInventoryRequest {
	return &DeleteInventoryRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteInventoryResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewDeleteInventoryResponse(payload bool) *DeleteInventoryResponse {
	return &DeleteInventoryResponse{
		Payload: payload,
	}
}

func (s *InventoryService) DeleteInventory(req *DeleteInventoryRequest) (*DeleteInventoryResponse, error) {
	/*
		1. Delete Inventory
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

	// get inventory
	inventory, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete sales order
	err = s.Repo.Inventory.DeleteInventory(req.Ctx, tx, inventory)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := DeleteInventoryResponse{
		Payload: true,
	}

	return &resp, nil
}

// GET INVENTORY EXISTS BY ITEM ID
type GetInventoryExistsByItemIDRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ItemID int
}

func (s *InventoryService) NewGetInventoryExistsByItemIDRequest(ctx context.Context, itemID int) *GetInventoryExistsByItemIDRequest {
	return &GetInventoryExistsByItemIDRequest{
		Ctx:    ctx,
		ItemID: itemID,
	}
}

type GetInventoryExistsByItemIDResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewGetInventoryExistsByItemIDResponse(payload bool) *GetInventoryExistsByItemIDResponse {
	return &GetInventoryExistsByItemIDResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetInventoryExistsByItemID(req *GetInventoryExistsByItemIDRequest) (*GetInventoryExistsByItemIDResponse, error) {
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

	exists, err := s.Repo.Inventory.GetInventoryExistsByItemID(req.Ctx, tx, req.ItemID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetInventoryExistsByItemIDResponse{
		Payload: exists,
	}
	return &resp, nil
}
