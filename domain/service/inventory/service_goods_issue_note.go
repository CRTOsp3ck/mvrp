package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/merge"
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

	// Pagination
	totalCount, err := s.Repo.Inventory.GetGoodsIssueNoteTotalCount(req.Ctx, tx)
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
	resp := SearchGoodsIssueNoteResponse{
		Payload:    res,
		Pagination: pd,
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
	Payload dto.GetGoodsIssueNoteDTO `json:"payload"`
}

func (s *InventoryService) NewGetGoodsIssueNoteResponse(payload dto.GetGoodsIssueNoteDTO) *GetGoodsIssueNoteResponse {
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

	ginRes, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// get goods issue note items
	ginItems, err := s.Repo.Inventory.GetGoodsIssueNoteItemsByGoodsIssueNoteID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	ginItemRes := make([]dto.GetGoodsIssueNoteItemDTO, 0)
	for _, item := range ginItems {
		ginItemRes = append(ginItemRes, dto.GetGoodsIssueNoteItemDTO{
			GoodsIssueNoteItem: *item,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetGoodsIssueNoteResponse{
		Payload: dto.GetGoodsIssueNoteDTO{
			GoodsIssueNote: *ginRes,
			Items:          ginItemRes,
		},
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
		err = proc.ProcessGoodsIssueNoteItemAmounts(&item.GoodsIssueNoteItem)
		if err != nil {
			return nil, err
		}
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
		err = proc.ProcessInventoryAmounts(inv)
		if err != nil {
			return nil, err
		}
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeIssuance,
			Quantity:        item.Quantity,
			Reason:          null.StringFrom("Goods Issue Note Creation"),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

	}

	// get created goods issue note
	gin, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.Payload.GoodsIssueNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateGoodsIssueNoteResponse{
		Payload: *gin,
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

	// merge values
	err = merge.MergeNilOrEmptyValueFields(currGin, &req.Payload.GoodsIssueNote, true)
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
	currGinItems, err := s.Repo.Inventory.GetGoodsIssueNoteItemsByGoodsIssueNoteID(req.Ctx, tx, currGin.ID)
	if err != nil {
		return nil, err
	}
	for _, currGinItem := range currGinItems {
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
			err = proc.ProcessInventoryAmounts(inv)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     currGinItem.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeIssuanceCancellation,
				Quantity:        currGinItem.Quantity,
				Reason:          null.StringFrom("Goods Issue Note Adjustment"),
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
			err = proc.ProcessGoodsIssueNoteItemAmounts(&item.GoodsIssueNoteItem)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateGoodsIssueNoteItem(req.Ctx, tx, &item.GoodsIssueNoteItem)
			if err != nil {
				return nil, err
			}

			quantityChanged := amountOffset.Big.Cmp(decimal.New(0, 2)) != 0
			if quantityChanged {
				// update inventory
				inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
				if err != nil {
					return nil, err
				}
				inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, amountOffset.Big)
				err = proc.ProcessInventoryAmounts(inv)
				if err != nil {
					return nil, err
				}
				err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
				if err != nil {
					return nil, err
				}

				// create inventory transaction
				invTx := &inventory.InventoryTransaction{
					InventoryID:     item.InventoryID,
					TransactionType: inventory.InventoryTransactionTypeIssuanceAdjustment,
					Quantity:        types.Decimal(amountOffset),
					Reason:          null.StringFrom("Goods Issue Note Adjustment"),
				}
				err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
				if err != nil {
					return nil, err
				}
			}
		} else {
			// create goods issue note item
			err = proc.ProcessGoodsIssueNoteItemAmounts(&item.GoodsIssueNoteItem)
			if err != nil {
				return nil, err
			}
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
			err = proc.ProcessInventoryAmounts(inv)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     item.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeIssuance,
				Quantity:        item.Quantity,
				Reason:          null.StringFrom("Goods Issue Note Adjustment"),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated goods issue note
	gin, err := s.Repo.Inventory.GetGoodsIssueNoteByID(req.Ctx, tx, req.Payload.GoodsIssueNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateGoodsIssueNoteResponse{
		Payload: *gin,
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
	ginItems, err := s.Repo.Inventory.GetGoodsIssueNoteItemsByGoodsIssueNoteID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range ginItems {
		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, item.Quantity.Big)
		err = proc.ProcessInventoryAmounts(inv)
		if err != nil {
			return nil, err
		}
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeIssuanceCancellation,
			Quantity:        item.Quantity,
			Reason:          null.StringFrom("Goods Issue Note Cancellation"),
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
