package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
)

// LIST DELIVERY NOTE
type ListCreditNoteViewRequest struct {
	Ctx context.Context
}

func (s *InvoiceService) NewListCreditNoteViewRequest(ctx context.Context) *ListCreditNoteViewRequest {
	return &ListCreditNoteViewRequest{Ctx: ctx}
}

type ListCreditNoteViewResponse struct {
	Payload invoice.CreditNoteViewSlice `json:"payload"`
}

func (s *InvoiceService) NewListCreditNoteViewResponse(payload invoice.CreditNoteViewSlice) *ListCreditNoteViewResponse {
	return &ListCreditNoteViewResponse{Payload: payload}
}

func (s *InvoiceService) ListCreditNoteView(req *ListCreditNoteViewRequest) (*ListCreditNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.ListAllCreditNoteViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListCreditNoteViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE
type SearchCreditNoteViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchCreditNoteDTO
}

func (s *InvoiceService) NewSearchCreditNoteViewRequest(ctx context.Context, payload dto.SearchCreditNoteDTO) *SearchCreditNoteViewRequest {
	return &SearchCreditNoteViewRequest{Ctx: ctx, Payload: payload}
}

type SearchCreditNoteViewResponse struct {
	Payload    invoice.CreditNoteViewSlice `json:"payload"`
	Pagination dto.PaginationDTO           `json:"pagination"`
}

func (s *InvoiceService) NewSearchCreditNoteViewResponse(payload invoice.CreditNoteViewSlice) *SearchCreditNoteViewResponse {
	return &SearchCreditNoteViewResponse{Payload: payload}
}

func (s *InvoiceService) SearchCreditNoteView(req *SearchCreditNoteViewRequest) (*SearchCreditNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Invoice.SearchCreditNoteViews(req.Ctx, tx, req.Payload)
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
	resp := SearchCreditNoteViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetCreditNoteViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewGetCreditNoteViewRequest(ctx context.Context, id int) *GetCreditNoteViewRequest {
	return &GetCreditNoteViewRequest{Ctx: ctx, ID: id}
}

type GetCreditNoteViewResponse struct {
	Payload invoice.CreditNoteView `json:"payload"`
}

func (s *InvoiceService) NewGetCreditNoteViewResponse(payload invoice.CreditNoteView) *GetCreditNoteViewResponse {
	return &GetCreditNoteViewResponse{Payload: payload}
}

func (s *InvoiceService) GetCreditNoteView(req *GetCreditNoteViewRequest) (*GetCreditNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.GetCreditNoteViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetCreditNoteViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
