package sale

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/sale"
	"mvrp/data/repo"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/jinzhu/copier"
)

// LIST ORDER CONFIRMATION
type ListOrderConfirmationRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *SaleService) NewListOrderConfirmationRequest(ctx context.Context) *ListOrderConfirmationRequest {
	return &ListOrderConfirmationRequest{Ctx: ctx}
}

type ListOrderConfirmationResponse struct {
	Payload sale.OrderConfirmationSlice `json:"payload"`
}

func (s *SaleService) NewListOrderConfirmationResponse(payload sale.OrderConfirmationSlice) *ListOrderConfirmationResponse {
	return &ListOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) ListOrderConfirmation(req *ListOrderConfirmationRequest) (*ListOrderConfirmationResponse, error) {
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

	res, err := s.Repo.Sale.ListAllOrderConfirmations(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListOrderConfirmationResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW ORDER CONFIRMATION
type PreviewOrderConfirmationRequest struct {
	Ctx     context.Context
	Payload dto.CreateOrderConfirmationDTO
}

func (s *SaleService) NewPreviewOrderConfirmationRequest(ctx context.Context, payload dto.CreateOrderConfirmationDTO) *PreviewOrderConfirmationRequest {
	return &PreviewOrderConfirmationRequest{Ctx: ctx, Payload: payload}
}

type PreviewOrderConfirmationResponse struct {
	Payload dto.CreateOrderConfirmationDTO `json:"payload"`
}

func (s *SaleService) NewPreviewOrderConfirmationResponse(payload dto.CreateOrderConfirmationDTO) *PreviewOrderConfirmationResponse {
	return &PreviewOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) PreviewOrderConfirmation(req *PreviewOrderConfirmationRequest) (*PreviewOrderConfirmationResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	resp := PreviewOrderConfirmationResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH ORDER CONFIRMATION
type SearchOrderConfirmationRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchOrderConfirmationDTO
}

func (s *SaleService) NewSearchOrderConfirmationRequest(ctx context.Context, payload dto.SearchOrderConfirmationDTO) *SearchOrderConfirmationRequest {
	return &SearchOrderConfirmationRequest{Ctx: ctx, Payload: payload}
}

type SearchOrderConfirmationResponse struct {
	Payload    sale.OrderConfirmationSlice `json:"payload"`
	Pagination dto.PaginationDTO           `json:"pagination"`
}

func (s *SaleService) NewSearchOrderConfirmationResponse(payload sale.OrderConfirmationSlice) *SearchOrderConfirmationResponse {
	return &SearchOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) SearchOrderConfirmation(req *SearchOrderConfirmationRequest) (*SearchOrderConfirmationResponse, error) {
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

	res, totalCount, err := s.Repo.Sale.SearchOrderConfirmations(req.Ctx, tx, req.Payload)
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
	resp := SearchOrderConfirmationResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET ORDER CONFIRMATION
type GetOrderConfirmationRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *SaleService) NewGetOrderConfirmationRequest(ctx context.Context, id int) *GetOrderConfirmationRequest {
	return &GetOrderConfirmationRequest{Ctx: ctx, ID: id}
}

type GetOrderConfirmationResponse struct {
	Payload sale.OrderConfirmation `json:"payload"`
}

func (s *SaleService) NewGetOrderConfirmationResponse(payload sale.OrderConfirmation) *GetOrderConfirmationResponse {
	return &GetOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) GetOrderConfirmation(req *GetOrderConfirmationRequest) (*GetOrderConfirmationResponse, error) {
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

	res, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetOrderConfirmationResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE ORDER CONFIRMATION
type CreateOrderConfirmationRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.CreateOrderConfirmationDTO
}

func (s *SaleService) NewCreateOrderConfirmationRequest(ctx context.Context, payload dto.CreateOrderConfirmationDTO) *CreateOrderConfirmationRequest {
	return &CreateOrderConfirmationRequest{Ctx: ctx, Payload: payload}
}

type CreateOrderConfirmationResponse struct {
	Payload sale.OrderConfirmation `json:"payload"`
}

func (s *SaleService) NewCreateOrderConfirmationResponse(payload sale.OrderConfirmation) *CreateOrderConfirmationResponse {
	return &CreateOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) CreateOrderConfirmation(req *CreateOrderConfirmationRequest) (*CreateOrderConfirmationResponse, error) {
	/*
		1. Create Order Confirmation
		2. Create Order Confirmation Items
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

	if req.Payload.ToCreateFromSalesOrder {
		// get sales order
		salesOrder, err := s.Repo.Sale.GetSalesOrderByID(req.Ctx, tx, req.Payload.OrderConfirmation.SalesOrderID)
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
		var orderConfirmationBaseDocument base.BaseDocument
		copier.Copy(&orderConfirmationBaseDocument, &salesOrderBaseDocument)
		nextID, err := s.Repo.Base.GetNextEntryBaseDocumentID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		orderConfirmationBaseDocument.ID = nextID
		err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &orderConfirmationBaseDocument)
		if err != nil {
			return nil, err
		}

		// copy & create order confirmation
		var orderConfirmation sale.OrderConfirmation
		copier.Copy(&orderConfirmation, &req.Payload.OrderConfirmation)
		nextID, err = s.Repo.Sale.GetNextEntryOrderConfirmationID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		orderConfirmation.ID = nextID
		orderConfirmation.BaseDocumentID = orderConfirmationBaseDocument.ID
		if orderConfirmation.OrderConfirmationNumber == "" {
			orderConfirmation.OrderConfirmationNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateOrderConfirmation(req.Ctx, tx, &orderConfirmation)
		if err != nil {
			return nil, err
		}

		// copy & create base document items
		var bdItemMap = make(map[int]int)
		for _, item := range salesOrderBaseDocumentItems {
			var orderConfirmationBaseDocumentItem base.BaseDocumentItem
			copier.Copy(&orderConfirmationBaseDocumentItem, &item)
			nextID, err = s.Repo.Base.GetNextEntryBaseDocumentItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			orderConfirmationBaseDocumentItem.ID = nextID
			orderConfirmationBaseDocumentItem.BaseDocumentID = orderConfirmationBaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &orderConfirmationBaseDocumentItem)
			if err != nil {
				return nil, err
			}
			bdItemMap[item.ID] = orderConfirmationBaseDocumentItem.ID
		}

		// copy & create order confirmation items
		for _, item := range salesOrderItems {
			var orderConfirmationItem sale.OrderConfirmationItem
			copier.Copy(&orderConfirmationItem, &item)
			nextID, err = s.Repo.Sale.GetNextEntryOrderConfirmationItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			orderConfirmationItem.ID = nextID
			orderConfirmationItem.OrderConfirmationID = orderConfirmation.ID
			orderConfirmationItem.BaseDocumentItemID = bdItemMap[item.BaseDocumentItemID]
			err = s.Repo.Sale.CreateOrderConfirmationItem(req.Ctx, tx, &orderConfirmationItem)
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
		nextID, err := s.Repo.Base.GetNextEntryBaseDocumentID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.BaseDocument.ID = nextID
		err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
		if err != nil {
			return nil, err
		}

		// create order confirmation
		nextID, err = s.Repo.Sale.GetNextEntryOrderConfirmationID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.OrderConfirmation.ID = nextID
		req.Payload.OrderConfirmation.BaseDocumentID = req.Payload.BaseDocument.ID
		if req.Payload.OrderConfirmation.OrderConfirmationNumber == "" {
			req.Payload.OrderConfirmation.OrderConfirmationNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateOrderConfirmation(req.Ctx, tx, &req.Payload.OrderConfirmation)
		if err != nil {
			return nil, err
		}

		for _, item := range req.Payload.Items {
			// create base document items
			nextID, err = s.Repo.Base.GetNextEntryBaseDocumentItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.BaseDocumentItem.ID = nextID
			item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create order confirmation items
			nextID, err = s.Repo.Sale.GetNextEntryOrderConfirmationItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.OrderConfirmationItem.ID = nextID
			item.OrderConfirmationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.OrderConfirmationItem.OrderConfirmationID = req.Payload.OrderConfirmation.ID
			err = s.Repo.Sale.CreateOrderConfirmationItem(req.Ctx, tx, &item.OrderConfirmationItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// get created order confirmation
	oc, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.Payload.OrderConfirmation.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := CreateOrderConfirmationResponse{
		Payload: *oc,
	}

	return &resp, nil
}

// UPDATE ORDER CONFIRMATION
type UpdateOrderConfirmationRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.UpdateOrderConfirmationDTO
}

func (s *SaleService) NewUpdateOrderConfirmationRequest(ctx context.Context, payload dto.UpdateOrderConfirmationDTO) *UpdateOrderConfirmationRequest {
	return &UpdateOrderConfirmationRequest{Ctx: ctx, Payload: payload}
}

type UpdateOrderConfirmationResponse struct {
	Payload sale.OrderConfirmation `json:"payload"`
}

func (s *SaleService) NewUpdateOrderConfirmationResponse(payload sale.OrderConfirmation) *UpdateOrderConfirmationResponse {
	return &UpdateOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) UpdateOrderConfirmation(req *UpdateOrderConfirmationRequest) (*UpdateOrderConfirmationResponse, error) {
	/*
		1. Update Order Confirmation
		2. Update Order Confirmation Items
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

	currOc, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.Payload.OrderConfirmation.ID)
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

	// update order confirmation
	err = s.Repo.Sale.UpdateOrderConfirmation(req.Ctx, tx, &req.Payload.OrderConfirmation)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list but not in the new list
	currItems, err := s.Repo.Sale.GetOrderConfirmationItemsByOrderConfirmationID(req.Ctx, tx, currOc.ID)
	if err != nil {
		return nil, err
	}
	for _, currItem := range currItems {
		found := false
		for _, item := range req.Payload.Items {
			if currItem.ID == item.OrderConfirmationItem.ID {
				found = true
				break
			}
		}
		if !found {
			// delete order confirmation item
			err = s.Repo.Sale.DeleteOrderConfirmationItem(req.Ctx, tx, currItem)
			if err != nil {
				return nil, err
			}

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
		}
	}

	// create or update order confirmation items
	for _, item := range req.Payload.Items {
		// check if the item is new
		itemExists, err := s.Repo.Sale.OrderConfirmationItemExists(req.Ctx, tx, item.OrderConfirmationItem.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			// update base document items
			err = s.Repo.Base.UpdateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// update order confirmation items
			err = s.Repo.Sale.UpdateOrderConfirmationItem(req.Ctx, tx, &item.OrderConfirmationItem)
			if err != nil {
				return nil, err
			}
		} else {
			// create base document items
			nextID, err := s.Repo.Base.GetNextEntryBaseDocumentItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.BaseDocumentItem.ID = nextID
			item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create order confirmation items
			nextID, err = s.Repo.Sale.GetNextEntryOrderConfirmationItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.OrderConfirmationItem.ID = nextID
			item.OrderConfirmationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.OrderConfirmationItem.OrderConfirmationID = req.Payload.OrderConfirmation.ID
			err = s.Repo.Sale.CreateOrderConfirmationItem(req.Ctx, tx, &item.OrderConfirmationItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated order confirmation
	oc, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.Payload.OrderConfirmation.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := UpdateOrderConfirmationResponse{
		Payload: *oc,
	}

	return &resp, nil
}

// DELETE ORDER CONFIRMATION
type DeleteOrderConfirmationRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *SaleService) NewDeleteOrderConfirmationRequest(ctx context.Context, id int) *DeleteOrderConfirmationRequest {
	return &DeleteOrderConfirmationRequest{Ctx: ctx, ID: id}
}

type DeleteOrderConfirmationResponse struct {
	Payload bool `json:"payload"`
}

func (s *SaleService) NewDeleteOrderConfirmationResponse(payload bool) *DeleteOrderConfirmationResponse {
	return &DeleteOrderConfirmationResponse{Payload: payload}
}

func (s *SaleService) DeleteOrderConfirmation(req *DeleteOrderConfirmationRequest) (*DeleteOrderConfirmationResponse, error) {
	/*
		1. Delete Order Confirmation
		2. Delete Order Confirmation Items
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

	// get order confirmation
	OrderConfirmation, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete order confirmation
	err = s.Repo.Sale.DeleteOrderConfirmation(req.Ctx, tx, OrderConfirmation)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, OrderConfirmation.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.Sale.GetOrderConfirmationItemsByOrderConfirmationID(req.Ctx, tx, OrderConfirmation.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
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

		// delete order confirmation item
		err = s.Repo.Sale.DeleteOrderConfirmationItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := DeleteOrderConfirmationResponse{
		Payload: true,
	}

	return &resp, nil
}
