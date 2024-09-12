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

func (r *SaleRepository) ListAllSalesOrderItems(ctx context.Context, exec boil.ContextExecutor) (sale.SalesOrderItemSlice, error) {
	return sale.SalesOrderItems().All(ctx, exec)
}

/*
func (r *SaleRepository) SearchSalesOrderItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchSalesOrderItemDTO) (sale.SalesOrderItemSlice, error) {
	return sale.SalesOrderItems(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *SaleRepository) SearchSalesOrderItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchSalesOrderItemDTO) (sale.SalesOrderItemSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return sale.SalesOrderItems(queryMods...).All(ctx, exec)
}

func (r *SaleRepository) GetSalesOrderItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*sale.SalesOrderItem, error) {
	return sale.FindSalesOrderItem(ctx, exec, id)
}

func (r *SaleRepository) CreateSalesOrderItem(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesOrderItem) error {
	/*
		id, err := r.GetNextEntrySalesOrderItemID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *SaleRepository) UpdateSalesOrderItem(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesOrderItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *SaleRepository) UpsertSalesOrderItem(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesOrderItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *SaleRepository) DeleteSalesOrderItem(ctx context.Context, exec boil.ContextExecutor, m *sale.SalesOrderItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *SaleRepository) SalesOrderItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return sale.SalesOrderItemExists(ctx, exec, id)
}

func (r *SaleRepository) GetSalesOrderItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.SalesOrderItems().Count(ctx, exec)
	return int(count), err
}

func (r *SaleRepository) GetMostRecentSalesOrderItem(ctx context.Context, exec boil.ContextExecutor) (*sale.SalesOrderItem, error) {
	return sale.SalesOrderItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *SaleRepository) GetNextEntrySalesOrderItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := sale.SalesOrderItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentSalesOrderItem(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *SaleRepository) GetSalesOrderItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := sale.SalesOrderItems().Count(ctx, exec)
	return int(count), err
}