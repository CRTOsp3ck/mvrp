// Code generated by MVRP Codegen Util. DO NOT EDIT.

package sale

import (
	"context"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) ListAllDeliveryNoteViews(ctx context.Context, exec boil.ContextExecutor) (sale.DeliveryNoteViewSlice, error) {
	return sale.DeliveryNoteViews().All(ctx, exec)
}
func (r *SaleRepository) SearchDeliveryNoteViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchDeliveryNoteDTO) (sale.DeliveryNoteViewSlice, error) {
	return sale.DeliveryNoteViews(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *SaleRepository) GetDeliveryNoteViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.DeliveryNoteView, error) {
	return sale.DeliveryNoteViews(qm.Where(sale.DeliveryNoteViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *SaleRepository) DeliveryNoteViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.DeliveryNoteViews(qm.Where(sale.DeliveryNoteViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *SaleRepository) GetDeliveryNoteViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.DeliveryNoteViews().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentDeliveryNoteView(ctx context.Context, exec boil.ContextExecutor) (*sale.DeliveryNoteView, error) {
	return sale.DeliveryNoteViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}