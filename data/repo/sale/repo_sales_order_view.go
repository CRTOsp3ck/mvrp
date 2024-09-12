// Code generated by MVRP Codegen Util. DO NOT EDIT.

package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) ListAllSalesOrderViews(ctx context.Context, exec boil.ContextExecutor) (sale.SalesOrderViewSlice, error) {
	return sale.SalesOrderViews().All(ctx, exec)
}

/*
func (r *SaleRepository) SearchSalesOrderViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchSalesOrderDTO) (sale.SalesOrderViewSlice, error) {
	return sale.SalesOrderViews(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *SaleRepository) SearchSalesOrderViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchSalesOrderDTO) (sale.SalesOrderViewSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return sale.SalesOrderViews(queryMods...).All(ctx, exec)
}

func (r *SaleRepository) GetSalesOrderViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.SalesOrderView, error) {
	return sale.SalesOrderViews(qm.Where(sale.SalesOrderViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *SaleRepository) SalesOrderViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.SalesOrderViews(qm.Where(sale.SalesOrderViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *SaleRepository) GetSalesOrderViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.SalesOrderViews().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentSalesOrderView(ctx context.Context, exec boil.ContextExecutor) (*sale.SalesOrderView, error) {
	return sale.SalesOrderViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}