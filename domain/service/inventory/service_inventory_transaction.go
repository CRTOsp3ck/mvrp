package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
)

// LIST INVENTORY TRANSACTION
type ListInventoryTransactionRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListInventoryTransactionRequest(ctx context.Context) *ListInventoryTransactionRequest {
	return &ListInventoryTransactionRequest{
		Ctx: ctx,
	}
}

type ListInventoryTransactionResponse struct {
	Payload inventory.InventoryTransactionSlice `json:"payload"`
}

func (s *InventoryService) NewListInventoryTransactionResponse(payload inventory.InventoryTransactionSlice) *ListInventoryTransactionResponse {
	return &ListInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListInventoryTransaction(req *ListInventoryTransactionRequest) (*ListInventoryTransactionResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllInventoryTransactions(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListInventoryTransactionResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH INVENTORY TRANSACTION
type SearchInventoryTransactionRequest struct {
	Ctx     context.Context
	Payload dto.SearchInventoryTransactionDTO
}

func (s *InventoryService) NewSearchInventoryTransactionRequest(ctx context.Context, payload dto.SearchInventoryTransactionDTO) *SearchInventoryTransactionRequest {
	return &SearchInventoryTransactionRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchInventoryTransactionResponse struct {
	Payload    inventory.InventoryTransactionSlice `json:"payload"`
	Pagination dto.PaginationDTO                   `json:"pagination"`
}

func (s *InventoryService) NewSearchInventoryTransactionResponse(payload inventory.InventoryTransactionSlice) *SearchInventoryTransactionResponse {
	return &SearchInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchInventoryTransaction(req *SearchInventoryTransactionRequest) (*SearchInventoryTransactionResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchInventoryTransactions(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Inventory.GetInventoryTransactionTotalCountByInventoryID(req.Ctx, tx, req.Payload.InventoryId)
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
	resp := SearchInventoryTransactionResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET INVENTORY TRANSACTION
type GetInventoryTransactionRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetInventoryTransactionRequest(ctx context.Context, id int) *GetInventoryTransactionRequest {
	return &GetInventoryTransactionRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetInventoryTransactionResponse struct {
	Payload inventory.InventoryTransaction `json:"payload"`
}

func (s *InventoryService) NewGetInventoryTransactionResponse(payload inventory.InventoryTransaction) *GetInventoryTransactionResponse {
	return &GetInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetInventoryTransaction(req *GetInventoryTransactionRequest) (*GetInventoryTransactionResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetInventoryTransactionByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetInventoryTransactionResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE INVENTORY TRANSACTION
type CreateInventoryTransactionRequest struct {
	Ctx     context.Context
	Payload dto.CreateInventoryTransactionDTO
}

func (s *InventoryService) NewCreateInventoryTransactionRequest(ctx context.Context, payload dto.CreateInventoryTransactionDTO) *CreateInventoryTransactionRequest {
	return &CreateInventoryTransactionRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateInventoryTransactionResponse struct {
	Payload inventory.InventoryTransaction `json:"payload"`
}

func (s *InventoryService) NewCreateInventoryTransactionResponse(payload inventory.InventoryTransaction) *CreateInventoryTransactionResponse {
	return &CreateInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) CreateInventoryTransaction(req *CreateInventoryTransactionRequest) (*CreateInventoryTransactionResponse, error) {
	/*
		1. Create Inventory
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create inventory
	nextID, err := s.Repo.Inventory.GetNextEntryInventoryTransactionID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.InventoryTransaction.ID = nextID
	err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, &req.Payload.InventoryTransaction)
	if err != nil {
		return nil, err
	}

	// get created inventory
	invTx, err := s.Repo.Inventory.GetInventoryTransactionByID(req.Ctx, tx, req.Payload.InventoryTransaction.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateInventoryTransactionResponse{
		Payload: *invTx,
	}

	return &resp, nil
}

// UPDATE INVENTORY TRANSACTION
type UpdateInventoryTransactionRequest struct {
	Ctx     context.Context
	Payload dto.UpdateInventoryTransactionDTO
}

func (s *InventoryService) NewUpdateInventoryTransactionRequest(ctx context.Context, payload dto.UpdateInventoryTransactionDTO) *UpdateInventoryTransactionRequest {
	return &UpdateInventoryTransactionRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateInventoryTransactionResponse struct {
	Payload inventory.InventoryTransaction `json:"payload"`
}

func (s *InventoryService) NewUpdateInventoryTransactionResponse(payload inventory.InventoryTransaction) *UpdateInventoryTransactionResponse {
	return &UpdateInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) UpdateInventoryTransaction(req *UpdateInventoryTransactionRequest) (*UpdateInventoryTransactionResponse, error) {
	/*
		1. Update Inventory
		2. Create Inventory Transaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// update inventory
	err = s.Repo.Inventory.UpdateInventoryTransaction(req.Ctx, tx, &req.Payload.InventoryTransaction)
	if err != nil {
		return nil, err
	}

	// get updated inventory
	invTx, err := s.Repo.Inventory.GetInventoryTransactionByID(req.Ctx, tx, req.Payload.InventoryTransaction.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateInventoryTransactionResponse{
		Payload: *invTx,
	}

	return &resp, nil
}

// DELETE INVENTORY TRANSACTION
type DeleteInventoryTransactionRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewDeleteInventoryTransactionRequest(ctx context.Context, id int) *DeleteInventoryTransactionRequest {
	return &DeleteInventoryTransactionRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteInventoryTransactionResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewDeleteInventoryTransactionResponse(payload bool) *DeleteInventoryTransactionResponse {
	return &DeleteInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) DeleteInventoryTransaction(req *DeleteInventoryTransactionRequest) (*DeleteInventoryTransactionResponse, error) {
	/*
		1. Delete Inventory
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get inventory
	invTx, err := s.Repo.Inventory.GetInventoryTransactionByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete sales order
	err = s.Repo.Inventory.DeleteInventoryTransaction(req.Ctx, tx, invTx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteInventoryTransactionResponse{
		Payload: true,
	}

	return &resp, nil
}

// SEARCH ALL INVENTORY TRANSACTIONS
type SearchAllInventoryTransactionRequest struct {
	Ctx     context.Context
	Payload dto.SearchInventoryTransactionDTO
}

func (s *InventoryService) NewSearchAllInventoryTransactionRequest(ctx context.Context, payload dto.SearchInventoryTransactionDTO) *SearchAllInventoryTransactionRequest {
	return &SearchAllInventoryTransactionRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchAllInventoryTransactionResponse struct {
	Payload    inventory.InventoryTransactionSlice `json:"payload"`
	Pagination dto.PaginationDTO                   `json:"pagination"`
}

func (s *InventoryService) NewSearchAllInventoryTransactionResponse(payload inventory.InventoryTransactionSlice) *SearchAllInventoryTransactionResponse {
	return &SearchAllInventoryTransactionResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchAllInventoryTransaction(req *SearchAllInventoryTransactionRequest) (*SearchAllInventoryTransactionResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchAllInventoryTransactions(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Inventory.GetInventoryTransactionTotalCount(req.Ctx, tx)
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
	resp := SearchAllInventoryTransactionResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}
