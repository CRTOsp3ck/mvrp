// Code generated by MVRP Codegen Util. DO NOT EDIT.

package purchase

import (
	"context"
	"mvrp/data/model/purchase"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *PurchaseRepository) ListAllPurchaseOrderViews(ctx context.Context, exec boil.ContextExecutor) (purchase.PurchaseOrderViewSlice, error) {
	return purchase.PurchaseOrderViews().All(ctx, exec)
}

func (r *PurchaseRepository) GetPurchaseOrderViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.PurchaseOrderView, error) {
	return purchase.PurchaseOrderViews(qm.Where(purchase.PurchaseOrderViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *PurchaseRepository) PurchaseOrderViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.PurchaseOrderViews(qm.Where(purchase.PurchaseOrderViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *PurchaseRepository) GetPurchaseOrderViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.PurchaseOrderViews().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentPurchaseOrderView(ctx context.Context, exec boil.ContextExecutor) (*purchase.PurchaseOrderView, error) {
	return purchase.PurchaseOrderViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}