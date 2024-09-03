// Code generated by MVRP Codegen Util. DO NOT EDIT.

package purchase

import (
	"context"
	"mvrp/data/model/purchase"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *PurchaseRepository) ListAllGoodsReceiptNoteItemViews(ctx context.Context, exec boil.ContextExecutor) (purchase.GoodsReceiptNoteItemViewSlice, error) {
	return purchase.GoodsReceiptNoteItemViews().All(ctx, exec)
}

func (r *PurchaseRepository) GetGoodsReceiptNoteItemViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.GoodsReceiptNoteItemView, error) {
	return purchase.GoodsReceiptNoteItemViews(qm.Where(purchase.GoodsReceiptNoteItemViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *PurchaseRepository) GoodsReceiptNoteItemViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.GoodsReceiptNoteItemViews(qm.Where(purchase.GoodsReceiptNoteItemViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *PurchaseRepository) GetGoodsReceiptNoteItemViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.GoodsReceiptNoteItemViews().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentGoodsReceiptNoteItemView(ctx context.Context, exec boil.ContextExecutor) (*purchase.GoodsReceiptNoteItemView, error) {
	return purchase.GoodsReceiptNoteItemViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}