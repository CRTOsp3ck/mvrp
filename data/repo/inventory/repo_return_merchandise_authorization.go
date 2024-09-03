// Code generated by ZERP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"database/sql"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) ListAllReturnMerchandiseAuthorizations(ctx context.Context, exec boil.ContextExecutor) (inventory.ReturnMerchandiseAuthorizationSlice, error) {
	return inventory.ReturnMerchandiseAuthorizations().All(ctx, exec)
}
func (r *InventoryRepository) SearchReturnMerchandiseAuthorizations(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchReturnMerchandiseAuthorizationDTO) (inventory.ReturnMerchandiseAuthorizationSlice, error) {
	return inventory.ReturnMerchandiseAuthorizations(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.ReturnMerchandiseAuthorization, error) {
	return inventory.FindReturnMerchandiseAuthorization(ctx, exec, id)
}

func (r *InventoryRepository) CreateReturnMerchandiseAuthorization(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorization) error {
	id, err := r.GetNextEntryReturnMerchandiseAuthorizationID(ctx, exec)
	if err != nil {
		return err
	}
	m.ID = id
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InventoryRepository) UpdateReturnMerchandiseAuthorization(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorization) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InventoryRepository) UpsertReturnMerchandiseAuthorization(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorization) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InventoryRepository) DeleteReturnMerchandiseAuthorization(ctx context.Context, exec boil.ContextExecutor, m *inventory.ReturnMerchandiseAuthorization) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InventoryRepository) ReturnMerchandiseAuthorizationExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.ReturnMerchandiseAuthorizationExists(ctx, exec, id)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.ReturnMerchandiseAuthorizations().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentReturnMerchandiseAuthorization(ctx context.Context, exec boil.ContextExecutor) (*inventory.ReturnMerchandiseAuthorization, error) {
	return inventory.ReturnMerchandiseAuthorizations(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) GetNextEntryReturnMerchandiseAuthorizationID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	currID, err := r.GetMostRecentReturnMerchandiseAuthorization(ctx, exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1, nil
		}
		return 0, err
	}
	return currID.ID + 1, nil
}