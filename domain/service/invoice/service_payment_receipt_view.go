package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/data/repo"
	"mvrp/domain/dto"
)

// LIST DELIVERY NOTE
type ListPaymentReceiptViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *InvoiceService) NewListPaymentReceiptViewRequest(ctx context.Context) *ListPaymentReceiptViewRequest {
	return &ListPaymentReceiptViewRequest{Ctx: ctx}
}

type ListPaymentReceiptViewResponse struct {
	Payload invoice.PaymentReceiptViewSlice `json:"payload"`
}

func (s *InvoiceService) NewListPaymentReceiptViewResponse(payload invoice.PaymentReceiptViewSlice) *ListPaymentReceiptViewResponse {
	return &ListPaymentReceiptViewResponse{Payload: payload}
}

func (s *InvoiceService) ListPaymentReceiptView(req *ListPaymentReceiptViewRequest) (*ListPaymentReceiptViewResponse, error) {
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

	res, err := s.Repo.Invoice.ListAllPaymentReceiptViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListPaymentReceiptViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE
type SearchPaymentReceiptViewRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchPaymentReceiptDTO
}

func (s *InvoiceService) NewSearchPaymentReceiptViewRequest(ctx context.Context, payload dto.SearchPaymentReceiptDTO) *SearchPaymentReceiptViewRequest {
	return &SearchPaymentReceiptViewRequest{Ctx: ctx, Payload: payload}
}

type SearchPaymentReceiptViewResponse struct {
	Payload    invoice.PaymentReceiptViewSlice `json:"payload"`
	Pagination dto.PaginationDTO               `json:"pagination"`
}

func (s *InvoiceService) NewSearchPaymentReceiptViewResponse(payload invoice.PaymentReceiptViewSlice) *SearchPaymentReceiptViewResponse {
	return &SearchPaymentReceiptViewResponse{Payload: payload}
}

func (s *InvoiceService) SearchPaymentReceiptView(req *SearchPaymentReceiptViewRequest) (*SearchPaymentReceiptViewResponse, error) {
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

	res, totalCount, err := s.Repo.Invoice.SearchPaymentReceiptViews(req.Ctx, tx, req.Payload)
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
	resp := SearchPaymentReceiptViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetPaymentReceiptViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *InvoiceService) NewGetPaymentReceiptViewRequest(ctx context.Context, id int) *GetPaymentReceiptViewRequest {
	return &GetPaymentReceiptViewRequest{Ctx: ctx, ID: id}
}

type GetPaymentReceiptViewResponse struct {
	Payload invoice.PaymentReceiptView `json:"payload"`
}

func (s *InvoiceService) NewGetPaymentReceiptViewResponse(payload invoice.PaymentReceiptView) *GetPaymentReceiptViewResponse {
	return &GetPaymentReceiptViewResponse{Payload: payload}
}

func (s *InvoiceService) GetPaymentReceiptView(req *GetPaymentReceiptViewRequest) (*GetPaymentReceiptViewResponse, error) {
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

	res, err := s.Repo.Invoice.GetPaymentReceiptViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetPaymentReceiptViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
