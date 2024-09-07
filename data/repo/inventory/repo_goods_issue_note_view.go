// Code generated by MVRP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) ListAllGoodsIssueNoteViews(ctx context.Context, exec boil.ContextExecutor) (inventory.GoodsIssueNoteViewSlice, error) {
	return inventory.GoodsIssueNoteViews().All(ctx, exec)
}
func (r *InventoryRepository) SearchGoodsIssueNoteViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchGoodsIssueNoteDTO) (inventory.GoodsIssueNoteViewSlice, error) {
	return inventory.GoodsIssueNoteViews(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetGoodsIssueNoteViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.GoodsIssueNoteView, error) {
	return inventory.GoodsIssueNoteViews(qm.Where(inventory.GoodsIssueNoteViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *InventoryRepository) GoodsIssueNoteViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.GoodsIssueNoteViews(qm.Where(inventory.GoodsIssueNoteViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *InventoryRepository) GetGoodsIssueNoteViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.GoodsIssueNoteViews().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentGoodsIssueNoteView(ctx context.Context, exec boil.ContextExecutor) (*inventory.GoodsIssueNoteView, error) {
	return inventory.GoodsIssueNoteViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}