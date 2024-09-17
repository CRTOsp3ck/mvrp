package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetOrderConfirmationByBaseDocumentID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.OrderConfirmation, error) {
	return sale.OrderConfirmations(qm.Where("base_document_id = ?", id)).One(ctx, exec)
}
