// Code generated by MVRP Codegen Util. DO NOT EDIT.

package base

import (
	"context"
	"database/sql"
	"mvrp/data/model/base"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *BaseRepository) ListAllBaseDocuments(ctx context.Context, exec boil.ContextExecutor) (base.BaseDocumentSlice, error) {
	return base.BaseDocuments().All(ctx, exec)
}

func (r *BaseRepository) GetBaseDocumentByID(ctx context.Context, exec boil.ContextExecutor, id int) (*base.BaseDocument, error) {
	return base.FindBaseDocument(ctx, exec, id)
}

func (r *BaseRepository) CreateBaseDocument(ctx context.Context, exec boil.ContextExecutor, m *base.BaseDocument) error {
	id, err := r.GetNextEntryBaseDocumentID(ctx, exec)
	if err != nil {
		return err
	}
	m.ID = id
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *BaseRepository) UpdateBaseDocument(ctx context.Context, exec boil.ContextExecutor, m *base.BaseDocument) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *BaseRepository) UpsertBaseDocument(ctx context.Context, exec boil.ContextExecutor, m *base.BaseDocument) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *BaseRepository) DeleteBaseDocument(ctx context.Context, exec boil.ContextExecutor, m *base.BaseDocument) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *BaseRepository) BaseDocumentExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return base.BaseDocumentExists(ctx, exec, id)
}

func (r *BaseRepository) GetBaseDocumentRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := base.BaseDocuments().Count(ctx, exec)
	return int(count), err
}

func (r *BaseRepository) GetMostRecentBaseDocument(ctx context.Context, exec boil.ContextExecutor) (*base.BaseDocument, error) {
	return base.BaseDocuments(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *BaseRepository) GetNextEntryBaseDocumentID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	currID, err := r.GetMostRecentBaseDocument(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		}
		return 0, err
	}
	return currID.ID + 1, nil
}