package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetGoodsReturnNoteItemsByGoodsReturnNoteID(ctx context.Context, exec boil.ContextExecutor, id int) (sale.GoodsReturnNoteItemSlice, error) {
	return sale.GoodsReturnNoteItems(qm.Where("goods_return_note_id = ?", id)).All(ctx, exec)
}
