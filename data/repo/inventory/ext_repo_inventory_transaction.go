package inventory

import (
	"context"
	"mvrp/data/model/inventory"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetInventoryTransactionTotalCountByInventoryID(ctx context.Context, exec boil.ContextExecutor, inventoryID string) (int, error) {
	count, err := inventory.InventoryTransactions(qm.Where("inventory_id = ?", inventoryID)).Count(ctx, exec)
	return int(count), err
}
