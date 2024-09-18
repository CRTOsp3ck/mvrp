package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/data/repo"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"
)

// LIST PAYMENT RECEIPT
type ListPaymentReceiptRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
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

	res, err := s.Repo.Invoice.ListAllPaymentReceipts(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
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
	pris := make([]*invoice.PaymentReceiptItem, 0)
	for _, item := range req.Payload.Items {
		pris = append(pris, &item.PaymentReceiptItem)
	}
	err := proc.ProcessPaymentReceiptAmounts(&req.Payload.PaymentReceipt, pris)
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
	RepoTx  *repo.RepoTx
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

	res, totalCount, err := s.Repo.Invoice.SearchPaymentReceipts(req.Ctx, tx, req.Payload)
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
	resp := SearchPaymentReceiptResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET PAYMENT RECEIPT
type GetPaymentReceiptRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
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

	res, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetPaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE PAYMENT RECEIPT
type CreatePaymentReceiptRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
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
		1. Create PaymentReceipt
		2. Create PaymentReceipt Items
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

	// create payment receipt
	nextID, err := s.Repo.Invoice.GetNextEntryPaymentReceiptID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.PaymentReceipt.ID = nextID
	if req.Payload.PaymentReceipt.PaymentReceiptNumber == "" {
		req.Payload.PaymentReceipt.PaymentReceiptNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Invoice.CreatePaymentReceipt(req.Ctx, tx, &req.Payload.PaymentReceipt)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Payload.Items {
		// create payment receipt item
		nextID, err = s.Repo.Invoice.GetNextEntryPaymentReceiptItemID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		item.PaymentReceiptItem.ID = nextID
		item.PaymentReceiptItem.PaymentReceiptID = req.Payload.PaymentReceipt.ID
		err = s.Repo.Invoice.CreatePaymentReceiptItem(req.Ctx, tx, &item.PaymentReceiptItem)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.Repo.Invoice.GetPaymentReceiptByID(req.Ctx, tx, req.Payload.PaymentReceipt.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := CreatePaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE PAYMENT RECEIPT
type UpdatePaymentReceiptRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
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
		1. Update PaymentReceipt
		2. Update PaymentReceipt Items
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

	// preprocess amounts
	pris := make([]*invoice.PaymentReceiptItem, 0)
	for _, item := range req.Payload.Items {
		pris = append(pris, &item.PaymentReceiptItem)
	}
	err = proc.ProcessPaymentReceiptAmounts(&req.Payload.PaymentReceipt, pris)
	if err != nil {
		return nil, err
	}

	// update payment receipt
	err = s.Repo.Invoice.UpdatePaymentReceipt(req.Ctx, tx, &req.Payload.PaymentReceipt)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	items, err := s.Repo.Invoice.GetPaymentReceiptItemsByPaymentReceiptID(req.Ctx, tx, req.Payload.PaymentReceipt.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		found := false
		for _, newItem := range req.Payload.Items {
			if item.ID == newItem.PaymentReceiptItem.ID {
				found = true
				break
			}
		}
		if !found {
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
			// preprocess amounts
			err = proc.ProcessPaymentReceiptItemAmounts(&item.PaymentReceiptItem)
			if err != nil {
				return nil, err
			}

			// update payment receipt items
			err = s.Repo.Invoice.UpdatePaymentReceiptItem(req.Ctx, tx, &item.PaymentReceiptItem)
			if err != nil {
				return nil, err
			}
		} else {
			// create payment receipt items
			nextID, err := s.Repo.Invoice.GetNextEntryPaymentReceiptItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.PaymentReceiptItem.ID = nextID
			item.PaymentReceiptItem.PaymentReceiptID = req.Payload.PaymentReceipt.ID

			// preprocess amounts
			err = proc.ProcessPaymentReceiptItemAmounts(&item.PaymentReceiptItem)
			if err != nil {
				return nil, err
			}

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

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}
	resp := UpdatePaymentReceiptResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE PAYMENT RECEIPT
type DeletePaymentReceiptRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
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

	items, err := s.Repo.Invoice.GetPaymentReceiptItemsByPaymentReceiptID(req.Ctx, tx, inv.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		// delete payment receipt item
		err = s.Repo.Invoice.DeletePaymentReceiptItem(req.Ctx, tx, item)
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

	resp := DeletePaymentReceiptResponse{
		Payload: true,
	}
	return &resp, nil
}
