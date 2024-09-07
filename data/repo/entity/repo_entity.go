// Code generated by MVRP Codegen Util. DO NOT EDIT.

package entity

import (
	"context"
	"database/sql"
	"mvrp/data/model/entity"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *EntityRepository) ListAllEntities(ctx context.Context, exec boil.ContextExecutor) (entity.EntitySlice, error) {
	return entity.Entities().All(ctx, exec)
}
func (r *EntityRepository) SearchEntities(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchEntityDTO) (entity.EntitySlice, error) {
	return entity.Entities(
		qm.Where("type = ?", dto.Type),
        qm.And(
			"name ILIKE ? or address ILIKE ? or email ILIKE ?",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
		),
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *EntityRepository) GetEntityByID(ctx context.Context, exec boil.ContextExecutor, id int) (*entity.Entity, error) {
	return entity.FindEntity(ctx, exec, id)
}

func (r *EntityRepository) CreateEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	id, err := r.GetNextEntryEntityID(ctx, exec)
	if err != nil {
		return err
	}
	m.ID = id
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *EntityRepository) UpdateEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *EntityRepository) UpsertEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *EntityRepository) DeleteEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *EntityRepository) EntityExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return entity.EntityExists(ctx, exec, id)
}

func (r *EntityRepository) GetEntityRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := entity.Entities().Count(ctx, exec)
	return int(count), err
}

func (r *EntityRepository) GetMostRecentEntity(ctx context.Context, exec boil.ContextExecutor) (*entity.Entity, error) {
	return entity.Entities(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *EntityRepository) GetNextEntryEntityID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	currID, err := r.GetMostRecentEntity(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		}
		return 0, err
	}
	return currID.ID + 1, nil
}

func (r *EntityRepository) GetEntityTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := entity.Entities().Count(ctx, exec)
	return int(count), err
}