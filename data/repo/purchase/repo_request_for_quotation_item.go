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

func (r *PurchaseRepository) ListAllRequestForQuotationItems(ctx context.Context, exec boil.ContextExecutor) (purchase.RequestForQuotationItemSlice, error) {
	return purchase.RequestForQuotationItems().All(ctx, exec)
}
func (r *PurchaseRepository) SearchRequestForQuotationItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchRequestForQuotationItemDTO) (purchase.RequestForQuotationItemSlice, error) {
	return purchase.RequestForQuotationItems(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *PurchaseRepository) GetRequestForQuotationItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.RequestForQuotationItem, error) {
	return purchase.FindRequestForQuotationItem(ctx, exec, id)
}

func (r *PurchaseRepository) CreateRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	/*
		id, err := r.GetNextEntryRequestForQuotationItemID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *PurchaseRepository) UpdateRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *PurchaseRepository) UpsertRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *PurchaseRepository) DeleteRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *PurchaseRepository) RequestForQuotationItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.RequestForQuotationItemExists(ctx, exec, id)
}

func (r *PurchaseRepository) GetRequestForQuotationItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotationItems().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor) (*purchase.RequestForQuotationItem, error) {
	return purchase.RequestForQuotationItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *PurchaseRepository) GetNextEntryRequestForQuotationItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := purchase.RequestForQuotationItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentRequestForQuotationItem(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *PurchaseRepository) GetRequestForQuotationItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotationItems().Count(ctx, exec)
	return int(count), err
}