package inventory

import (
	"context"
	"mvrp/data/model/inventory"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetGoodsIssueNoteItemsByGoodsIssueNoteID(ctx context.Context, exec boil.ContextExecutor, gin_id int) (inventory.GoodsIssueNoteItemSlice, error) {
	return inventory.GoodsIssueNoteItems(
		qm.Where("gin_id=?", gin_id),
	).All(ctx, exec)
}
