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

func (r *InventoryRepository) ListAllInventories(ctx context.Context, exec boil.ContextExecutor) (inventory.InventorySlice, error) {
	return inventory.Inventories().All(ctx, exec)
}
func (r *InventoryRepository) SearchInventories(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryDTO) (inventory.InventorySlice, error) {
	return inventory.Inventories(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetInventoryByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.Inventory, error) {
	return inventory.FindInventory(ctx, exec, id)
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, exec boil.ContextExecutor, m *inventory.Inventory) error {
	id, err := r.GetNextEntryInventoryID(ctx, exec)
	if err != nil {
		return err
	}
	m.ID = id
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InventoryRepository) UpdateInventory(ctx context.Context, exec boil.ContextExecutor, m *inventory.Inventory) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InventoryRepository) UpsertInventory(ctx context.Context, exec boil.ContextExecutor, m *inventory.Inventory) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InventoryRepository) DeleteInventory(ctx context.Context, exec boil.ContextExecutor, m *inventory.Inventory) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InventoryRepository) InventoryExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.InventoryExists(ctx, exec, id)
}

func (r *InventoryRepository) GetInventoryRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.Inventories().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentInventory(ctx context.Context, exec boil.ContextExecutor) (*inventory.Inventory, error) {
	return inventory.Inventories(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) GetNextEntryInventoryID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	currID, err := r.GetMostRecentInventory(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		}
		return 0, err
	}
	return currID.ID + 1, nil
}

func (r *InventoryRepository) GetInventoryTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.Inventories().Count(ctx, exec)
	return int(count), err
}