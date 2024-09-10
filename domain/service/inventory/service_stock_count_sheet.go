package inventory

import (
	"context"
	"errors"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/merge"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST STOCK COUNT SHEET
type ListStockCountSheetRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListStockCountSheetRequest(ctx context.Context) *ListStockCountSheetRequest {
	return &ListStockCountSheetRequest{
		Ctx: ctx,
	}
}

type ListStockCountSheetResponse struct {
	Payload inventory.StockCountSheetSlice `json:"payload"`
}

func (s *InventoryService) NewListStockCountSheetResponse(payload inventory.StockCountSheetSlice) *ListStockCountSheetResponse {
	return &ListStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListStockCountSheet(req *ListStockCountSheetRequest) (*ListStockCountSheetResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllStockCountSheets(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListStockCountSheetResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH STOCK COUNT SHEET
type SearchStockCountSheetRequest struct {
	Ctx     context.Context
	Payload dto.SearchStockCountSheetDTO
}

func (s *InventoryService) NewSearchStockCountSheetRequest(ctx context.Context, payload dto.SearchStockCountSheetDTO) *SearchStockCountSheetRequest {
	return &SearchStockCountSheetRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchStockCountSheetResponse struct {
	Payload    inventory.StockCountSheetSlice `json:"payload"`
	Pagination dto.PaginationDTO              `json:"pagination"`
}

func (s *InventoryService) NewSearchStockCountSheetResponse(payload inventory.StockCountSheetSlice) *SearchStockCountSheetResponse {
	return &SearchStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchStockCountSheet(req *SearchStockCountSheetRequest) (*SearchStockCountSheetResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchStockCountSheets(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Inventory.GetStockCountSheetTotalCount(req.Ctx, tx)
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
	resp := SearchStockCountSheetResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET STOCK COUNT SHEET
type GetStockCountSheetRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetStockCountSheetRequest(ctx context.Context, id int) *GetStockCountSheetRequest {
	return &GetStockCountSheetRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetStockCountSheetResponse struct {
	Payload inventory.StockCountSheet `json:"payload"`
}

func (s *InventoryService) NewGetStockCountSheetResponse(payload inventory.StockCountSheet) *GetStockCountSheetResponse {
	return &GetStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetStockCountSheet(req *GetStockCountSheetRequest) (*GetStockCountSheetResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetStockCountSheetByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetStockCountSheetResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE STOCK COUNT SHEET
type CreateStockCountSheetRequest struct {
	Ctx     context.Context
	Payload dto.CreateStockCountSheetDTO
}

func (s *InventoryService) NewCreateStockCountSheetRequest(ctx context.Context, payload dto.CreateStockCountSheetDTO) *CreateStockCountSheetRequest {
	return &CreateStockCountSheetRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateStockCountSheetResponse struct {
	Payload inventory.StockCountSheet `json:"payload"`
}

func (s *InventoryService) NewCreateStockCountSheetResponse(payload inventory.StockCountSheet) *CreateStockCountSheetResponse {
	return &CreateStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) CreateStockCountSheet(req *CreateStockCountSheetRequest) (*CreateStockCountSheetResponse, error) {
	/*
		1. Create StockCountSheet
		2. Update Inventory
		3. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create stock count sheet
	if req.Payload.StockCountSheet.SCSNumber == "" {
		nextID, err := s.Repo.Inventory.GetNextEntryStockCountSheetID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.StockCountSheet.SCSNumber = util.Util.Str.ToString(nextID)
	}
	inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.StockCountSheet.InventoryID.Int)
	if err != nil {
		return nil, err
	}
	err = proc.ProcessStockCountSheetAmounts(&req.Payload.StockCountSheet, inv)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.CreateStockCountSheet(req.Ctx, tx, &req.Payload.StockCountSheet)
	if err != nil {
		return nil, err
	}

	// update inventory
	inv.QuantityAvailable = types.NullDecimal(req.Payload.StockCountSheet.TotalQuantity)
	err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
	if err != nil {
		return nil, err
	}

	// create inventory transaction
	invTxn := inventory.InventoryTransaction{
		InventoryID:     null.IntFrom(inv.ID),
		TransactionType: inventory.InventoryTransactionTypeStockCount,
		Quantity:        req.Payload.StockCountSheet.TotalQuantity,
	}
	err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, &invTxn)
	if err != nil {
		return nil, err
	}

	// get created stock count sheet
	inventory, err := s.Repo.Inventory.GetStockCountSheetByID(req.Ctx, tx, req.Payload.StockCountSheet.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateStockCountSheetResponse{
		Payload: *inventory,
	}

	return &resp, nil
}

// UPDATE STOCK COUNT SHEET
type UpdateStockCountSheetRequest struct {
	Ctx     context.Context
	Payload dto.UpdateStockCountSheetDTO
}

func (s *InventoryService) NewUpdateStockCountSheetRequest(ctx context.Context, payload dto.UpdateStockCountSheetDTO) *UpdateStockCountSheetRequest {
	return &UpdateStockCountSheetRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateStockCountSheetResponse struct {
	Payload inventory.StockCountSheet `json:"payload"`
}

func (s *InventoryService) NewUpdateStockCountSheetResponse(payload inventory.StockCountSheet) *UpdateStockCountSheetResponse {
	return &UpdateStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) UpdateStockCountSheet(req *UpdateStockCountSheetRequest) (*UpdateStockCountSheetResponse, error) {
	/*
		1. Update StockCountSheet
		2. Update Inventory
		3. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// update stock count sheet
	currScs, err := s.Repo.Inventory.GetStockCountSheetByID(req.Ctx, tx, req.Payload.StockCountSheet.ID)
	if err != nil {
		return nil, err
	}
	currScsQuantity, ok := currScs.TotalQuantity.Float64()
	if !ok {
		return nil, errors.New("invalid stock count sheet quantity found while updating stock count sheet")
	}
	newScsQuantity, ok := req.Payload.StockCountSheet.TotalQuantity.Float64()
	if !ok {
		return nil, errors.New("invalid stock count sheet quantity found while updating stock count sheet")
	}
	offsetQuantity := newScsQuantity - currScsQuantity
	req.Payload.StockCountSheet.DiscrepanciesGen.Add(
		req.Payload.StockCountSheet.DiscrepanciesGen.Big,
		types.NewNullDecimal(decimal.New(int64(offsetQuantity*100), 2)).Big,
	)

	// merge empty values
	err = merge.MergeNilOrEmptyValueFields(currScs, &req.Payload.StockCountSheet, true)
	if err != nil {
		return nil, err
	}

	err = s.Repo.Inventory.UpdateStockCountSheet(req.Ctx, tx, &req.Payload.StockCountSheet)
	if err != nil {
		return nil, err
	}

	// update inventory
	inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, req.Payload.StockCountSheet.InventoryID.Int)
	if err != nil {
		return nil, err
	}
	inv.QuantityAvailable = types.NullDecimal(req.Payload.StockCountSheet.TotalQuantity)
	err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
	if err != nil {
		return nil, err
	}

	// create inventory transaction
	invTxn := inventory.InventoryTransaction{
		InventoryID:     null.IntFrom(inv.ID),
		TransactionType: inventory.InventoryTransactionTypeStockCountAdjustment,
		Quantity:        types.NewDecimal(decimal.New(int64(offsetQuantity*100), 2)),
	}
	err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, &invTxn)
	if err != nil {
		return nil, err
	}

	// get updated stock count sheet
	inventory, err := s.Repo.Inventory.GetStockCountSheetByID(req.Ctx, tx, req.Payload.StockCountSheet.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateStockCountSheetResponse{
		Payload: *inventory,
	}

	return &resp, nil
}

// DELETE STOCK COUNT SHEET
type DeleteStockCountSheetRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewDeleteStockCountSheetRequest(ctx context.Context, id int) *DeleteStockCountSheetRequest {
	return &DeleteStockCountSheetRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteStockCountSheetResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewDeleteStockCountSheetResponse(payload bool) *DeleteStockCountSheetResponse {
	return &DeleteStockCountSheetResponse{
		Payload: payload,
	}
}

func (s *InventoryService) DeleteStockCountSheet(req *DeleteStockCountSheetRequest) (*DeleteStockCountSheetResponse, error) {
	/*
		1. Delete StockCountSheet
		2. Update Inventory
		3. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get stock count sheet
	scs, err := s.Repo.Inventory.GetStockCountSheetByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete stock count sheet
	err = s.Repo.Inventory.DeleteStockCountSheet(req.Ctx, tx, scs)
	if err != nil {
		return nil, err
	}

	// update inventory
	inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, scs.InventoryID.Int)
	if err != nil {
		return nil, err
	}
	discrepencyValue, ok := scs.DiscrepanciesGen.Float64()
	if !ok {
		return nil, errors.New("invalid stock count sheet discrepancies found while deleting stock count sheet")
	}
	inv.QuantityAvailable.Add(
		inv.QuantityAvailable.Big,
		types.NewDecimal(decimal.New(int64(discrepencyValue*100), 2)).Big,
	)
	err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
	if err != nil {
		return nil, err
	}

	// create inventory transaction
	invTxn := inventory.InventoryTransaction{
		InventoryID:     null.IntFrom(inv.ID),
		TransactionType: inventory.InventoryTransactionTypeStockCountCancellation,
		Quantity:        types.NewDecimal(decimal.New(int64(discrepencyValue*100), 2)),
	}
	err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, &invTxn)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteStockCountSheetResponse{
		Payload: true,
	}

	return &resp, nil
}
