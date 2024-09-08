package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetInventoryTransactionTotalCountByInventoryID(ctx context.Context, exec boil.ContextExecutor, inventoryID string) (int, error) {
	count, err := inventory.InventoryTransactions(qm.Where("inventory_id = ?", inventoryID)).Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) SearchAllInventoryTransactions(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryTransactionDTO) (inventory.InventoryTransactionSlice, error) {
	return inventory.InventoryTransactions(
		// qm.Where("inventory_id = ?", dto.InventoryId),
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
