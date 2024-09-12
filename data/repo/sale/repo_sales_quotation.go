// Code generated by MVRP Codegen Util. DO NOT EDIT.

package sale

import (
	"context"
	"database/sql"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *SaleRepository) ListAllSalesQuotations(ctx context.Context, exec boil.ContextExecutor) (sale.SalesQuotationSlice, error) {
	return sale.SalesQuotations().All(ctx, exec)
}
func (r *SaleRepository) SearchSalesQuotations(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchSalesQuotationDTO) (sale.SalesQuotationSlice, error) {
	return sale.SalesQuotations(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *SaleRepository) GetSalesQuotationByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.SalesQuotation, error) {
	return sale.FindSalesQuotation(ctx, exec, id)
}

func (r *SaleRepository) CreateSalesQuotation(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesQuotation) error {
	/*
		id, err := r.GetNextEntrySalesQuotationID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *SaleRepository) UpdateSalesQuotation(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesQuotation) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *SaleRepository) UpsertSalesQuotation(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesQuotation) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *SaleRepository) DeleteSalesQuotation(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesQuotation) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *SaleRepository) SalesQuotationExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.SalesQuotationExists(ctx, exec, id)
}

func (r *SaleRepository) GetSalesQuotationRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.SalesQuotations().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentSalesQuotation(ctx context.Context, exec boil.ContextExecutor) (*sale.SalesQuotation, error) {
	return sale.SalesQuotations(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *SaleRepository) GetNextEntrySalesQuotationID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := sale.SalesQuotations(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentSalesQuotation(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *SaleRepository) GetSalesQuotationTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.SalesQuotations().Count(ctx, exec)
	return int(count), err
}