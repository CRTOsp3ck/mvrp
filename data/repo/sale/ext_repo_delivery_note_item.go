package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetDeliveryNoteItemsByDeliveryNoteID(ctx context.Context, exec boil.ContextExecutor, id int) (sale.DeliveryNoteItemSlice, error) {
	return sale.DeliveryNoteItems(qm.Where("delivery_note_id = ?", id)).All(ctx, exec)
}
