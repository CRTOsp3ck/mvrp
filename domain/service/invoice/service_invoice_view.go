package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/data/repo"
	"mvrp/domain/dto"
)

// LIST DELIVERY NOTE
type ListInvoiceViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *InvoiceService) NewListInvoiceViewRequest(ctx context.Context) *ListInvoiceViewRequest {
	return &ListInvoiceViewRequest{Ctx: ctx}
}

type ListInvoiceViewResponse struct {
	Payload invoice.InvoiceViewSlice `json:"payload"`
}

func (s *InvoiceService) NewListInvoiceViewResponse(payload invoice.InvoiceViewSlice) *ListInvoiceViewResponse {
	return &ListInvoiceViewResponse{Payload: payload}
}

func (s *InvoiceService) ListInvoiceView(req *ListInvoiceViewRequest) (*ListInvoiceViewResponse, error) {
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

	res, err := s.Repo.Invoice.ListAllInvoiceViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListInvoiceViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE
type SearchInvoiceViewRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchInvoiceDTO
}

func (s *InvoiceService) NewSearchInvoiceViewRequest(ctx context.Context, payload dto.SearchInvoiceDTO) *SearchInvoiceViewRequest {
	return &SearchInvoiceViewRequest{Ctx: ctx, Payload: payload}
}

type SearchInvoiceViewResponse struct {
	Payload    invoice.InvoiceViewSlice `json:"payload"`
	Pagination dto.PaginationDTO        `json:"pagination"`
}

func (s *InvoiceService) NewSearchInvoiceViewResponse(payload invoice.InvoiceViewSlice) *SearchInvoiceViewResponse {
	return &SearchInvoiceViewResponse{Payload: payload}
}

func (s *InvoiceService) SearchInvoiceView(req *SearchInvoiceViewRequest) (*SearchInvoiceViewResponse, error) {
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

	res, totalCount, err := s.Repo.Invoice.SearchInvoiceViews(req.Ctx, tx, req.Payload)
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
	resp := SearchInvoiceViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetInvoiceViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *InvoiceService) NewGetInvoiceViewRequest(ctx context.Context, id int) *GetInvoiceViewRequest {
	return &GetInvoiceViewRequest{Ctx: ctx, ID: id}
}

type GetInvoiceViewResponse struct {
	Payload invoice.InvoiceView `json:"payload"`
}

func (s *InvoiceService) NewGetInvoiceViewResponse(payload invoice.InvoiceView) *GetInvoiceViewResponse {
	return &GetInvoiceViewResponse{Payload: payload}
}

func (s *InvoiceService) GetInvoiceView(req *GetInvoiceViewRequest) (*GetInvoiceViewResponse, error) {
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

	res, err := s.Repo.Invoice.GetInvoiceViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetInvoiceViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
