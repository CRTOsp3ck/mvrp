package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/data/repo"
	"mvrp/domain/dto"
)

// LIST ORDER CONFIRMATION VIEW
type ListOrderConfirmationViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *SaleService) NewListOrderConfirmationViewRequest(ctx context.Context) *ListOrderConfirmationViewRequest {
	return &ListOrderConfirmationViewRequest{
		Ctx: ctx,
	}
}

type ListOrderConfirmationViewResponse struct {
	Payload sale.OrderConfirmationViewSlice `json:"payload"`
}

func (s *SaleService) NewListOrderConfirmationViewResponse(payload sale.OrderConfirmationViewSlice) *ListOrderConfirmationViewResponse {
	return &ListOrderConfirmationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) ListOrderConfirmationView(req *ListOrderConfirmationViewRequest) (*ListOrderConfirmationViewResponse, error) {
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

	res, err := s.Repo.Sale.ListAllOrderConfirmationViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListOrderConfirmationViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH ORDER CONFIRMATION VIEW
type SearchOrderConfirmationViewRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchOrderConfirmationDTO
}

func (s *SaleService) NewSearchOrderConfirmationViewRequest(ctx context.Context, payload dto.SearchOrderConfirmationDTO) *SearchOrderConfirmationViewRequest {
	return &SearchOrderConfirmationViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchOrderConfirmationViewResponse struct {
	Payload    sale.OrderConfirmationViewSlice `json:"payload"`
	Pagination dto.PaginationDTO               `json:"pagination"`
}

func (s *SaleService) NewSearchOrderConfirmationViewResponse(payload sale.OrderConfirmationViewSlice) *SearchOrderConfirmationViewResponse {
	return &SearchOrderConfirmationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) SearchOrderConfirmationView(req *SearchOrderConfirmationViewRequest) (*SearchOrderConfirmationViewResponse, error) {
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

	res, totalCount, err := s.Repo.Sale.SearchOrderConfirmationViews(req.Ctx, tx, req.Payload)
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
	resp := SearchOrderConfirmationViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET ORDER CONFIRMATION VIEW
type GetOrderConfirmationViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *SaleService) NewGetOrderConfirmationViewRequest(ctx context.Context, id int) *GetOrderConfirmationViewRequest {
	return &GetOrderConfirmationViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetOrderConfirmationViewResponse struct {
	Payload sale.OrderConfirmationView `json:"payload"`
}

func (s *SaleService) NewGetOrderConfirmationViewResponse(payload sale.OrderConfirmationView) *GetOrderConfirmationViewResponse {
	return &GetOrderConfirmationViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) GetOrderConfirmationView(req *GetOrderConfirmationViewRequest) (*GetOrderConfirmationViewResponse, error) {
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

	res, err := s.Repo.Sale.GetOrderConfirmationViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetOrderConfirmationViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
