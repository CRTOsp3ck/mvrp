package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/data/repo"
	"mvrp/domain/dto"
)

// LIST DELIVERY NOTE VIEW
type ListDeliveryNoteViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *SaleService) NewListDeliveryNoteViewRequest(ctx context.Context) *ListDeliveryNoteViewRequest {
	return &ListDeliveryNoteViewRequest{
		Ctx: ctx,
	}
}

type ListDeliveryNoteViewResponse struct {
	Payload sale.DeliveryNoteViewSlice `json:"payload"`
}

func (s *SaleService) NewListDeliveryNoteViewResponse(payload sale.DeliveryNoteViewSlice) *ListDeliveryNoteViewResponse {
	return &ListDeliveryNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) ListDeliveryNoteView(req *ListDeliveryNoteViewRequest) (*ListDeliveryNoteViewResponse, error) {
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

	res, err := s.Repo.Sale.ListAllDeliveryNoteViews(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListDeliveryNoteViewResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE VIEW
type SearchDeliveryNoteViewRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchDeliveryNoteDTO
}

func (s *SaleService) NewSearchDeliveryNoteViewRequest(ctx context.Context, payload dto.SearchDeliveryNoteDTO) *SearchDeliveryNoteViewRequest {
	return &SearchDeliveryNoteViewRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchDeliveryNoteViewResponse struct {
	Payload    sale.DeliveryNoteViewSlice `json:"payload"`
	Pagination dto.PaginationDTO          `json:"pagination"`
}

func (s *SaleService) NewSearchDeliveryNoteViewResponse(payload sale.DeliveryNoteViewSlice) *SearchDeliveryNoteViewResponse {
	return &SearchDeliveryNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) SearchDeliveryNoteView(req *SearchDeliveryNoteViewRequest) (*SearchDeliveryNoteViewResponse, error) {
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

	res, totalCount, err := s.Repo.Sale.SearchDeliveryNoteViews(req.Ctx, tx, req.Payload)
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
	resp := SearchDeliveryNoteViewResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE VIEW
type GetDeliveryNoteViewRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *SaleService) NewGetDeliveryNoteViewRequest(ctx context.Context, id int) *GetDeliveryNoteViewRequest {
	return &GetDeliveryNoteViewRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetDeliveryNoteViewResponse struct {
	Payload sale.DeliveryNoteView `json:"payload"`
}

func (s *SaleService) NewGetDeliveryNoteViewResponse(payload sale.DeliveryNoteView) *GetDeliveryNoteViewResponse {
	return &GetDeliveryNoteViewResponse{
		Payload: payload,
	}
}

func (s *SaleService) GetDeliveryNoteView(req *GetDeliveryNoteViewRequest) (*GetDeliveryNoteViewResponse, error) {
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

	res, err := s.Repo.Sale.GetDeliveryNoteViewByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetDeliveryNoteViewResponse{
		Payload: *res,
	}
	return &resp, nil
}
