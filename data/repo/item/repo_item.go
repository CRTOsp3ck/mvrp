// Code generated by MVRP Codegen Util. DO NOT EDIT.

package item

import (
	"context"
	"database/sql"
	"mvrp/data/model/item"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *ItemRepository) ListAllItems(ctx context.Context, exec boil.ContextExecutor) (item.ItemSlice, error) {
	return item.Items().All(ctx, exec)
}
func (r *ItemRepository) SearchItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchItemDTO) (item.ItemSlice, error) {
	return item.Items(
		qm.Where("type = ?", dto.Type),
        qm.And(
			"code ILIKE ? or name ILIKE ? or brand ILIKE ? or description ILIKE ? or category ILIKE ? or origin ILIKE ?",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
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

func (r *ItemRepository) GetItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*item.Item, error) {
	return item.FindItem(ctx, exec, id)
}

func (r *ItemRepository) CreateItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	/*
		id, err := r.GetNextEntryItemID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *ItemRepository) UpdateItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *ItemRepository) UpsertItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *ItemRepository) DeleteItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *ItemRepository) ItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return item.ItemExists(ctx, exec, id)
}

func (r *ItemRepository) GetItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := item.Items().Count(ctx, exec)
	return int(count), err
}

func (r *ItemRepository) GetMostRecentItem(ctx context.Context, exec boil.ContextExecutor) (*item.Item, error) {
	return item.Items(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *ItemRepository) GetNextEntryItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := item.Items(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentItem(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *ItemRepository) GetItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := item.Items().Count(ctx, exec)
	return int(count), err
}