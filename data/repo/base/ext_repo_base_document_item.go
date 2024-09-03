package base

import (
	"context"
	"mvrp/data/model/base"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *BaseRepository) GetBaseDocumentItemsByBaseDocumentID(ctx context.Context, exec boil.ContextExecutor, id int) (base.BaseDocumentItemSlice, error) {
	return base.BaseDocumentItems(qm.Where("base_document_id = ?", id)).All(ctx, exec)
}
