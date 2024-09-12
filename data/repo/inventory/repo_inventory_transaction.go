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

func (r *InventoryRepository) ListAllInventoryTransactions(ctx context.Context, exec boil.ContextExecutor) (inventory.InventoryTransactionSlice, error) {
	return inventory.InventoryTransactions().All(ctx, exec)
}
func (r *InventoryRepository) SearchInventoryTransactions(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryTransactionDTO) (inventory.InventoryTransactionSlice, error) {
	return inventory.InventoryTransactions(
		qm.Where("inventory_id = ?", dto.InventoryId),
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetInventoryTransactionByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.InventoryTransaction, error) {
	return inventory.FindInventoryTransaction(ctx, exec, id)
}

func (r *InventoryRepository) CreateInventoryTransaction(ctx context.Context, exec boil.ContextExecutor, m *inventory.InventoryTransaction) error {
	/*
		id, err := r.GetNextEntryInventoryTransactionID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InventoryRepository) UpdateInventoryTransaction(ctx context.Context, exec boil.ContextExecutor, m *inventory.InventoryTransaction) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InventoryRepository) UpsertInventoryTransaction(ctx context.Context, exec boil.ContextExecutor, m *inventory.InventoryTransaction) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InventoryRepository) DeleteInventoryTransaction(ctx context.Context, exec boil.ContextExecutor, m *inventory.InventoryTransaction) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InventoryRepository) InventoryTransactionExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.InventoryTransactionExists(ctx, exec, id)
}

func (r *InventoryRepository) GetInventoryTransactionRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.InventoryTransactions().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentInventoryTransaction(ctx context.Context, exec boil.ContextExecutor) (*inventory.InventoryTransaction, error) {
	return inventory.InventoryTransactions(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) GetNextEntryInventoryTransactionID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := inventory.InventoryTransactions(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentInventoryTransaction(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *InventoryRepository) GetInventoryTransactionTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.InventoryTransactions().Count(ctx, exec)
	return int(count), err
}