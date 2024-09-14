package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
)

// LIST RETURN MERCHANDISE AUTHORIZATION VIEW
type ListReturnMerchandiseAuthorizationViewRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListReturnMerchandiseAuthorizationViewRequest(ctx context.Context) *ListReturnMerchandiseAuthorizationViewRequest {
	return &ListReturnMerchandiseAuthorizationViewRequest{
		Ctx: ctx,
	}
}

type ListReturnMerchandiseAuthorizationViewResponse struct {
	Payload inventory.ReturnMerchandiseAuthorizationViewSlice `json:"payload"`
}

func (s *InventoryService) NewListReturnMerchandiseAuthorizationViewResponse(payload inventory.ReturnMerchandiseAuthorizationViewSlice) *ListReturnMerchandiseAuthorizationViewResponse {
	return &ListReturnMerchandiseAuthorizationViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListReturnMerchandiseAuthorizationView(req *ListReturnMerchandiseAuthorizationViewRequest) (*ListReturnMerchandiseAuthorizationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllReturnMerchandiseAuthorizationViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListReturnMerchandiseAuthorizationViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH RETURN MERCHANDISE AUTHORIZATION VIEW
type SearchReturnMerchandiseAuthorizationViewRequest struct {
	Ctx     context.Context
	Payload dto.SearchReturnMerchandiseAuthorizationDTO
}

func (s *InventoryService) NewSearchReturnMerchandiseAuthorizationViewRequest(ctx context.Context, payload dto.SearchReturnMerchandiseAuthorizationDTO) *SearchReturnMerchandiseAuthorizationViewRequest {
	return &SearchReturnMerchandiseAuthorizationViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchReturnMerchandiseAuthorizationViewResponse struct {
	Payload    inventory.ReturnMerchandiseAuthorizationViewSlice `json:"payload"`
	Pagination dto.PaginationDTO                                 `json:"pagination"`
}

func (s *InventoryService) NewSearchReturnMerchandiseAuthorizationViewResponse(payload inventory.ReturnMerchandiseAuthorizationViewSlice) *SearchReturnMerchandiseAuthorizationViewResponse {
	return &SearchReturnMerchandiseAuthorizationViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchReturnMerchandiseAuthorizationView(req *SearchReturnMerchandiseAuthorizationViewRequest) (*SearchReturnMerchandiseAuthorizationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, totalCount, err := s.Repo.Inventory.SearchReturnMerchandiseAuthorizationViews(req.Ctx, tx, req.Payload)
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
	resp := SearchReturnMerchandiseAuthorizationViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET RETURN MERCHANDISE AUTHORIZATION VIEW
type GetReturnMerchandiseAuthorizationViewRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetReturnMerchandiseAuthorizationViewRequest(ctx context.Context, id int) *GetReturnMerchandiseAuthorizationViewRequest {
	return &GetReturnMerchandiseAuthorizationViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetReturnMerchandiseAuthorizationViewResponse struct {
	Payload inventory.ReturnMerchandiseAuthorizationView `json:"payload"`
}

func (s *InventoryService) NewGetReturnMerchandiseAuthorizationViewResponse(payload inventory.ReturnMerchandiseAuthorizationView) *GetReturnMerchandiseAuthorizationViewResponse {
	return &GetReturnMerchandiseAuthorizationViewResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetReturnMerchandiseAuthorizationView(req *GetReturnMerchandiseAuthorizationViewRequest) (*GetReturnMerchandiseAuthorizationViewResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetReturnMerchandiseAuthorizationViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
