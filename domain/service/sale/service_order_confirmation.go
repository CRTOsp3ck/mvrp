package sale

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/jinzhu/copier"
)

// LIST ORDER CONFIRMATION
type ListOrderConfirmationRequest struct {
	Ctx context.Context
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllOrderConfirmations(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchOrderConfirmations(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchOrderConfirmationResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET ORDER CONFIRMATION
type GetOrderConfirmationRequest struct {
	Ctx context.Context
	ID  int
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
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetOrderConfirmationResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE ORDER CONFIRMATION
type CreateOrderConfirmationRequest struct {
	Ctx     context.Context
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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
		orderConfirmationBaseDocument.ID = -1
		err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &orderConfirmationBaseDocument)
		if err != nil {
			return nil, err
		}

		// copy & create order confirmation
		var orderConfirmation sale.OrderConfirmation
		copier.Copy(&orderConfirmation, &req.Payload.OrderConfirmation)
		orderConfirmation.BaseDocumentID = orderConfirmationBaseDocument.ID
		if orderConfirmation.OrderConfirmationNumber == "" {
			nextID, err := s.Repo.Sale.GetNextEntryOrderConfirmationID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			orderConfirmation.OrderConfirmationNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateOrderConfirmation(req.Ctx, tx, &orderConfirmation)
		if err != nil {
			return nil, err
		}

		// copy & create base document items
		for _, item := range salesOrderBaseDocumentItems {
			var orderConfirmationBaseDocumentItem base.BaseDocumentItem
			copier.Copy(&orderConfirmationBaseDocumentItem, &item)
			orderConfirmationBaseDocumentItem.ID = -1
			orderConfirmationBaseDocumentItem.BaseDocumentID = orderConfirmationBaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &orderConfirmationBaseDocumentItem)
			if err != nil {
				return nil, err
			}
		}

		// copy & create order confirmation items
		for _, item := range salesOrderItems {
			var orderConfirmationItem sale.OrderConfirmationItem
			copier.Copy(&orderConfirmationItem, &item)
			orderConfirmationItem.ID = -1
			orderConfirmationItem.OrderConfirmationID = orderConfirmation.ID
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
		err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
		if err != nil {
			return nil, err
		}

		// create order confirmation
		req.Payload.OrderConfirmation.BaseDocumentID = req.Payload.BaseDocument.ID
		if req.Payload.OrderConfirmation.OrderConfirmationNumber == "" {
			nextID, err := s.Repo.Sale.GetNextEntryOrderConfirmationID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			req.Payload.OrderConfirmation.OrderConfirmationNumber = util.Util.Str.ToString(nextID)
		}
		err = s.Repo.Sale.CreateOrderConfirmation(req.Ctx, tx, &req.Payload.OrderConfirmation)
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

			// create order confirmation items
			item.OrderConfirmationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.OrderConfirmationItem.OrderConfirmationID = req.Payload.OrderConfirmation.ID
			err = s.Repo.Sale.CreateOrderConfirmationItem(req.Ctx, tx, &item.OrderConfirmationItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// get created order confirmation
	OrderConfirmation, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.Payload.OrderConfirmation.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateOrderConfirmationResponse{
		Payload: *OrderConfirmation,
	}

	return &resp, nil
}

// UPDATE ORDER CONFIRMATION
type UpdateOrderConfirmationRequest struct {
	Ctx     context.Context
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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
	for _, currItem := range currOc.R.OrderConfirmationItems {
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
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create order confirmation items
			item.OrderConfirmationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.OrderConfirmationItem.OrderConfirmationID = req.Payload.OrderConfirmation.ID
			err = s.Repo.Sale.CreateOrderConfirmationItem(req.Ctx, tx, &item.OrderConfirmationItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated order confirmation
	OrderConfirmation, err := s.Repo.Sale.GetOrderConfirmationByID(req.Ctx, tx, req.Payload.OrderConfirmation.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateOrderConfirmationResponse{
		Payload: *OrderConfirmation,
	}

	return &resp, nil
}

// DELETE ORDER CONFIRMATION
type DeleteOrderConfirmationRequest struct {
	Ctx context.Context
	ID  int
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

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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

	for _, item := range OrderConfirmation.R.OrderConfirmationItems {
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

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteOrderConfirmationResponse{
		Payload: true,
	}

	return &resp, nil
}
