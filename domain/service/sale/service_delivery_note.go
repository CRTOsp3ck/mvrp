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
	"github.com/jinzhu/copier"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST DELIVERY NOTE
type ListDeliveryNoteRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListDeliveryNoteRequest(ctx context.Context) *ListDeliveryNoteRequest {
	return &ListDeliveryNoteRequest{Ctx: ctx}
}

type ListDeliveryNoteResponse struct {
	Payload sale.DeliveryNoteSlice `json:"payload"`
}

func (s *SaleService) NewListDeliveryNoteResponse(payload sale.DeliveryNoteSlice) *ListDeliveryNoteResponse {
	return &ListDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) ListDeliveryNote(req *ListDeliveryNoteRequest) (*ListDeliveryNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllDeliveryNotes(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListDeliveryNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW DELIVERY NOTE
type PreviewDeliveryNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateDeliveryNoteDTO
}

func (s *SaleService) NewPreviewDeliveryNoteRequest(ctx context.Context, payload dto.CreateDeliveryNoteDTO) *PreviewDeliveryNoteRequest {
	return &PreviewDeliveryNoteRequest{Ctx: ctx, Payload: payload}
}

type PreviewDeliveryNoteResponse struct {
	Payload dto.CreateDeliveryNoteDTO `json:"payload"`
}

func (s *SaleService) NewPreviewDeliveryNoteResponse(payload dto.CreateDeliveryNoteDTO) *PreviewDeliveryNoteResponse {
	return &PreviewDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) PreviewDeliveryNote(req *PreviewDeliveryNoteRequest) (*PreviewDeliveryNoteResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	resp := PreviewDeliveryNoteResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE
type SearchDeliveryNoteRequest struct {
	Ctx     context.Context
	Payload dto.SearchDeliveryNoteDTO
}

func (s *SaleService) NewSearchDeliveryNoteRequest(ctx context.Context, payload dto.SearchDeliveryNoteDTO) *SearchDeliveryNoteRequest {
	return &SearchDeliveryNoteRequest{Ctx: ctx, Payload: payload}
}

type SearchDeliveryNoteResponse struct {
	Payload    sale.DeliveryNoteSlice `json:"payload"`
	Pagination dto.PaginationDTO      `json:"pagination"`
}

func (s *SaleService) NewSearchDeliveryNoteResponse(payload sale.DeliveryNoteSlice) *SearchDeliveryNoteResponse {
	return &SearchDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) SearchDeliveryNote(req *SearchDeliveryNoteRequest) (*SearchDeliveryNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchDeliveryNotes(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Sale.GetDeliveryNoteTotalCount(req.Ctx, tx)
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
	resp := SearchDeliveryNoteResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetDeliveryNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetDeliveryNoteRequest(ctx context.Context, id int) *GetDeliveryNoteRequest {
	return &GetDeliveryNoteRequest{Ctx: ctx, ID: id}
}

type GetDeliveryNoteResponse struct {
	Payload sale.DeliveryNote `json:"payload"`
}

func (s *SaleService) NewGetDeliveryNoteResponse(payload sale.DeliveryNote) *GetDeliveryNoteResponse {
	return &GetDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) GetDeliveryNote(req *GetDeliveryNoteRequest) (*GetDeliveryNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetDeliveryNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetDeliveryNoteResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE DELIVERY NOTE
type CreateDeliveryNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateDeliveryNoteDTO
}

func (s *SaleService) NewCreateDeliveryNoteRequest(ctx context.Context, payload dto.CreateDeliveryNoteDTO) *CreateDeliveryNoteRequest {
	return &CreateDeliveryNoteRequest{Ctx: ctx, Payload: payload}
}

type CreateDeliveryNoteResponse struct {
	Payload sale.DeliveryNote `json:"payload"`
}

func (s *SaleService) NewCreateDeliveryNoteResponse(payload sale.DeliveryNote) *CreateDeliveryNoteResponse {
	return &CreateDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) CreateDeliveryNote(req *CreateDeliveryNoteRequest) (*CreateDeliveryNoteResponse, error) {
	/*
		1. Preprocess Amounts
		2. Create Base Document
		3. Create Delivery Note
		4. Create Base Document Items
		5. Create Delivery Note Items
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if req.Payload.ToCreateFromSalesOrder {
		// get sales order
		salesOrder, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.Payload.DeliveryNote.SalesOrderID)
		if err != nil {
			return nil, err
		}
		// get base document
		salesOrderBaseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, salesOrder.BaseDocumentID)
		if err != nil {
			return nil, err
		}
		// get sales order items
		salesOrderItems, err := s.Repo.Sale.GetSalesOrderItemsBySalesOrderID(req.Ctx, tx, salesOrder.ID)
		if err != nil {
			return nil, err
		}
		// get base document items
		salesOrderBaseDocumentItems, err := s.Repo.Base.GetBaseDocumentItemsByBaseDocumentID(req.Ctx, tx, salesOrder.BaseDocumentID)
		if err != nil {
			return nil, err
		}

		// copy & create base document
		var deliveryNoteBaseDocument base.BaseDocument
		copier.Copy(&deliveryNoteBaseDocument, &salesOrderBaseDocument)
		deliveryNoteBaseDocument.ID = -1
		err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &deliveryNoteBaseDocument)
		if err != nil {
			return nil, err
		}

		// copy & create delivery note
		var deliveryNote sale.DeliveryNote
		copier.Copy(&deliveryNote, &req.Payload.DeliveryNote)
		deliveryNote.BaseDocumentID = deliveryNoteBaseDocument.ID
		// additional required fields
		/*
			    bill_to_information TEXT,
				delivery_date DATE,
				shipping_personnel_information TEXT,
				received_by TEXT,
				goods_condition TEXT
		*/
		if deliveryNote.DeliveryNoteNumber == "" {
			nextID, err := s.Repo.Sale.GetNextEntryDeliveryNoteID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			deliveryNote.DeliveryNoteNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateDeliveryNote(req.Ctx, tx, &deliveryNote)
		if err != nil {
			return nil, err
		}

		// copy & create base document items
		for _, item := range salesOrderBaseDocumentItems {
			var deliveryNoteBaseDocumentItem base.BaseDocumentItem
			copier.Copy(&deliveryNoteBaseDocumentItem, &item)
			deliveryNoteBaseDocumentItem.ID = -1
			deliveryNoteBaseDocumentItem.BaseDocumentID = deliveryNoteBaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &deliveryNoteBaseDocumentItem)
			if err != nil {
				return nil, err
			}
		}

		// copy & create delivery note items
		for _, item := range salesOrderItems {
			var deliveryNoteItem sale.DeliveryNoteItem
			copier.Copy(&deliveryNoteItem, &item)
			deliveryNoteItem.ID = -1
			deliveryNoteItem.DeliveryNoteID = deliveryNote.ID
			err = s.Repo.Sale.CreateDeliveryNoteItem(req.Ctx, tx, &deliveryNoteItem)
			if err != nil {
				return nil, err
			}
		}
	} else {
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

		// create delivery note
		req.Payload.DeliveryNote.BaseDocumentID = req.Payload.BaseDocument.ID
		if req.Payload.DeliveryNote.DeliveryNoteNumber == "" {
			nextID, err := s.Repo.Sale.GetNextEntryDeliveryNoteID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			req.Payload.DeliveryNote.DeliveryNoteNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateDeliveryNote(req.Ctx, tx, &req.Payload.DeliveryNote)
		if err != nil {
			return nil, err
		}

		for _, item := range req.Payload.Items {
			// create base document items
			item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create delivery note items
			item.DeliveryNoteItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.DeliveryNoteItem.DeliveryNoteID = req.Payload.DeliveryNote.ID
			err = s.Repo.Sale.CreateDeliveryNoteItem(req.Ctx, tx, &item.DeliveryNoteItem)
			if err != nil {
				return nil, err
			}

			// update inventory (remove from reserved)
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReserved.Sub(inv.QuantityReserved.Big, item.BaseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeShipping,
				Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get created delivery note
	DeliveryNote, err := s.Repo.Sale.GetDeliveryNoteByID(req.Ctx, tx, req.Payload.DeliveryNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateDeliveryNoteResponse{
		Payload: *DeliveryNote,
	}

	return &resp, nil
}

// UPDATE DELIVERY NOTE
type UpdateDeliveryNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateDeliveryNoteDTO
}

func (s *SaleService) NewUpdateDeliveryNoteRequest(ctx context.Context, payload dto.UpdateDeliveryNoteDTO) *UpdateDeliveryNoteRequest {
	return &UpdateDeliveryNoteRequest{Ctx: ctx, Payload: payload}
}

type UpdateDeliveryNoteResponse struct {
	Payload sale.DeliveryNote `json:"payload"`
}

func (s *SaleService) NewUpdateDeliveryNoteResponse(payload sale.DeliveryNote) *UpdateDeliveryNoteResponse {
	return &UpdateDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) UpdateDeliveryNote(req *UpdateDeliveryNoteRequest) (*UpdateDeliveryNoteResponse, error) {
	/*
		1. Preprocess Amounts
		2. Update Base Document
		3. Update Delivery Note
		4. Update Base Document Items
		5. Update Delivery Note Items
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currDn, err := s.Repo.Sale.GetDeliveryNoteByID(req.Ctx, tx, req.Payload.DeliveryNote.ID)
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

	// update delivery note
	err = s.Repo.Sale.UpdateDeliveryNote(req.Ctx, tx, &req.Payload.DeliveryNote)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list but not in the new list
	for _, currDni := range currDn.R.DeliveryNoteItems {
		found := false
		for _, newDni := range req.Payload.Items {
			if currDni.ID == newDni.DeliveryNoteItem.ID {
				found = true
				break
			}
		}
		if !found {
			// get base document item
			currBaseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, currDni.BaseDocumentItemID)
			if err != nil {
				return nil, err
			}
			// get delivery note item
			currDeliveryNoteItem := currDni
			// update inventory (add back to reserved)
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, currBaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReserved.Add(inv.QuantityReserved.Big, currBaseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}
			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeShippingCancellation,
				Quantity:        types.NewDecimal(currBaseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
			// delete base document item
			err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, currBaseDocumentItem)
			if err != nil {
				return nil, err
			}
			// delete delivery note item
			err = s.Repo.Sale.DeleteDeliveryNoteItem(req.Ctx, tx, currDeliveryNoteItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update delivery note items
	for _, item := range req.Payload.Items {
		// check if the item is new or existing
		itemExists, err := s.Repo.Sale.DeliveryNoteItemExists(req.Ctx, tx, item.DeliveryNoteItem.ID)
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

			// update delivery note items
			err = s.Repo.Sale.UpdateDeliveryNoteItem(req.Ctx, tx, &item.DeliveryNoteItem)
			if err != nil {
				return nil, err
			}

			// update inventory (add back to reserved)
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReserved.Add(inv.QuantityReserved.Big, amountOffset.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeShippingAdjustment,
				Quantity:        types.NewDecimal(amountOffset.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		} else {
			// create base document items
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create delivery note items
			item.DeliveryNoteItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.DeliveryNoteItem.DeliveryNoteID = req.Payload.DeliveryNote.ID
			err = s.Repo.Sale.CreateDeliveryNoteItem(req.Ctx, tx, &item.DeliveryNoteItem)
			if err != nil {
				return nil, err
			}

			// update inventory (remove from reserved)
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityReserved.Sub(inv.QuantityReserved.Big, item.BaseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeShipping,
				Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated delivery note
	DeliveryNote, err := s.Repo.Sale.GetDeliveryNoteByID(req.Ctx, tx, req.Payload.DeliveryNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateDeliveryNoteResponse{
		Payload: *DeliveryNote,
	}

	return &resp, nil
}

// DELETE DELIVERY NOTE
type DeleteDeliveryNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewDeleteDeliveryNoteRequest(ctx context.Context, id int) *DeleteDeliveryNoteRequest {
	return &DeleteDeliveryNoteRequest{Ctx: ctx, ID: id}
}

type DeleteDeliveryNoteResponse struct {
	Payload bool `json:"payload"`
}

func (s *SaleService) NewDeleteDeliveryNoteResponse(payload bool) *DeleteDeliveryNoteResponse {
	return &DeleteDeliveryNoteResponse{Payload: payload}
}

func (s *SaleService) DeleteDeliveryNote(req *DeleteDeliveryNoteRequest) (*DeleteDeliveryNoteResponse, error) {
	/*
		1. Get Delivery Note
		2. Delete Delivery Note
		3. Get Base Document
		4. Delete Base Document
		5. Get Base Document Items
		6. Delete Base Document Items
		7. Get Delivery Note Items
		8. Delete Delivery Note Items
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get delivery note
	DeliveryNote, err := s.Repo.Sale.GetDeliveryNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete delivery note
	err = s.Repo.Sale.DeleteDeliveryNote(req.Ctx, tx, DeliveryNote)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, DeliveryNote.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	for _, item := range DeliveryNote.R.DeliveryNoteItems {
		// get base document item
		baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItemID)
		if err != nil {
			return nil, err
		}

		// update inventory (add back to reserved)
		inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, baseDocumentItem.ItemID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityReserved.Add(inv.QuantityReserved.Big, baseDocumentItem.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     null.IntFrom(inv.ID),
			TransactionType: inventory.InventoryTransactionTypeShippingCancellation,
			Quantity:        types.NewDecimal(baseDocumentItem.Quantity.Big),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

		// delete base document item
		err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
		if err != nil {
			return nil, err
		}

		// delete delivery note item
		err = s.Repo.Sale.DeleteDeliveryNoteItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteDeliveryNoteResponse{
		Payload: true,
	}

	return &resp, nil
}
