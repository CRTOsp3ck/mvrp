package item

import (
	"context"
	"mvrp/data/model/item"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *ItemRepository) GetItemTotalCountByType(ctx context.Context, exec boil.ContextExecutor, itemType string) (int, error) {
	count, err := item.Items(qm.Where("type = ?", itemType)).Count(ctx, exec)
	return int(count), err
}
