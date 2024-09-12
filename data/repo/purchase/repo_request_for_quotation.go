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

func (r *PurchaseRepository) ListAllRequestForQuotations(ctx context.Context, exec boil.ContextExecutor) (purchase.RequestForQuotationSlice, error) {
	return purchase.RequestForQuotations().All(ctx, exec)
}
func (r *PurchaseRepository) SearchRequestForQuotations(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchRequestForQuotationDTO) (purchase.RequestForQuotationSlice, error) {
	return purchase.RequestForQuotations(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *PurchaseRepository) GetRequestForQuotationByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.RequestForQuotation, error) {
	return purchase.FindRequestForQuotation(ctx, exec, id)
}

func (r *PurchaseRepository) CreateRequestForQuotation(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotation) error {
	/*
		id, err := r.GetNextEntryRequestForQuotationID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *PurchaseRepository) UpdateRequestForQuotation(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotation) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *PurchaseRepository) UpsertRequestForQuotation(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotation) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *PurchaseRepository) DeleteRequestForQuotation(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotation) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *PurchaseRepository) RequestForQuotationExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.RequestForQuotationExists(ctx, exec, id)
}

func (r *PurchaseRepository) GetRequestForQuotationRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotations().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentRequestForQuotation(ctx context.Context, exec boil.ContextExecutor) (*purchase.RequestForQuotation, error) {
	return purchase.RequestForQuotations(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *PurchaseRepository) GetNextEntryRequestForQuotationID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := purchase.RequestForQuotations(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentRequestForQuotation(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *PurchaseRepository) GetRequestForQuotationTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotations().Count(ctx, exec)
	return int(count), err
}