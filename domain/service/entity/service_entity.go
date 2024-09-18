package entity

import (
	"context"
	"mvrp/data/model/entity"
	"mvrp/data/repo"
	"mvrp/domain/dto"
)

// LIST ENTITY
type ListEntityRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
}

func (s *EntityService) NewListEntityRequest(ctx context.Context) *ListEntityRequest {
	return &ListEntityRequest{
		Ctx: ctx,
	}
}

type ListEntityResponse struct {
	Payload entity.EntitySlice `json:"payload"`
}

func (s *EntityService) NewListEntityResponse(payload entity.EntitySlice) *ListEntityResponse {
	return &ListEntityResponse{
		Payload: payload,
	}
}
func (s *EntityService) ListEntity(req *ListEntityRequest) (*ListEntityResponse, error) {
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

	res, err := s.Repo.Entity.ListAllEntities(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := ListEntityResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH ENTITY
type SearchEntityRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.SearchEntityDTO
}

func (s *EntityService) NewSearchEntityRequest(ctx context.Context, payload dto.SearchEntityDTO) *SearchEntityRequest {
	return &SearchEntityRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchEntityResponse struct {
	Payload    entity.EntitySlice `json:"payload"`
	Pagination dto.PaginationDTO  `json:"pagination"`
}

func (s *EntityService) NewSearchEntityResponse(payload entity.EntitySlice) *SearchEntityResponse {
	return &SearchEntityResponse{
		Payload: payload,
	}
}

func (s *EntityService) SearchEntity(req *SearchEntityRequest) (*SearchEntityResponse, error) {
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

	res, count, err := s.Repo.Entity.SearchEntities(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	// Pagination
	pd := dto.PaginationDTO{
		TotalItems:   count,
		ItemsPerPage: req.Payload.ItemsPerPage,
		Page:         req.Payload.Page,
		SortBy:       req.Payload.SortBy,
		OrderBy:      req.Payload.OrderBy,
	}
	resp := SearchEntityResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET ENTITY
type GetEntityRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *EntityService) NewGetEntityRequest(ctx context.Context, id int) *GetEntityRequest {
	return &GetEntityRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetEntityResponse struct {
	Payload entity.Entity `json:"payload"`
}

func (s *EntityService) NewGetEntityResponse(payload entity.Entity) *GetEntityResponse {
	return &GetEntityResponse{
		Payload: payload,
	}
}

func (s *EntityService) GetEntity(req *GetEntityRequest) (*GetEntityResponse, error) {
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

	res, err := s.Repo.Entity.GetEntityByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := GetEntityResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE ENTITY
type CreateEntityRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.CreateEntityDTO
}

func (s *EntityService) NewCreateEntityRequest(ctx context.Context, payload dto.CreateEntityDTO) *CreateEntityRequest {
	return &CreateEntityRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateEntityResponse struct {
	Payload entity.Entity `json:"payload"`
}

func (s *EntityService) NewCreateEntityResponse(payload entity.Entity) *CreateEntityResponse {
	return &CreateEntityResponse{
		Payload: payload,
	}
}

func (s *EntityService) CreateEntity(req *CreateEntityRequest) (*CreateEntityResponse, error) {
	/*
		1. Create Entity
	*/

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

	// create entity
	nextID, err := s.Repo.Entity.GetNextEntryEntityID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.Entity.ID = nextID
	err = s.Repo.Entity.CreateEntity(req.Ctx, tx, &req.Payload.Entity)
	if err != nil {
		return nil, err
	}

	// get entity
	res, err := s.Repo.Entity.GetEntityByID(req.Ctx, tx, req.Payload.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := CreateEntityResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE ENTITY
type UpdateEntityRequest struct {
	Ctx     context.Context
	RepoTx  *repo.RepoTx
	Payload dto.UpdateEntityDTO
}

func (s *EntityService) NewUpdateEntityRequest(ctx context.Context, payload dto.UpdateEntityDTO) *UpdateEntityRequest {
	return &UpdateEntityRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateEntityResponse struct {
	Payload entity.Entity `json:"payload"`
}

func (s *EntityService) NewUpdateEntityResponse(payload entity.Entity) *UpdateEntityResponse {
	return &UpdateEntityResponse{
		Payload: payload,
	}
}

func (s *EntityService) UpdateEntity(req *UpdateEntityRequest) (*UpdateEntityResponse, error) {
	/*
		1. Update Entity
	*/

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

	// update entity
	err = s.Repo.Entity.UpdateEntity(req.Ctx, tx, &req.Payload.Entity)
	if err != nil {
		return nil, err
	}

	// get updated entity
	res, err := s.Repo.Entity.GetEntityByID(req.Ctx, tx, req.Payload.ID)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := UpdateEntityResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE ENTITY
type DeleteEntityRequest struct {
	Ctx    context.Context
	RepoTx *repo.RepoTx
	ID     int
}

func (s *EntityService) NewDeleteEntityRequest(ctx context.Context, id int) *DeleteEntityRequest {
	return &DeleteEntityRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteEntityResponse struct {
	Payload bool `json:"payload"`
}

func (s *EntityService) NewDeleteEntityResponse(payload bool) *DeleteEntityResponse {
	return &DeleteEntityResponse{
		Payload: payload,
	}
}

func (s *EntityService) DeleteEntity(req *DeleteEntityRequest) (*DeleteEntityResponse, error) {
	/*
		1. Delete Entity
	*/

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

	// get entity
	entity, err := s.Repo.Entity.GetEntityByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete entity
	err = s.Repo.Entity.DeleteEntity(req.Ctx, tx, entity)
	if err != nil {
		return nil, err
	}

	if req.RepoTx == nil {
		err = tx.Commit()
		if err != nil {
			return nil, err
		}
	}

	resp := DeleteEntityResponse{
		Payload: true,
	}
	return &resp, nil
}
