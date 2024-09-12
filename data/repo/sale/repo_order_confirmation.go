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

func (r *SaleRepository) ListAllOrderConfirmations(ctx context.Context, exec boil.ContextExecutor) (sale.OrderConfirmationSlice, error) {
	return sale.OrderConfirmations().All(ctx, exec)
}

/*
func (r *SaleRepository) SearchOrderConfirmations(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchOrderConfirmationDTO) (sale.OrderConfirmationSlice, error) {
	return sale.OrderConfirmations(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *SaleRepository) SearchOrderConfirmations(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchOrderConfirmationDTO) (sale.OrderConfirmationSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return sale.OrderConfirmations(queryMods...).All(ctx, exec)
}

func (r *SaleRepository) GetOrderConfirmationByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.OrderConfirmation, error) {
	return sale.FindOrderConfirmation(ctx, exec, id)
}

func (r *SaleRepository) CreateOrderConfirmation(ctx context.Context, exec boil.ContextExecutor, m *sale.OrderConfirmation) error {
	/*
		id, err := r.GetNextEntryOrderConfirmationID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *SaleRepository) UpdateOrderConfirmation(ctx context.Context, exec boil.ContextExecutor, m *sale.OrderConfirmation) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *SaleRepository) UpsertOrderConfirmation(ctx context.Context, exec boil.ContextExecutor, m *sale.OrderConfirmation) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *SaleRepository) DeleteOrderConfirmation(ctx context.Context, exec boil.ContextExecutor, m *sale.OrderConfirmation) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *SaleRepository) OrderConfirmationExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.OrderConfirmationExists(ctx, exec, id)
}

func (r *SaleRepository) GetOrderConfirmationRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.OrderConfirmations().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentOrderConfirmation(ctx context.Context, exec boil.ContextExecutor) (*sale.OrderConfirmation, error) {
	return sale.OrderConfirmations(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *SaleRepository) GetNextEntryOrderConfirmationID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := sale.OrderConfirmations(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentOrderConfirmation(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *SaleRepository) GetOrderConfirmationTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.OrderConfirmations().Count(ctx, exec)
	return int(count), err
}