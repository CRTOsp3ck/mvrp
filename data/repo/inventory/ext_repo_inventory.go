package inventory

import (
	"context"
	"mvrp/data/model/inventory"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetInventoryByItemID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.Inventory, error) {
	return inventory.Inventories(
		qm.Where("item_id=?", id),
	).One(ctx, exec)
}

func (r *InventoryRepository) GetInventoryExistsByItemID(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.Inventories(
		qm.Where("item_id=?", id),
	).Exists(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(ctx context.Context, exec boil.ContextExecutor, id int) (inventory.ReturnMerchandiseAuthorizationItemSlice, error) {
	return inventory.ReturnMerchandiseAuthorizationItems(
		qm.Where("rma_id=?", id),
	).All(ctx, exec)
}

func (r *InventoryRepository) ListInventoriesByItemType(ctx context.Context, exec boil.ContextExecutor, itemType string) (inventory.InventorySlice, error) {
	return inventory.Inventories(
		qm.InnerJoin("item.item ON inventory.inventory.item_id = item.item.id"),
		qm.Where("item.item.type = ?", itemType),
	).All(ctx, exec)
}
