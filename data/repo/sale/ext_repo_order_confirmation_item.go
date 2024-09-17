package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetOrderConfirmationItemsByOrderConfirmationID(ctx context.Context, exec boil.ContextExecutor, id int) (sale.OrderConfirmationItemSlice, error) {
	return sale.OrderConfirmationItems(qm.Where("order_confirmation_id = ?", id)).All(ctx, exec)
}

func (r *SaleRepository) GetOrderConfirmationItemByBaseDocumentItemID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.OrderConfirmationItem, error) {
	return sale.OrderConfirmationItems(qm.Where("base_document_item_id = ?", id)).One(ctx, exec)
}
