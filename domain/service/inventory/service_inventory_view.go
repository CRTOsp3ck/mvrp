package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
)

// LIST INVENTORY
type ListInventoryViewRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListInventoryViewRequest(ctx context.Context) *ListInventoryViewRequest {
	return &ListInventoryViewRequest{
		Ctx: ctx,
	}
}

type ListInventoryViewResponse struct {
	Payload inventory.InventoryViewSlice `json:"payload"`
}

func (s *InventoryService) NewListInventoryViewResponse(payload inventory.InventoryViewSlice) *ListInventoryViewResponse {
	return &ListInventoryViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListInventoryView(req *ListInventoryViewRequest) (*ListInventoryViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllInventoryViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListInventoryViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH INVENTORY
type SearchInventoryViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchInventoryDTO
}

func (s *InventoryService) NewSearchInventoryViewRequest(ctx context.Context, payload dto.SearchInventoryDTO) *SearchInventoryViewRequest {
	return &SearchInventoryViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchInventoryViewResponse struct {
	Payload    inventory.InventoryViewSlice `json:"payload"`
	Pagination dto.PaginationDTO            `json:"pagination"`
}

func (s *InventoryService) NewSearchInventoryViewResponse(payload inventory.InventoryViewSlice) *SearchInventoryViewResponse {
	return &SearchInventoryViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchInventoryView(req *SearchInventoryViewRequest) (*SearchInventoryViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Inventory.SearchInventoryViews(req.Ctx, tx, req.Payload)
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
	resp := SearchInventoryViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET INVENTORY
type GetInventoryViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetInventoryViewRequest(ctx context.Context, id int) *GetInventoryViewRequest {
	return &GetInventoryViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetInventoryViewResponse struct {
	Payload inventory.InventoryView `json:"payload"`
}

func (s *InventoryService) NewGetInventoryViewResponse(payload inventory.InventoryView) *GetInventoryViewResponse {
	return &GetInventoryViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetInventoryView(req *GetInventoryViewRequest) (*GetInventoryViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetInventoryViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetInventoryViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
