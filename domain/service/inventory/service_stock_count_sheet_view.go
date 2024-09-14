package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
)

// LIST GOODS ISSUE NOTE VIEW
type ListStockCountSheetViewRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListStockCountSheetViewRequest(ctx context.Context) *ListStockCountSheetViewRequest {
	return &ListStockCountSheetViewRequest{
		Ctx: ctx,
	}
}

type ListStockCountSheetViewResponse struct {
	Payload inventory.StockCountSheetViewSlice `json:"payload"`
}

func (s *InventoryService) NewListStockCountSheetViewResponse(payload inventory.StockCountSheetViewSlice) *ListStockCountSheetViewResponse {
	return &ListStockCountSheetViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListStockCountSheetView(req *ListStockCountSheetViewRequest) (*ListStockCountSheetViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllStockCountSheetViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListStockCountSheetViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH GOODS ISSUE NOTE VIEW
type SearchStockCountSheetViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchStockCountSheetDTO
}

func (s *InventoryService) NewSearchStockCountSheetViewRequest(ctx context.Context, payload dto.SearchStockCountSheetDTO) *SearchStockCountSheetViewRequest {
	return &SearchStockCountSheetViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchStockCountSheetViewResponse struct {
	Payload    inventory.StockCountSheetViewSlice `json:"payload"`
	Pagination dto.PaginationDTO                  `json:"pagination"`
}

func (s *InventoryService) NewSearchStockCountSheetViewResponse(payload inventory.StockCountSheetViewSlice) *SearchStockCountSheetViewResponse {
	return &SearchStockCountSheetViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchStockCountSheetView(req *SearchStockCountSheetViewRequest) (*SearchStockCountSheetViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Inventory.SearchStockCountSheetViews(req.Ctx, tx, req.Payload)
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
	resp := SearchStockCountSheetViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET GOODS ISSUE NOTE VIEW
type GetStockCountSheetViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetStockCountSheetViewRequest(ctx context.Context, id int) *GetStockCountSheetViewRequest {
	return &GetStockCountSheetViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetStockCountSheetViewResponse struct {
	Payload inventory.StockCountSheetView `json:"payload"`
}

func (s *InventoryService) NewGetStockCountSheetViewResponse(payload inventory.StockCountSheetView) *GetStockCountSheetViewResponse {
	return &GetStockCountSheetViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetStockCountSheetView(req *GetStockCountSheetViewRequest) (*GetStockCountSheetViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetStockCountSheetViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetStockCountSheetViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
