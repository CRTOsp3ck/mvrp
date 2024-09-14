package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
)

// LIST GOODS ISSUE NOTE VIEW
type ListGoodsIssueNoteViewRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListGoodsIssueNoteViewRequest(ctx context.Context) *ListGoodsIssueNoteViewRequest {
	return &ListGoodsIssueNoteViewRequest{
		Ctx: ctx,
	}
}

type ListGoodsIssueNoteViewResponse struct {
	Payload inventory.GoodsIssueNoteViewSlice `json:"payload"`
}

func (s *InventoryService) NewListGoodsIssueNoteViewResponse(payload inventory.GoodsIssueNoteViewSlice) *ListGoodsIssueNoteViewResponse {
	return &ListGoodsIssueNoteViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListGoodsIssueNoteView(req *ListGoodsIssueNoteViewRequest) (*ListGoodsIssueNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllGoodsIssueNoteViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListGoodsIssueNoteViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH GOODS ISSUE NOTE VIEW
type SearchGoodsIssueNoteViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchGoodsIssueNoteDTO
}

func (s *InventoryService) NewSearchGoodsIssueNoteViewRequest(ctx context.Context, payload dto.SearchGoodsIssueNoteDTO) *SearchGoodsIssueNoteViewRequest {
	return &SearchGoodsIssueNoteViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchGoodsIssueNoteViewResponse struct {
	Payload    inventory.GoodsIssueNoteViewSlice `json:"payload"`
	Pagination dto.PaginationDTO                 `json:"pagination"`
}

func (s *InventoryService) NewSearchGoodsIssueNoteViewResponse(payload inventory.GoodsIssueNoteViewSlice) *SearchGoodsIssueNoteViewResponse {
	return &SearchGoodsIssueNoteViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchGoodsIssueNoteView(req *SearchGoodsIssueNoteViewRequest) (*SearchGoodsIssueNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Inventory.SearchGoodsIssueNoteViews(req.Ctx, tx, req.Payload)
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
	resp := SearchGoodsIssueNoteViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET GOODS ISSUE NOTE VIEW
type GetGoodsIssueNoteViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetGoodsIssueNoteViewRequest(ctx context.Context, id int) *GetGoodsIssueNoteViewRequest {
	return &GetGoodsIssueNoteViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetGoodsIssueNoteViewResponse struct {
	Payload inventory.GoodsIssueNoteView `json:"payload"`
}

func (s *InventoryService) NewGetGoodsIssueNoteViewResponse(payload inventory.GoodsIssueNoteView) *GetGoodsIssueNoteViewResponse {
	return &GetGoodsIssueNoteViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetGoodsIssueNoteView(req *GetGoodsIssueNoteViewRequest) (*GetGoodsIssueNoteViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetGoodsIssueNoteViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetGoodsIssueNoteViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
