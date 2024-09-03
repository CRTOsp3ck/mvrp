package sale

import (
	"context"
	"mvrp/data/model/sale"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) GetSalesOrderItemsBySalesOrderID(ctx context.Context, exec boil.ContextExecutor, id int) (sale.SalesOrderItemSlice, error) {
	return sale.SalesOrderItems(qm.Where("sales_order_id = ?", id)).All(ctx, exec)
}
