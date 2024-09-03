package sale

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/inventory"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST GOODS RETURN NOTE
type ListGoodsReturnNoteRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListGoodsReturnNoteRequest(ctx context.Context) *ListGoodsReturnNoteRequest {
	return &ListGoodsReturnNoteRequest{Ctx: ctx}
}

type ListGoodsReturnNoteResponse struct {
	Payload sale.GoodsReturnNoteSlice
}

func (s *SaleService) NewListGoodsReturnNoteResponse(payload sale.GoodsReturnNoteSlice) *ListGoodsReturnNoteResponse {
	return &ListGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) ListGoodsReturnNote(req *ListGoodsReturnNoteRequest) (*ListGoodsReturnNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllGoodsReturnNotes(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListGoodsReturnNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW GOODS RETURN NOTE
type PreviewGoodsReturnNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateGoodsReturnNoteDTO
}

func (s *SaleService) NewPreviewGoodsReturnNoteRequest(ctx context.Context, payload dto.CreateGoodsReturnNoteDTO) *PreviewGoodsReturnNoteRequest {
	return &PreviewGoodsReturnNoteRequest{Ctx: ctx, Payload: payload}
}

type PreviewGoodsReturnNoteResponse struct {
	Payload dto.CreateGoodsReturnNoteDTO
}

func (s *SaleService) NewPreviewGoodsReturnNoteResponse(payload dto.CreateGoodsReturnNoteDTO) *PreviewGoodsReturnNoteResponse {
	return &PreviewGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) PreviewGoodsReturnNote(req *PreviewGoodsReturnNoteRequest) (*PreviewGoodsReturnNoteResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	resp := PreviewGoodsReturnNoteResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH GOODS RETURN NOTE
type SearchGoodsReturnNoteRequest struct {
	Ctx     context.Context
	Payload dto.SearchGoodsReturnNoteDTO
}

func (s *SaleService) NewSearchGoodsReturnNoteRequest(ctx context.Context, payload dto.SearchGoodsReturnNoteDTO) *SearchGoodsReturnNoteRequest {
	return &SearchGoodsReturnNoteRequest{Ctx: ctx, Payload: payload}
}

type SearchGoodsReturnNoteResponse struct {
	Payload sale.GoodsReturnNoteSlice
}

func (s *SaleService) NewSearchGoodsReturnNoteResponse(payload sale.GoodsReturnNoteSlice) *SearchGoodsReturnNoteResponse {
	return &SearchGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) SearchGoodsReturnNote(req *SearchGoodsReturnNoteRequest) (*SearchGoodsReturnNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchGoodsReturnNotes(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchGoodsReturnNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET GOODS RETURN NOTE
type GetGoodsReturnNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetGoodsReturnNoteRequest(ctx context.Context, id int) *GetGoodsReturnNoteRequest {
	return &GetGoodsReturnNoteRequest{Ctx: ctx, ID: id}
}

type GetGoodsReturnNoteResponse struct {
	Payload sale.GoodsReturnNote
}

func (s *SaleService) NewGetGoodsReturnNoteResponse(payload sale.GoodsReturnNote) *GetGoodsReturnNoteResponse {
	return &GetGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) GetGoodsReturnNote(req *GetGoodsReturnNoteRequest) (*GetGoodsReturnNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetGoodsReturnNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetGoodsReturnNoteResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE GOODS RETURN NOTE
type CreateGoodsReturnNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateGoodsReturnNoteDTO
}

func (s *SaleService) NewCreateGoodsReturnNoteRequest(ctx context.Context, payload dto.CreateGoodsReturnNoteDTO) *CreateGoodsReturnNoteRequest {
	return &CreateGoodsReturnNoteRequest{Ctx: ctx, Payload: payload}
}

type CreateGoodsReturnNoteResponse struct {
	Payload sale.GoodsReturnNote
}

func (s *SaleService) NewCreateGoodsReturnNoteResponse(payload sale.GoodsReturnNote) *CreateGoodsReturnNoteResponse {
	return &CreateGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) CreateGoodsReturnNote(req *CreateGoodsReturnNoteRequest) (*CreateGoodsReturnNoteResponse, error) {
	/*
		1. Preprocess Amounts
		2. Create Base Document
		3. Create Goods Return Note
		4. Create Base Document Items
		5. Create Goods Return Note Items
		6. Update Inventory
		7. Create Inventory Transaction
		8. Create Return Merchandise Authorization
		9. Create Return Merchandise Authorization Items
		10. Create Credit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err = proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	// create base document
	err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// create goods return note
	req.Payload.GoodsReturnNote.BaseDocumentID = req.Payload.BaseDocument.ID
	if req.Payload.GoodsReturnNote.GoodsReturnNoteNumber == "" {
		nextID, err := s.Repo.Sale.GetNextEntryGoodsReturnNoteID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.GoodsReturnNote.GoodsReturnNoteNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Sale.CreateGoodsReturnNote(req.Ctx, tx, &req.Payload.GoodsReturnNote)
	if err != nil {
		return nil, err
	}

	// create return merchandise authorization
	nextID, err := s.Repo.Inventory.GetNextEntryReturnMerchandiseAuthorizationID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	rma := &inventory.ReturnMerchandiseAuthorization{
		RmaNumber: util.Util.Str.ToString(nextID),
	}
	err = s.Repo.Inventory.CreateReturnMerchandiseAuthorization(req.Ctx, tx, rma)
	if err != nil {
		return nil, err
	}
	rmaItems := make([]*inventory.ReturnMerchandiseAuthorizationItem, 0)

	for _, item := range req.Payload.Items {
		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityReturned.Add(inv.QuantityReturned.Big, item.BaseDocumentItem.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     null.IntFrom(inv.ID),
			TransactionType: inventory.InventoryTransactionTypeReturn,
			Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

		// create return merchandise authorization items
		rmaItem := &inventory.ReturnMerchandiseAuthorizationItem{
			RmaID:       null.IntFrom(rma.ID),
			InventoryID: null.IntFrom(inv.ID),
			Quantity:    types.Decimal(item.BaseDocumentItem.Quantity),
			UnitValue:   types.Decimal(item.BaseDocumentItem.UnitPrice),
		}
		err = s.Repo.Inventory.CreateReturnMerchandiseAuthorizationItem(req.Ctx, tx, rmaItem)
		if err != nil {
			return nil, err
		}
		rmaItems = append(rmaItems, rmaItem)

		// create base document items
		item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
		err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
		if err != nil {
			return nil, err
		}

		// create goods return note items
		item.GoodsReturnNoteItem.BaseDocumentItemID = item.BaseDocumentItem.ID
		item.GoodsReturnNoteItem.GoodsReturnNoteID = req.Payload.GoodsReturnNote.ID
		item.GoodsReturnNoteItem.RmaItemID = null.IntFrom(rmaItem.ID)
		err = s.Repo.Sale.CreateGoodsReturnNoteItem(req.Ctx, tx, &item.GoodsReturnNoteItem)
		if err != nil {
			return nil, err
		}
	}

	// preprocess amounts (update total value)
	err = proc.ProcessReturnMerchandiseAuthorizationAmounts(rma, rmaItems)
	if err != nil {
		return nil, err
	}

	// update return merchandise authorization with the total value
	err = s.Repo.Inventory.UpdateReturnMerchandiseAuthorization(req.Ctx, tx, rma)
	if err != nil {
		return nil, err
	}

	// get created goods return note
	goodsReturnNote, err := s.Repo.Sale.GetGoodsReturnNoteByID(req.Ctx, tx, req.Payload.GoodsReturnNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateGoodsReturnNoteResponse{
		Payload: *goodsReturnNote,
	}

	return &resp, nil
}

// UPDATE GOODS RETURN NOTE
type UpdateGoodsReturnNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateGoodsReturnNoteDTO
}

func (s *SaleService) NewUpdateGoodsReturnNoteRequest(ctx context.Context, payload dto.UpdateGoodsReturnNoteDTO) *UpdateGoodsReturnNoteRequest {
	return &UpdateGoodsReturnNoteRequest{Ctx: ctx, Payload: payload}
}

type UpdateGoodsReturnNoteResponse struct {
	Payload sale.GoodsReturnNote
}

func (s *SaleService) NewUpdateGoodsReturnNoteResponse(payload sale.GoodsReturnNote) *UpdateGoodsReturnNoteResponse {
	return &UpdateGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) UpdateGoodsReturnNote(req *UpdateGoodsReturnNoteRequest) (*UpdateGoodsReturnNoteResponse, error) {
	/*
		1. Preprocess Amounts
		2. Update Base Document
		3. Update Goods Return Note
		4. Update Base Document Items
		5. Update Goods Return Note Items
		6. Update Inventory
		7. Create Inventory Transaction
		8. Update Return Merchandise Authorization
		9. Update Return Merchandise Authorization Items
		10. Update Credit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currGrn, err := s.Repo.Sale.GetGoodsReturnNoteByID(req.Ctx, tx, req.Payload.GoodsReturnNote.ID)
	if err != nil {
		return nil, err
	}

	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err = proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	// update base document
	err = s.Repo.Base.UpdateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// update goods return note
	err = s.Repo.Sale.UpdateGoodsReturnNote(req.Ctx, tx, &req.Payload.GoodsReturnNote)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	for _, currItem := range currGrn.R.GoodsReturnNoteItems {
		found := false
		for _, item := range req.Payload.Items {
			if currItem.ID == item.GoodsReturnNoteItem.ID {
				found = true
				break
			}
		}
		if !found {
			// get base document item
			baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, currItem.BaseDocumentItemID)
			if err != nil {
				return nil, err
			}

			// delete base document item
			err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
			if err != nil {
				return nil, err
			}

			// delete return merchandise authorization item
			rmaItem, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemByID(req.Ctx, tx, currItem.RmaItemID.Int)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorizationItem(req.Ctx, tx, rmaItem)
			if err != nil {
				return nil, err
			}

			// delete goods return note item
			err = s.Repo.Sale.DeleteGoodsReturnNoteItem(req.Ctx, tx, currItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, baseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReturned.Sub(inv.QuantityReturned.Big, baseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeReturnCancellation,
				Quantity:        types.NewDecimal(baseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update goods return note items
	for _, item := range req.Payload.Items {
		// check if the item is new
		itemExists, err := s.Repo.Sale.GoodsReturnNoteItemExists(req.Ctx, tx, item.GoodsReturnNoteItem.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			currBaseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItem.ID)
			if err != nil {
				return nil, err
			}
			amountOffset := types.NewNullDecimal(decimal.New(0, 2))
			amountOffset.Sub(item.Quantity.Big, currBaseDocumentItem.Quantity.Big)

			// update base document items
			err = s.Repo.Base.UpdateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// update goods return note items
			err = s.Repo.Sale.UpdateGoodsReturnNoteItem(req.Ctx, tx, &item.GoodsReturnNoteItem)
			if err != nil {
				return nil, err
			}

			// update return merchandise authorization item
			rmaItem, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemByID(req.Ctx, tx, item.GoodsReturnNoteItem.RmaItemID.Int)
			if err != nil {
				return nil, err
			}
			rmaItem.Quantity = types.Decimal(item.BaseDocumentItem.Quantity)
			rmaItem.UnitValue = types.Decimal(item.BaseDocumentItem.UnitPrice)
			err = s.Repo.Inventory.UpdateReturnMerchandiseAuthorizationItem(req.Ctx, tx, rmaItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReturned.Add(inv.QuantityReturned.Big, amountOffset.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeReturnAdjustment,
				Quantity:        types.NewDecimal(amountOffset.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		} else {
			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReturned.Add(inv.QuantityReturned.Big, item.BaseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeReturn,
				Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}

			// create return merchandise authorization item
			rmaItem := &inventory.ReturnMerchandiseAuthorizationItem{
				RmaID:       req.Payload.GoodsReturnNote.RmaID,
				InventoryID: null.IntFrom(inv.ID),
				Quantity:    types.Decimal(item.BaseDocumentItem.Quantity),
				UnitValue:   types.Decimal(item.BaseDocumentItem.UnitPrice),
			}
			err = s.Repo.Inventory.CreateReturnMerchandiseAuthorizationItem(req.Ctx, tx, rmaItem)
			if err != nil {
				return nil, err
			}

			// create base document items
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create goods return note items
			item.GoodsReturnNoteItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.GoodsReturnNoteItem.GoodsReturnNoteID = req.Payload.GoodsReturnNote.ID
			err = s.Repo.Sale.CreateGoodsReturnNoteItem(req.Ctx, tx, &item.GoodsReturnNoteItem)
			if err != nil {
				return nil, err
			}

		}
	}

	// update return merchandise authorization with the total value
	rma, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.Payload.GoodsReturnNote.RmaID.Int)
	if err != nil {
		return nil, err
	}
	rmaItems, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(req.Ctx, tx, rma.ID)
	if err != nil {
		return nil, err
	}
	err = proc.ProcessReturnMerchandiseAuthorizationAmounts(rma, rmaItems)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.UpdateReturnMerchandiseAuthorization(req.Ctx, tx, rma)
	if err != nil {
		return nil, err
	}

	// get updated goods return note
	goodsReturnNote, err := s.Repo.Sale.GetGoodsReturnNoteByID(req.Ctx, tx, req.Payload.GoodsReturnNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateGoodsReturnNoteResponse{
		Payload: *goodsReturnNote,
	}

	return &resp, nil
}

// DELETE GOODS RETURN NOTE
type DeleteGoodsReturnNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewDeleteGoodsReturnNoteRequest(ctx context.Context, id int) *DeleteGoodsReturnNoteRequest {
	return &DeleteGoodsReturnNoteRequest{Ctx: ctx, ID: id}
}

type DeleteGoodsReturnNoteResponse struct {
	Payload bool
}

func (s *SaleService) NewDeleteGoodsReturnNoteResponse(payload bool) *DeleteGoodsReturnNoteResponse {
	return &DeleteGoodsReturnNoteResponse{Payload: payload}
}

func (s *SaleService) DeleteGoodsReturnNote(req *DeleteGoodsReturnNoteRequest) (*DeleteGoodsReturnNoteResponse, error) {
	/*
		1. Delete Goods Return Note
		2. Delete Base Document
		3. Delete Base Document Items
		4. Delete Goods Return Note Items
		5. Update Inventory
		6. Create Inventory Transaction
		7. Delete Return Merchandise Authorization
		8. Delete Return Merchandise Authorization Items
		9. Delete Credit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get goods return note
	goodsReturnNote, err := s.Repo.Sale.GetGoodsReturnNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete goods return note
	err = s.Repo.Sale.DeleteGoodsReturnNote(req.Ctx, tx, goodsReturnNote)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, goodsReturnNote.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	// get return merchandise authorization
	rma, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, goodsReturnNote.RmaID.Int)
	if err != nil {
		return nil, err
	}

	// delete return merchandise authorization
	err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorization(req.Ctx, tx, rma)
	if err != nil {
		return nil, err
	}

	for _, item := range goodsReturnNote.R.GoodsReturnNoteItems {
		// get base document item
		baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItemID)
		if err != nil {
			return nil, err
		}

		// delete base document item
		err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
		if err != nil {
			return nil, err
		}

		// delete goods return note item
		err = s.Repo.Sale.DeleteGoodsReturnNoteItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}

		// get return merchandise authorization item
		rmaItem, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemByID(req.Ctx, tx, item.RmaItemID.Int)
		if err != nil {
			return nil, err
		}

		// delete return merchandise authorization item
		err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorizationItem(req.Ctx, tx, rmaItem)
		if err != nil {
			return nil, err
		}

		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, baseDocumentItem.ItemID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityReturned.Sub(inv.QuantityReturned.Big, baseDocumentItem.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     null.IntFrom(inv.ID),
			TransactionType: inventory.InventoryTransactionTypeReturnCancellation,
			Quantity:        types.NewDecimal(baseDocumentItem.Quantity.Big),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteGoodsReturnNoteResponse{
		Payload: true,
	}

	return &resp, nil
}
