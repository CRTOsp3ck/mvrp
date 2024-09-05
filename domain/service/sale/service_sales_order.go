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

// LIST SALES ORDER
type ListSalesOrderRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListSalesOrderRequest(ctx context.Context) *ListSalesOrderRequest {
	return &ListSalesOrderRequest{Ctx: ctx}
}

type ListSalesOrderResponse struct {
	Payload sale.SalesOrderSlice `json:"payload"`
}

func (s *SaleService) NewListSalesOrderResponse(payload sale.SalesOrderSlice) *ListSalesOrderResponse {
	return &ListSalesOrderResponse{Payload: payload}
}

func (s *SaleService) ListSalesOrder(req *ListSalesOrderRequest) (*ListSalesOrderResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllSalesOrders(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListSalesOrderResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW SALES ORDER
type PreviewSalesOrderRequest struct {
	Ctx     context.Context
	Payload dto.CreateSalesOrderDTO
}

func (s *SaleService) NewPreviewSalesOrderRequest(ctx context.Context, payload dto.CreateSalesOrderDTO) *PreviewSalesOrderRequest {
	return &PreviewSalesOrderRequest{Ctx: ctx, Payload: payload}
}

type PreviewSalesOrderResponse struct {
	Payload dto.CreateSalesOrderDTO `json:"payload"`
}

func (s *SaleService) NewPreviewSalesOrderResponse(payload dto.CreateSalesOrderDTO) *PreviewSalesOrderResponse {
	return &PreviewSalesOrderResponse{Payload: payload}
}

func (s *SaleService) PreviewSalesOrder(req *PreviewSalesOrderRequest) (*PreviewSalesOrderResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	resp := PreviewSalesOrderResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH SALES ORDER
type SearchSalesOrderRequest struct {
	Ctx     context.Context
	Payload dto.SearchSalesOrderDTO
}

func (s *SaleService) NewSearchSalesOrderRequest(ctx context.Context, payload dto.SearchSalesOrderDTO) *SearchSalesOrderRequest {
	return &SearchSalesOrderRequest{Ctx: ctx, Payload: payload}
}

type SearchSalesOrderResponse struct {
	Payload    sale.SalesOrderSlice `json:"payload"`
	Pagination dto.PaginationDTO    `json:"pagination"`
}

func (s *SaleService) NewSearchSalesOrderResponse(payload sale.SalesOrderSlice) *SearchSalesOrderResponse {
	return &SearchSalesOrderResponse{Payload: payload}
}

func (s *SaleService) SearchSalesOrder(req *SearchSalesOrderRequest) (*SearchSalesOrderResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchSalesOrders(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Sale.GetSalesOrderTotalCount(req.Ctx, tx)
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
	resp := SearchSalesOrderResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET SALES ORDER
type GetSalesOrderRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetSalesOrderRequest(ctx context.Context, id int) *GetSalesOrderRequest {
	return &GetSalesOrderRequest{Ctx: ctx, ID: id}
}

type GetSalesOrderResponse struct {
	Payload sale.SalesOrder `json:"payload"`
}

func (s *SaleService) NewGetSalesOrderResponse(payload sale.SalesOrder) *GetSalesOrderResponse {
	return &GetSalesOrderResponse{Payload: payload}
}

func (s *SaleService) GetSalesOrder(req *GetSalesOrderRequest) (*GetSalesOrderResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetSalesOrderResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE SALES ORDER
type CreateSalesOrderRequest struct {
	Ctx     context.Context
	Payload dto.CreateSalesOrderDTO
}

func (s *SaleService) NewCreateSalesOrderRequest(ctx context.Context, payload dto.CreateSalesOrderDTO) *CreateSalesOrderRequest {
	return &CreateSalesOrderRequest{Ctx: ctx, Payload: payload}
}

type CreateSalesOrderResponse struct {
	Payload sale.SalesOrder `json:"payload"`
}

func (s *SaleService) NewCreateSalesOrderResponse(payload sale.SalesOrder) *CreateSalesOrderResponse {
	return &CreateSalesOrderResponse{Payload: payload}
}

func (s *SaleService) CreateSalesOrder(req *CreateSalesOrderRequest) (*CreateSalesOrderResponse, error) {
	/*
		1. Preprocess Amounts
		2. Create Base Document
		3. Create Sales Order
		4. Create Base Document Items
		5. Create Sales Order Items
		6. Update Inventory
		7. Create Inventory Transaction
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

	// create sales order
	req.Payload.SalesOrder.BaseDocumentID = req.Payload.BaseDocument.ID
	if req.Payload.SalesOrder.SalesOrderNumber == "" {
		nextID, err := s.Repo.Sale.GetNextEntrySalesOrderID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.SalesOrder.SalesOrderNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Sale.CreateSalesOrder(req.Ctx, tx, &req.Payload.SalesOrder)
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

		// create sales order items
		item.SalesOrderItem.BaseDocumentItemID = item.BaseDocumentItem.ID
		item.SalesOrderItem.SalesOrderID = req.Payload.SalesOrder.ID
		err = s.Repo.Sale.CreateSalesOrderItem(req.Ctx, tx, &item.SalesOrderItem)
		if err != nil {
			return nil, err
		}

		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.BaseDocumentItem.Quantity.Big)
		inv.QuantityReserved.Add(inv.QuantityReserved.Big, item.BaseDocumentItem.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     null.IntFrom(inv.ID),
			TransactionType: inventory.InventoryTransactionTypeSale,
			Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}
	}

	// get created sales order
	salesOrder, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.Payload.SalesOrder.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateSalesOrderResponse{
		Payload: *salesOrder,
	}

	return &resp, nil
}

// UPDATE SALES ORDER
type UpdateSalesOrderRequest struct {
	Ctx     context.Context
	Payload dto.UpdateSalesOrderDTO
}

func (s *SaleService) NewUpdateSalesOrderRequest(ctx context.Context, payload dto.UpdateSalesOrderDTO) *UpdateSalesOrderRequest {
	return &UpdateSalesOrderRequest{Ctx: ctx, Payload: payload}
}

type UpdateSalesOrderResponse struct {
	Payload sale.SalesOrder `json:"payload"`
}

func (s *SaleService) NewUpdateSalesOrderResponse(payload sale.SalesOrder) *UpdateSalesOrderResponse {
	return &UpdateSalesOrderResponse{Payload: payload}
}

func (s *SaleService) UpdateSalesOrder(req *UpdateSalesOrderRequest) (*UpdateSalesOrderResponse, error) {
	/*
		1. Preprocess Amounts
		2. Update Base Document
		3. Update Sales Order
		4. Update Base Document Items
		5. Update Sales Order Items
		6. Update Inventory
		7. Create Inventory Transaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currSo, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.Payload.SalesOrder.ID)
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

	// update sales order
	err = s.Repo.Sale.UpdateSalesOrder(req.Ctx, tx, &req.Payload.SalesOrder)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	for _, item := range currSo.R.SalesOrderItems {
		found := false
		for _, newItem := range req.Payload.Items {
			if item.ID == newItem.SalesOrderItem.ID {
				found = true
				break
			}
		}
		if !found {
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

			// delete sales order item
			err = s.Repo.Sale.DeleteSalesOrderItem(req.Ctx, tx, item)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, baseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, baseDocumentItem.Quantity.Big)
			inv.QuantityReserved.Sub(inv.QuantityReserved.Big, baseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeSaleCancellation,
				Quantity:        types.NewDecimal(baseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update sales order items
	for _, item := range req.Payload.Items {
		// check if the item is new or existing
		itemExists, err := s.Repo.Sale.SalesOrderItemExists(req.Ctx, tx, item.SalesOrderItem.ID)
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

			// update sales order items
			err = s.Repo.Sale.UpdateSalesOrderItem(req.Ctx, tx, &item.SalesOrderItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, amountOffset.Big)
			inv.QuantityReserved.Sub(inv.QuantityReserved.Big, amountOffset.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeSaleAdjustment,
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

			// create sales order items
			item.SalesOrderItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.SalesOrderItem.SalesOrderID = req.Payload.SalesOrder.ID
			err = s.Repo.Sale.CreateSalesOrderItem(req.Ctx, tx, &item.SalesOrderItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, item.BaseDocumentItem.ItemID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.BaseDocumentItem.Quantity.Big)
			inv.QuantityReserved.Add(inv.QuantityReserved.Big, item.BaseDocumentItem.Quantity.Big)
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     null.IntFrom(inv.ID),
				TransactionType: inventory.InventoryTransactionTypeSale,
				Quantity:        types.NewDecimal(item.BaseDocumentItem.Quantity.Big),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated sales order
	salesOrder, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.Payload.SalesOrder.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateSalesOrderResponse{
		Payload: *salesOrder,
	}

	return &resp, nil
}

// DELETE SALES ORDER
type DeleteSalesOrderRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewDeleteSalesOrderRequest(ctx context.Context, id int) *DeleteSalesOrderRequest {
	return &DeleteSalesOrderRequest{Ctx: ctx, ID: id}
}

type DeleteSalesOrderResponse struct {
	Payload bool `json:"payload"`
}

func (s *SaleService) NewDeleteSalesOrderResponse(payload bool) *DeleteSalesOrderResponse {
	return &DeleteSalesOrderResponse{Payload: payload}
}

func (s *SaleService) DeleteSalesOrder(req *DeleteSalesOrderRequest) (*DeleteSalesOrderResponse, error) {
	/*
		1. Get Sales Order
		2. Delete Sales Order
		3. Get Base Document
		4. Delete Base Document
		5. Get Base Document Items
		6. Delete Base Document Items
		7. Get Sales Order Items
		8. Delete Sales Order Items
		9. Update Inventory
		10. Create Inventory Transaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get sales order
	salesOrder, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete sales order
	err = s.Repo.Sale.DeleteSalesOrder(req.Ctx, tx, salesOrder)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, salesOrder.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	for _, item := range salesOrder.R.SalesOrderItems {
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

		// delete sales order item
		err = s.Repo.Sale.DeleteSalesOrderItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}

		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByItemID(req.Ctx, tx, baseDocumentItem.ItemID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, baseDocumentItem.Quantity.Big)
		inv.QuantityReserved.Sub(inv.QuantityReserved.Big, baseDocumentItem.Quantity.Big)
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     null.IntFrom(inv.ID),
			TransactionType: inventory.InventoryTransactionTypeSaleCancellation,
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

	resp := DeleteSalesOrderResponse{
		Payload: true,
	}

	return &resp, nil
}
