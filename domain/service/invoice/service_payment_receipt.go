package invoice

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST PAYMENT RECEIPT
type ListPaymentReceiptRequest struct {
	Ctx context.Context
}

func (s *InvoiceService) NewListPaymentReceiptRequest(ctx context.Context) *ListPaymentReceiptRequest {
	return &ListPaymentReceiptRequest{Ctx: ctx}
}

type ListPaymentReceiptResponse struct {
	Payload invoice.PaymentReceiptSlice `json:"payload"`
}

func (s *InvoiceService) NewListPaymentReceiptResponse(payload invoice.PaymentReceiptSlice) *ListPaymentReceiptResponse {
	return &ListPaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) ListPaymentReceipt(req *ListPaymentReceiptRequest) (*ListPaymentReceiptResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.ListAllPaymentReceipts(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListPaymentReceiptResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW PAYMENT RECEIPT
type PreviewPaymentReceiptRequest struct {
	Ctx     context.Context
	Payload dto.CreatePaymentReceiptDTO
}

func (s *InvoiceService) NewPreviewPaymentReceiptRequest(ctx context.Context, payload dto.CreatePaymentReceiptDTO) *PreviewPaymentReceiptRequest {
	return &PreviewPaymentReceiptRequest{Ctx: ctx, Payload: payload}
}

type PreviewPaymentReceiptResponse struct {
	Payload dto.CreatePaymentReceiptDTO `json:"payload"`
}

func (s *InvoiceService) NewPreviewPaymentReceiptResponse(payload dto.CreatePaymentReceiptDTO) *PreviewPaymentReceiptResponse {
	return &PreviewPaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) PreviewPaymentReceipt(req *PreviewPaymentReceiptRequest) (*PreviewPaymentReceiptResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}
	resp := PreviewPaymentReceiptResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH PAYMENT RECEIPT
type SearchPaymentReceiptRequest struct {
	Ctx     context.Context
	Payload dto.SearchPaymentReceiptDTO
}

func (s *InvoiceService) NewSearchPaymentReceiptRequest(ctx context.Context, payload dto.SearchPaymentReceiptDTO) *SearchPaymentReceiptRequest {
	return &SearchPaymentReceiptRequest{Ctx: ctx, Payload: payload}
}

type SearchPaymentReceiptResponse struct {
	Payload    invoice.PaymentReceiptSlice `json:"payload"`
	Pagination dto.PaginationDTO           `json:"pagination"`
}

func (s *InvoiceService) NewSearchPaymentReceiptResponse(payload invoice.PaymentReceiptSlice) *SearchPaymentReceiptResponse {
	return &SearchPaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) SearchPaymentReceipt(req *SearchPaymentReceiptRequest) (*SearchPaymentReceiptResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.SearchPaymentReceipts(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Invoice.GetPaymentReceiptTotalCount(req.Ctx, tx)
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
	resp := SearchPaymentReceiptResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET PAYMENT RECEIPT
type GetPaymentReceiptRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewGetPaymentReceiptRequest(ctx context.Context, id int) *GetPaymentReceiptRequest {
	return &GetPaymentReceiptRequest{Ctx: ctx, ID: id}
}

type GetPaymentReceiptResponse struct {
	Payload invoice.PaymentReceipt `json:"payload"`
}

func (s *InvoiceService) NewGetPaymentReceiptResponse(payload invoice.PaymentReceipt) *GetPaymentReceiptResponse {
	return &GetPaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) GetPaymentReceipt(req *GetPaymentReceiptRequest) (*GetPaymentReceiptResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := GetPaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE PAYMENT RECEIPT
type CreatePaymentReceiptRequest struct {
	Ctx     context.Context
	Payload dto.CreatePaymentReceiptDTO
}

func (s *InvoiceService) NewCreatePaymentReceiptRequest(ctx context.Context, payload dto.CreatePaymentReceiptDTO) *CreatePaymentReceiptRequest {
	return &CreatePaymentReceiptRequest{Ctx: ctx, Payload: payload}
}

type CreatePaymentReceiptResponse struct {
	Payload invoice.PaymentReceipt `json:"payload"`
}

func (s *InvoiceService) NewCreatePaymentReceiptResponse(payload invoice.PaymentReceipt) *CreatePaymentReceiptResponse {
	return &CreatePaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) CreatePaymentReceipt(req *CreatePaymentReceiptRequest) (*CreatePaymentReceiptResponse, error) {
	/*
		1. Create Base Document
		2. Create PaymentReceipt
		3. Create Base Document Items
		4. Create PaymentReceipt Items
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

	// create payment receipt
	req.Payload.PaymentReceipt.BaseDocumentID = req.Payload.BaseDocument.ID
	if req.Payload.PaymentReceipt.PaymentReceiptNumber == "" {
		nextID, err := s.Repo.Invoice.GetNextEntryPaymentReceiptID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.PaymentReceipt.PaymentReceiptNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Invoice.CreatePaymentReceipt(req.Ctx, tx, &req.Payload.PaymentReceipt)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Payload.Items {
		// create base document item
		item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
		err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
		if err != nil {
			return nil, err
		}

		// create payment receipt item
		item.PaymentReceiptItem.PaymentReceiptID = req.Payload.PaymentReceipt.ID
		item.PaymentReceiptItem.BaseDocumentItemID = item.BaseDocumentItem.ID
		err = s.Repo.Invoice.CreatePaymentReceiptItem(req.Ctx, tx, &item.PaymentReceiptItem)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.Payload.PaymentReceipt.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreatePaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE PAYMENT RECEIPT
type UpdatePaymentReceiptRequest struct {
	Ctx     context.Context
	Payload dto.UpdatePaymentReceiptDTO
}

func (s *InvoiceService) NewUpdatePaymentReceiptRequest(ctx context.Context, payload dto.UpdatePaymentReceiptDTO) *UpdatePaymentReceiptRequest {
	return &UpdatePaymentReceiptRequest{Ctx: ctx, Payload: payload}
}

type UpdatePaymentReceiptResponse struct {
	Payload invoice.PaymentReceipt `json:"payload"`
}

func (s *InvoiceService) NewUpdatePaymentReceiptResponse(payload invoice.PaymentReceipt) *UpdatePaymentReceiptResponse {
	return &UpdatePaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) UpdatePaymentReceipt(req *UpdatePaymentReceiptRequest) (*UpdatePaymentReceiptResponse, error) {
	/*
		1. Update Base Document
		2. Update PaymentReceipt
		3. Update Base Document Items
		4. Update PaymentReceipt Items
	*/
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currInv, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.Payload.PaymentReceipt.ID)
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

	// update payment receipt
	err = s.Repo.Invoice.UpdatePaymentReceipt(req.Ctx, tx, &req.Payload.PaymentReceipt)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	for _, item := range currInv.R.PaymentReceiptItems {
		found := false
		for _, newItem := range req.Payload.Items {
			if item.ID == newItem.PaymentReceiptItem.ID {
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

			// delete payment receipt item
			err = s.Repo.Invoice.DeletePaymentReceiptItem(req.Ctx, tx, item)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update payment receipt items
	for _, item := range req.Payload.Items {
		// check if the item is new or existing
		itemExists, err := s.Repo.Invoice.PaymentReceiptItemExists(req.Ctx, tx, item.PaymentReceiptItem.ID)
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

			// update payment receipt items
			err = s.Repo.Invoice.UpdatePaymentReceiptItem(req.Ctx, tx, &item.PaymentReceiptItem)
			if err != nil {
				return nil, err
			}
		} else {
			// create base document items
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create payment receipt items
			item.PaymentReceiptItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.PaymentReceiptItem.PaymentReceiptID = req.Payload.PaymentReceipt.ID
			err = s.Repo.Invoice.CreatePaymentReceiptItem(req.Ctx, tx, &item.PaymentReceiptItem)
			if err != nil {
				return nil, err
			}
		}
	}

	res, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.Payload.PaymentReceipt.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdatePaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE PAYMENT RECEIPT
type DeletePaymentReceiptRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewDeletePaymentReceiptRequest(ctx context.Context, id int) *DeletePaymentReceiptRequest {
	return &DeletePaymentReceiptRequest{Ctx: ctx, ID: id}
}

type DeletePaymentReceiptResponse struct {
	Payload bool `json:"payload"`
}

func (s *InvoiceService) NewDeletePaymentReceiptResponse(payload bool) *DeletePaymentReceiptResponse {
	return &DeletePaymentReceiptResponse{Payload: payload}
}

func (s *InvoiceService) DeletePaymentReceipt(req *DeletePaymentReceiptRequest) (*DeletePaymentReceiptResponse, error) {
	/*
		1. Delete Base Document
		2. Delete PaymentReceipt
		3. Delete Base Document Items
		4. Delete PaymentReceipt Items
	*/
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get payment receipt
	inv, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete payment receipt
	err = s.Repo.Invoice.DeletePaymentReceipt(req.Ctx, tx, inv)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, inv.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	for _, item := range inv.R.PaymentReceiptItems {
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

		// delete payment receipt item
		err = s.Repo.Invoice.DeletePaymentReceiptItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeletePaymentReceiptResponse{
		Payload: true,
	}
	return &resp, nil
}
