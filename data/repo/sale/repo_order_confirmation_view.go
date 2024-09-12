// Code generated by MVRP Codegen Util. DO NOT EDIT.

package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) ListAllOrderConfirmationViews(ctx context.Context, exec boil.ContextExecutor) (sale.OrderConfirmationViewSlice, error) {
	return sale.OrderConfirmationViews().All(ctx, exec)
}

/*
func (r *SaleRepository) SearchOrderConfirmationViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchOrderConfirmationDTO) (sale.OrderConfirmationViewSlice, error) {
	return sale.OrderConfirmationViews(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *SaleRepository) SearchOrderConfirmationViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchOrderConfirmationDTO) (sale.OrderConfirmationViewSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return sale.OrderConfirmationViews(queryMods...).All(ctx, exec)
}

func (r *SaleRepository) GetOrderConfirmationViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.OrderConfirmationView, error) {
	return sale.OrderConfirmationViews(qm.Where(sale.OrderConfirmationViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *SaleRepository) OrderConfirmationViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.OrderConfirmationViews(qm.Where(sale.OrderConfirmationViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *SaleRepository) GetOrderConfirmationViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.OrderConfirmationViews().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentOrderConfirmationView(ctx context.Context, exec boil.ContextExecutor) (*sale.OrderConfirmationView, error) {
	return sale.OrderConfirmationViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}