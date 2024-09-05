package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/jinzhu/copier"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST INVENTORY
type ListInventoryRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListInventoryRequest(ctx context.Context) *ListInventoryRequest {
	return &ListInventoryRequest{
		Ctx: ctx,
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllInventories(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListInventoryResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH INVENTORY
type SearchInventoryRequest struct {
	Ctx     context.Context
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchInventories(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Inventory.GetInventoryTotalCount(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
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
	Ctx context.Context
	ID  int
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetInventoryResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE INVENTORY
type CreateInventoryRequest struct {
	Ctx     context.Context
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create sales order
	if req.Payload.Inventory.InventoryNumber == "" {
		nextID, err := s.Repo.Inventory.GetNextEntryInventoryID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.Inventory.InventoryNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Inventory.CreateInventory(req.Ctx, tx, &req.Payload.Inventory)
	if err != nil {
		return nil, err
	}

	// get created sales order
	inventory, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateInventoryResponse{
		Payload: *inventory,
	}

	return &resp, nil
}

// UPDATE INVENTORY
type UpdateInventoryRequest struct {
	Ctx     context.Context
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get current inventory amount
	currInventory, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}
	amountOffset := types.NewNullDecimal(decimal.New(0, 2))
	amountOffset.Sub(req.Payload.Inventory.QuantityAvailable.Big, currInventory.QuantityAvailable.Big)

	// preserve other fields (only modify QuantityAvailable)
	var newInventory inventory.Inventory
	copier.Copy(&newInventory, &req.Payload.Inventory)
	newInventory.QuantityAvailable = req.Payload.Inventory.QuantityAvailable

	// update inventory
	err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, &newInventory)
	if err != nil {
		return nil, err
	}

	// create inventory transaction
	invTx := &inventory.InventoryTransaction{
		InventoryID:     null.IntFrom(req.Payload.Inventory.ID),
		TransactionType: inventory.InventoryTransactionTypeGeneralAdjustment,
		Quantity:        types.NewDecimal(amountOffset.Big),
	}
	err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
	if err != nil {
		return nil, err
	}

	// get updated sales order
	salesOrder, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.Inventory.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateInventoryResponse{
		Payload: *salesOrder,
	}

	return &resp, nil
}

// DELETE INVENTORY
type DeleteInventoryRequest struct {
	Ctx context.Context
	ID  int
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteInventoryResponse{
		Payload: true,
	}

	return &resp, nil
}
