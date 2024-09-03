// Code generated by MVRP Codegen Util. DO NOT EDIT.

package purchase

import (
	"context"
	"database/sql"
	"mvrp/data/model/purchase"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *PurchaseRepository) ListAllGoodsReceiptNoteItems(ctx context.Context, exec boil.ContextExecutor) (purchase.GoodsReceiptNoteItemSlice, error) {
	return purchase.GoodsReceiptNoteItems().All(ctx, exec)
}
func (r *PurchaseRepository) SearchGoodsReceiptNoteItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchGoodsReceiptNoteItemDTO) (purchase.GoodsReceiptNoteItemSlice, error) {
	return purchase.GoodsReceiptNoteItems(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *PurchaseRepository) GetGoodsReceiptNoteItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.GoodsReceiptNoteItem, error) {
	return purchase.FindGoodsReceiptNoteItem(ctx, exec, id)
}

func (r *PurchaseRepository) CreateGoodsReceiptNoteItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.GoodsReceiptNoteItem) error {
	id, err := r.GetNextEntryGoodsReceiptNoteItemID(ctx, exec)
	if err != nil {
		return err
	}
	m.ID = id
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *PurchaseRepository) UpdateGoodsReceiptNoteItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.GoodsReceiptNoteItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *PurchaseRepository) UpsertGoodsReceiptNoteItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.GoodsReceiptNoteItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *PurchaseRepository) DeleteGoodsReceiptNoteItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.GoodsReceiptNoteItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *PurchaseRepository) GoodsReceiptNoteItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.GoodsReceiptNoteItemExists(ctx, exec, id)
}

func (r *PurchaseRepository) GetGoodsReceiptNoteItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.GoodsReceiptNoteItems().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentGoodsReceiptNoteItem(ctx context.Context, exec boil.ContextExecutor) (*purchase.GoodsReceiptNoteItem, error) {
	return purchase.GoodsReceiptNoteItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *PurchaseRepository) GetNextEntryGoodsReceiptNoteItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	currID, err := r.GetMostRecentGoodsReceiptNoteItem(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		}
		return 0, err
	}
	return currID.ID + 1, nil
}