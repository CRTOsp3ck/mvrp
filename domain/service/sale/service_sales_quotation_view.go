package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
)

// LIST SALES QUOTATION VIEW
type ListSalesQuotationViewRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListSalesQuotationViewRequest(ctx context.Context) *ListSalesQuotationViewRequest {
	return &ListSalesQuotationViewRequest{
		Ctx: ctx,
	}
}

type ListSalesQuotationViewResponse struct {
	Payload sale.SalesQuotationViewSlice `json:"payload"`
}

func (s *SaleService) NewListSalesQuotationViewResponse(payload sale.SalesQuotationViewSlice) *ListSalesQuotationViewResponse {
	return &ListSalesQuotationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) ListSalesQuotationView(req *ListSalesQuotationViewRequest) (*ListSalesQuotationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllSalesQuotationViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListSalesQuotationViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH SALES QUOTATION VIEW
type SearchSalesQuotationViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchSalesQuotationDTO
}

func (s *SaleService) NewSearchSalesQuotationViewRequest(ctx context.Context, payload dto.SearchSalesQuotationDTO) *SearchSalesQuotationViewRequest {
	return &SearchSalesQuotationViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchSalesQuotationViewResponse struct {
	Payload    sale.SalesQuotationViewSlice `json:"payload"`
	Pagination dto.PaginationDTO            `json:"pagination"`
}

func (s *SaleService) NewSearchSalesQuotationViewResponse(payload sale.SalesQuotationViewSlice) *SearchSalesQuotationViewResponse {
	return &SearchSalesQuotationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) SearchSalesQuotationView(req *SearchSalesQuotationViewRequest) (*SearchSalesQuotationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Sale.SearchSalesQuotationViews(req.Ctx, tx, req.Payload)
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
	resp := SearchSalesQuotationViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET SALES QUOTATION VIEW
type GetSalesQuotationViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetSalesQuotationViewRequest(ctx context.Context, id int) *GetSalesQuotationViewRequest {
	return &GetSalesQuotationViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetSalesQuotationViewResponse struct {
	Payload sale.SalesQuotationView `json:"payload"`
}

func (s *SaleService) NewGetSalesQuotationViewResponse(payload sale.SalesQuotationView) *GetSalesQuotationViewResponse {
	return &GetSalesQuotationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) GetSalesQuotationView(req *GetSalesQuotationViewRequest) (*GetSalesQuotationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetSalesQuotationViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetSalesQuotationViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
