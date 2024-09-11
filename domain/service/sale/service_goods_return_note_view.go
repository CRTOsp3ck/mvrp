package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
)

// LIST GOODS RETURN NOTE VIEW
type ListGoodsReturnNoteViewRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListGoodsReturnNoteViewRequest(ctx context.Context) *ListGoodsReturnNoteViewRequest {
	return &ListGoodsReturnNoteViewRequest{
		Ctx: ctx,
	}
}

type ListGoodsReturnNoteViewResponse struct {
	Payload sale.GoodsReturnNoteViewSlice `json:"payload"`
}

func (s *SaleService) NewListGoodsReturnNoteViewResponse(payload sale.GoodsReturnNoteViewSlice) *ListGoodsReturnNoteViewResponse {
	return &ListGoodsReturnNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) ListGoodsReturnNoteView(req *ListGoodsReturnNoteViewRequest) (*ListGoodsReturnNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllGoodsReturnNoteViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListGoodsReturnNoteViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH GOODS RETURN NOTE VIEW
type SearchGoodsReturnNoteViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchGoodsReturnNoteDTO
}

func (s *SaleService) NewSearchGoodsReturnNoteViewRequest(ctx context.Context, payload dto.SearchGoodsReturnNoteDTO) *SearchGoodsReturnNoteViewRequest {
	return &SearchGoodsReturnNoteViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchGoodsReturnNoteViewResponse struct {
	Payload    sale.GoodsReturnNoteViewSlice `json:"payload"`
	Pagination dto.PaginationDTO             `json:"pagination"`
}

func (s *SaleService) NewSearchGoodsReturnNoteViewResponse(payload sale.GoodsReturnNoteViewSlice) *SearchGoodsReturnNoteViewResponse {
	return &SearchGoodsReturnNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) SearchGoodsReturnNoteView(req *SearchGoodsReturnNoteViewRequest) (*SearchGoodsReturnNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchGoodsReturnNoteViews(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Sale.GetGoodsReturnNoteTotalCount(req.Ctx, tx)
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
	resp := SearchGoodsReturnNoteViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET GOODS RETURN NOTE VIEW
type GetGoodsReturnNoteViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetGoodsReturnNoteViewRequest(ctx context.Context, id int) *GetGoodsReturnNoteViewRequest {
	return &GetGoodsReturnNoteViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetGoodsReturnNoteViewResponse struct {
	Payload sale.GoodsReturnNoteView `json:"payload"`
}

func (s *SaleService) NewGetGoodsReturnNoteViewResponse(payload sale.GoodsReturnNoteView) *GetGoodsReturnNoteViewResponse {
	return &GetGoodsReturnNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) GetGoodsReturnNoteView(req *GetGoodsReturnNoteViewRequest) (*GetGoodsReturnNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetGoodsReturnNoteViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetGoodsReturnNoteViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
