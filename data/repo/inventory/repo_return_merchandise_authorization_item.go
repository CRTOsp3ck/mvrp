// Code generated by MVRP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"database/sql"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) ListAllReturnMerchandiseAuthorizationItems(ctx context.Context, exec boil.ContextExecutor) (inventory.ReturnMerchandiseAuthorizationItemSlice, error) {
	return inventory.ReturnMerchandiseAuthorizationItems().All(ctx, exec)
}
func (r *InventoryRepository) SearchReturnMerchandiseAuthorizationItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchReturnMerchandiseAuthorizationItemDTO) (inventory.ReturnMerchandiseAuthorizationItemSlice, error) {
	return inventory.ReturnMerchandiseAuthorizationItems(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.ReturnMerchandiseAuthorizationItem, error) {
	return inventory.FindReturnMerchandiseAuthorizationItem(ctx, exec, id)
}

func (r *InventoryRepository) CreateReturnMerchandiseAuthorizationItem(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorizationItem) error {
	/*
		id, err := r.GetNextEntryReturnMerchandiseAuthorizationItemID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InventoryRepository) UpdateReturnMerchandiseAuthorizationItem(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorizationItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InventoryRepository) UpsertReturnMerchandiseAuthorizationItem(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorizationItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InventoryRepository) DeleteReturnMerchandiseAuthorizationItem(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorizationItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InventoryRepository) ReturnMerchandiseAuthorizationItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.ReturnMerchandiseAuthorizationItemExists(ctx, exec, id)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.ReturnMerchandiseAuthorizationItems().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentReturnMerchandiseAuthorizationItem(ctx context.Context, exec boil.ContextExecutor) (*inventory.ReturnMerchandiseAuthorizationItem, error) {
	return inventory.ReturnMerchandiseAuthorizationItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) GetNextEntryReturnMerchandiseAuthorizationItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := inventory.ReturnMerchandiseAuthorizationItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentReturnMerchandiseAuthorizationItem(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.ReturnMerchandiseAuthorizationItems().Count(ctx, exec)
	return int(count), err
}