package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetSalesQuotationItemsBySalesQuotationID(ctx context.Context, exec boil.ContextExecutor, id int) (sale.SalesQuotationItemSlice, error) {
	return sale.SalesQuotationItems(qm.Where("sales_quotation_id = ?", id)).All(ctx, exec)
}
