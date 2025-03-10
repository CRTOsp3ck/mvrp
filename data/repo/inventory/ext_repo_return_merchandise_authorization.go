package inventory

import (
	"context"
	"mvrp/data/model/inventory"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(ctx context.Context, exec boil.ContextExecutor, id int) (inventory.ReturnMerchandiseAuthorizationItemSlice, error) {
	return inventory.ReturnMerchandiseAuthorizationItems(
		qm.Where("rma_id=?", id),
	).All(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationByBaseDocumentID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.ReturnMerchandiseAuthorization, error) {
	return inventory.ReturnMerchandiseAuthorizations(
		qm.Where("base_document_id=?", id),
	).One(ctx, exec)
}
