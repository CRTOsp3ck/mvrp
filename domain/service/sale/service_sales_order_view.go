package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
)

// LIST SALES ORDER VIEW
type ListSalesOrderViewRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListSalesOrderViewRequest(ctx context.Context) *ListSalesOrderViewRequest {
	return &ListSalesOrderViewRequest{
		Ctx: ctx,
	}
}

type ListSalesOrderViewResponse struct {
	Payload sale.SalesOrderViewSlice `json:"payload"`
}

func (s *SaleService) NewListSalesOrderViewResponse(payload sale.SalesOrderViewSlice) *ListSalesOrderViewResponse {
	return &ListSalesOrderViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) ListSalesOrderView(req *ListSalesOrderViewRequest) (*ListSalesOrderViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllSalesOrderViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListSalesOrderViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH SALES ORDER VIEW
type SearchSalesOrderViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchSalesOrderDTO
}

func (s *SaleService) NewSearchSalesOrderViewRequest(ctx context.Context, payload dto.SearchSalesOrderDTO) *SearchSalesOrderViewRequest {
	return &SearchSalesOrderViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchSalesOrderViewResponse struct {
	Payload    sale.SalesOrderViewSlice `json:"payload"`
	Pagination dto.PaginationDTO        `json:"pagination"`
}

func (s *SaleService) NewSearchSalesOrderViewResponse(payload sale.SalesOrderViewSlice) *SearchSalesOrderViewResponse {
	return &SearchSalesOrderViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) SearchSalesOrderView(req *SearchSalesOrderViewRequest) (*SearchSalesOrderViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchSalesOrderViews(req.Ctx, tx, req.Payload)
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
	resp := SearchSalesOrderViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET SALES ORDER VIEW
type GetSalesOrderViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetSalesOrderViewRequest(ctx context.Context, id int) *GetSalesOrderViewRequest {
	return &GetSalesOrderViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetSalesOrderViewResponse struct {
	Payload sale.SalesOrderView `json:"payload"`
}

func (s *SaleService) NewGetSalesOrderViewResponse(payload sale.SalesOrderView) *GetSalesOrderViewResponse {
	return &GetSalesOrderViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) GetSalesOrderView(req *GetSalesOrderViewRequest) (*GetSalesOrderViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetSalesOrderViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetSalesOrderViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
