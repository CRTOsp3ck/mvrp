package entity

import (
	"context"
	"mvrp/data/model/entity"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *EntityRepository) GetEntityTotalCountByType(ctx context.Context, exec boil.ContextExecutor, entityType string) (int, error) {
	count, err := entity.Entities(qm.Where("entity_type = ?", entityType)).Count(ctx, exec)
	return int(count), err
}
