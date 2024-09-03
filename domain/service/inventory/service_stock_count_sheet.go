package inventory

import (
	"context"
	"errors"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST STOCK COUNT SHEET
type ListStockCountSheetRequest struct {
	Ctx context.Context
}
type ListStockCountSheetResponse struct {
	Payload inventory.StockCountSheetSlice
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
type SearchStockCountSheetResponse struct {
	Payload inventory.StockCountSheetSlice
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
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchStockCountSheetResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET STOCK COUNT SHEET
type GetStockCountSheetRequest struct {
	Ctx context.Context
	ID  int
}
type GetStockCountSheetResponse struct {
	Payload inventory.StockCountSheet
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
type CreateStockCountSheetResponse struct {
	Payload inventory.StockCountSheet
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
	Payload dto.CreateStockCountSheetDTO
}
type UpdateStockCountSheetResponse struct {
	Payload inventory.StockCountSheet
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
	req.Payload.StockCountSheet.Discrepancies.Add(
		req.Payload.StockCountSheet.Discrepancies.Big,
		types.NewNullDecimal(decimal.New(int64(offsetQuantity*100), 2)).Big,
	)
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
type DeleteStockCountSheetResponse struct {
	Payload bool
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
	discrepencyValue, ok := scs.Discrepancies.Float64()
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
