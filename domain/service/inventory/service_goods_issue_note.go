package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST GOODS ISSUE NOTE
type ListGoodsIssueNoteRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListGoodsIssueNoteRequest(ctx context.Context) *ListGoodsIssueNoteRequest {
	return &ListGoodsIssueNoteRequest{
		Ctx: ctx,
	}
}

type ListGoodsIssueNoteResponse struct {
	Payload inventory.GoodsIssueNoteSlice `json:"payload"`
}

func (s *InventoryService) NewListGoodsIssueNoteResponse(payload inventory.GoodsIssueNoteSlice) *ListGoodsIssueNoteResponse {
	return &ListGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListGoodsIssueNote(req *ListGoodsIssueNoteRequest) (*ListGoodsIssueNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllGoodsIssueNotes(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListGoodsIssueNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH GOODS ISSUE NOTE
type SearchGoodsIssueNoteRequest struct {
	Ctx     context.Context
	Payload dto.SearchGoodsIssueNoteDTO
}

func (s *InventoryService) NewSearchGoodsIssueNoteRequest(ctx context.Context, payload dto.SearchGoodsIssueNoteDTO) *SearchGoodsIssueNoteRequest {
	return &SearchGoodsIssueNoteRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchGoodsIssueNoteResponse struct {
	Payload    inventory.GoodsIssueNoteSlice `json:"payload"`
	Pagination dto.PaginationDTO             `json:"pagination"`
}

func (s *InventoryService) NewSearchGoodsIssueNoteResponse(payload inventory.GoodsIssueNoteSlice) *SearchGoodsIssueNoteResponse {
	return &SearchGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchGoodsIssueNote(req *SearchGoodsIssueNoteRequest) (*SearchGoodsIssueNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchGoodsIssueNotes(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchGoodsIssueNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET GOODS ISSUE NOTE
type GetGoodsIssueNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetGoodsIssueNoteRequest(ctx context.Context, id int) *GetGoodsIssueNoteRequest {
	return &GetGoodsIssueNoteRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetGoodsIssueNoteResponse struct {
	Payload inventory.GoodsIssueNote `json:"payload"`
}

func (s *InventoryService) NewGetGoodsIssueNoteResponse(payload inventory.GoodsIssueNote) *GetGoodsIssueNoteResponse {
	return &GetGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetGoodsIssueNote(req *GetGoodsIssueNoteRequest) (*GetGoodsIssueNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetGoodsIssueNoteResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE GOODS ISSUE NOTE
type CreateGoodsIssueNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateGoodsIssueNoteDTO
}

func (s *InventoryService) NewCreateGoodsIssueNoteRequest(ctx context.Context, payload dto.CreateGoodsIssueNoteDTO) *CreateGoodsIssueNoteRequest {
	return &CreateGoodsIssueNoteRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateGoodsIssueNoteResponse struct {
	Payload inventory.GoodsIssueNote `json:"payload"`
}

func (s *InventoryService) NewCreateGoodsIssueNoteResponse(payload inventory.GoodsIssueNote) *CreateGoodsIssueNoteResponse {
	return &CreateGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) CreateGoodsIssueNote(req *CreateGoodsIssueNoteRequest) (*CreateGoodsIssueNoteResponse, error) {
	/*
		1. Create GoodsIssueNote
		2. Create GoodsIssueNoteItems
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create goods issue note
	if req.Payload.GoodsIssueNote.GinNumber == "" {
		nextID, err := s.Repo.Inventory.GetNextEntryGoodsIssueNoteID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.GoodsIssueNote.GinNumber = util.Util.Str.ToString(nextID)
	}
	var ginItems []*inventory.GoodsIssueNoteItem
	for _, item := range req.Payload.Items {
		ginItems = append(ginItems, &item.GoodsIssueNoteItem)
	}
	err = proc.ProcessGoodsIssueNoteAmounts(&req.Payload.GoodsIssueNote, ginItems)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.CreateGoodsIssueNote(req.Ctx, tx, &req.Payload.GoodsIssueNote)
	if err != nil {
		return nil, err
	}

	// create goods issue note items
	for _, item := range req.Payload.Items {
		item.GoodsIssueNoteItem.GinID = null.IntFrom(req.Payload.GoodsIssueNote.ID)
		err = s.Repo.Inventory.CreateGoodsIssueNoteItem(req.Ctx, tx, &item.GoodsIssueNoteItem)
		if err != nil {
			return nil, err
		}

		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeIssue,
			Quantity:        item.Quantity,
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

	}

	// get created goods issue note
	inventory, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.Payload.GoodsIssueNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateGoodsIssueNoteResponse{
		Payload: *inventory,
	}

	return &resp, nil
}

// UPDATE GOODS ISSUE NOTE
type UpdateGoodsIssueNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateGoodsIssueNoteDTO
}

func (s *InventoryService) NewUpdateGoodsIssueNoteRequest(ctx context.Context, payload dto.UpdateGoodsIssueNoteDTO) *UpdateGoodsIssueNoteRequest {
	return &UpdateGoodsIssueNoteRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateGoodsIssueNoteResponse struct {
	Payload inventory.GoodsIssueNote `json:"payload"`
}

func (s *InventoryService) NewUpdateGoodsIssueNoteResponse(payload inventory.GoodsIssueNote) *UpdateGoodsIssueNoteResponse {
	return &UpdateGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) UpdateGoodsIssueNote(req *UpdateGoodsIssueNoteRequest) (*UpdateGoodsIssueNoteResponse, error) {
	/*
		1. Update GoodsIssueNote
		2. Update GoodsIssueNote Items
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currGin, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.Payload.GoodsIssueNote.ID)
	if err != nil {
		return nil, err
	}

	// update goods issue note
	var ginItems []*inventory.GoodsIssueNoteItem
	for _, item := range req.Payload.Items {
		ginItems = append(ginItems, &item.GoodsIssueNoteItem)
	}
	err = proc.ProcessGoodsIssueNoteAmounts(&req.Payload.GoodsIssueNote, ginItems)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.UpdateGoodsIssueNote(req.Ctx, tx, &req.Payload.GoodsIssueNote)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	for _, currGinItem := range currGin.R.GinGoodsIssueNoteItems {
		found := false
		for _, item := range req.Payload.Items {
			if currGinItem.ID == item.ID {
				found = true
				break
			}
		}
		if !found {
			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, currGinItem.InventoryID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, currGinItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     currGinItem.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeIssueCancellation,
				Quantity:        currGinItem.Quantity,
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}

			// delete goods issue note item
			err = s.Repo.Inventory.DeleteGoodsIssueNoteItem(req.Ctx, tx, currGinItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update goods issue note items
	for _, item := range req.Payload.Items {
		// check if the item is new
		itemExists, err := s.Repo.Inventory.GoodsIssueNoteItemExists(req.Ctx, tx, item.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			currGinItem, err := s.Repo.Inventory.GetGoodsIssueNoteItemByID(req.Ctx, tx, item.ID)
			if err != nil {
				return nil, err
			}
			amountOffset := types.NewNullDecimal(decimal.New(0, 2))
			amountOffset.Sub(item.Quantity.Big, currGinItem.Quantity.Big)

			// update goods issue note item
			err = s.Repo.Inventory.UpdateGoodsIssueNoteItem(req.Ctx, tx, &item.GoodsIssueNoteItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, amountOffset.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     item.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeIssueAdjustment,
				Quantity:        types.Decimal(amountOffset),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		} else {
			// create goods issue note item
			err = s.Repo.Inventory.CreateGoodsIssueNoteItem(req.Ctx, tx, &item.GoodsIssueNoteItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     item.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeIssue,
				Quantity:        item.Quantity,
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated sales order
	salesOrder, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.Payload.GoodsIssueNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateGoodsIssueNoteResponse{
		Payload: *salesOrder,
	}

	return &resp, nil
}

// DELETE GOODS ISSUE NOTE
type DeleteGoodsIssueNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewDeleteGoodsIssueNoteRequest(ctx context.Context, id int) *DeleteGoodsIssueNoteRequest {
	return &DeleteGoodsIssueNoteRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteGoodsIssueNoteResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewDeleteGoodsIssueNoteResponse(payload bool) *DeleteGoodsIssueNoteResponse {
	return &DeleteGoodsIssueNoteResponse{
		Payload: payload,
	}
}

func (s *InventoryService) DeleteGoodsIssueNote(req *DeleteGoodsIssueNoteRequest) (*DeleteGoodsIssueNoteResponse, error) {
	/*
		1. Delete GoodsIssueNote
		2. Delete GoodsIssueNoteItems
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get goods issue note
	gin, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete goods issue note
	err = s.Repo.Inventory.DeleteGoodsIssueNote(req.Ctx, tx, gin)
	if err != nil {
		return nil, err
	}

	// delete goods issue note items
	for _, item := range gin.R.GinGoodsIssueNoteItems {
		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, item.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeIssueCancellation,
			Quantity:        item.Quantity,
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

		// delete goods issue note item
		err = s.Repo.Inventory.DeleteGoodsIssueNoteItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteGoodsIssueNoteResponse{
		Payload: true,
	}

	return &resp, nil
}
