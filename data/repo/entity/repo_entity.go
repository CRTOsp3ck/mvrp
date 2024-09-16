// Code generated by MVRP Codegen Util. DO NOT EDIT.

package entity

import (
	"context"
	"database/sql"
	"mvrp/data/model/entity"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *EntityRepository) ListAllEntities(ctx context.Context, exec boil.ContextExecutor) (entity.EntitySlice, error) {
	return entity.Entities().All(ctx, exec)
}

func (r *EntityRepository) GetEntityByID(ctx context.Context, exec boil.ContextExecutor, id int) (*entity.Entity, error) {
	return entity.FindEntity(ctx, exec, id)
}

func (r *EntityRepository) CreateEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *EntityRepository) UpdateEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *EntityRepository) UpsertEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *EntityRepository) DeleteEntity(ctx context.Context, exec boil.ContextExecutor, m *entity.Entity) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *EntityRepository) EntityExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return entity.EntityExists(ctx, exec, id)
}

func (r *EntityRepository) GetEntityRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := entity.Entities().Count(ctx, exec)
	return int(count), err
}

func (r *EntityRepository) GetMostRecentEntity(ctx context.Context, exec boil.ContextExecutor) (*entity.Entity, error) {
	return entity.Entities(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *EntityRepository) GetNextEntryEntityID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := entity.Entities(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *EntityRepository) GetEntityTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := entity.Entities().Count(ctx, exec)
	return int(count), err
}

func (r *EntityRepository) SearchEntities(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchEntityDTO) (entity.EntitySlice, int, error) {
	return r.BuildSearchQueryForEntities(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		if dto.Type != "" {
			queryMods = append(queryMods, qm.Where("type = ?", dto.Type))
		}
		queryMods = append(queryMods, qm.And(
			"name ILIKE ? or address ILIKE ? or email ILIKE ?",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
		))
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return entity.Entities(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *EntityRepository) BuildSearchQueryForEntities(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchEntityDTO) (entity.EntitySlice, int, error) {
	var queryMods []qm.QueryMod
	if dto.Type != "" {
		queryMods = append(queryMods, qm.Where("type = ?", dto.Type))
	}

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForEntities(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForEntities(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForEntities(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForEntities(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForEntities(request)
	if limitSQL > 0 {
		queryMods = append(queryMods, qm.Limit(limitSQL))
	}
	if offsetSQL > 0 {
		queryMods = append(queryMods, qm.Offset(offsetSQL))
	}

	// ---------------- Pagination Count Query Mods ---------------- 
	var countQueryMods []qm.QueryMod
	if dto.Type != "" {
		countQueryMods = append(countQueryMods, qm.Where("type = ?", dto.Type))
	}
	// --------------------------------------------------------------

	if whereSQL != "" {
		countQueryMods = append(countQueryMods, qm.Where(whereSQL))
	}

	count, err := entity.Entities(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := entity.Entities(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *EntityRepository) createSelectSQLForEntities(request query.IServerSideGetRowsRequest) string {
	rowGroupCols := request.RowGroupCols
	valueCols := request.ValueCols
	groupKeys := request.GroupKeys

	var colsToSelect []string

	if len(rowGroupCols) > len(groupKeys) {
		// Include group columns in SELECT
		for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
			colsToSelect = append(colsToSelect, rowGroupCols[i].Field)
		}

		// Include aggregated columns in SELECT
		for _, valueCol := range valueCols {
			colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", valueCol.Field, valueCol.Field))
		}
	} else {
		// If not grouping, select all columns directly
		for _, valueCol := range valueCols {
			colsToSelect = append(colsToSelect, valueCol.Field)
		}
	}

	return strings.Join(colsToSelect, ", ")
}

// 2. Create Where SQL (handle group keys and filters)
func (r *EntityRepository) createWhereSQLForEntities(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForEntities(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}

func (r *EntityRepository) createFilterSQLForEntities(colID string, filterItem query.FilterItem) string {
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForEntities(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForEntities(colID, filterItem)
	default:
		return "true"
	}
}

func (r *EntityRepository) createNumberFilterSQLForEntities(colID string, filterItem query.FilterItem) string {
	switch filterItem.Type {
	case "equals":
		return fmt.Sprintf("%s = %s", colID, filterItem.Filter)
	case "greaterThan":
		return fmt.Sprintf("%s > %s", colID, filterItem.Filter)
	case "lessThan":
		return fmt.Sprintf("%s < %s", colID, filterItem.Filter)
	default:
		return "true"
	}
}

func (r *EntityRepository) createTextFilterSQLForEntities(colID string, filterModel query.FilterItem) string {
	switch filterModel.Type {
	case "contains":
		// Ensure that id is filtered with '=' if it's a number
		if colID == "id" {
			return fmt.Sprintf("%s = %s", colID, filterModel.Filter)
		}
		return fmt.Sprintf("%s ILIKE '%%%s%%'", colID, filterModel.Filter)
	case "equals":
		return fmt.Sprintf("%s = '%s'", colID, filterModel.Filter)
	case "notEqual":
		return fmt.Sprintf("%s != '%s'", colID, filterModel.Filter)
	case "startsWith":
		return fmt.Sprintf("%s ILIKE '%s%%'", colID, filterModel.Filter)
	case "endsWith":
		return fmt.Sprintf("%s ILIKE '%%%s'", colID, filterModel.Filter)
	default:
		return "true"
	}
}

// 3. Create Group By SQL
func (r *EntityRepository) createGroupBySQLForEntities(request query.IServerSideGetRowsRequest) string {
	rowGroupCols := request.RowGroupCols
	groupKeys := request.GroupKeys

	var groupByCols []string

	if len(rowGroupCols) > len(groupKeys) {
		for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
			groupByCols = append(groupByCols, rowGroupCols[i].Field)
		}
	}

	return strings.Join(groupByCols, ", ")
}

// 4. Create Order By SQL
func (r *EntityRepository) createOrderBySQLForEntities(request query.IServerSideGetRowsRequest) string {
	sortModel := request.SortModel
	rowGroupCols := request.RowGroupCols
	groupKeys := request.GroupKeys

	var sortParts []string

	// Determine if we are doing grouping
	grouping := len(rowGroupCols) > len(groupKeys)
	groupColIds := make(map[string]struct{})

	// Create a map of grouped columns
	for i := 0; i < len(groupKeys)+1 && i < len(rowGroupCols); i++ {
		groupColIds[rowGroupCols[i].Field] = struct{}{}
	}

	for _, sort := range sortModel {
		if grouping {
			// Only allow sorting on grouped columns
			if _, exists := groupColIds[sort.ColId]; exists {
				sortParts = append(sortParts, fmt.Sprintf("%s %s", sort.ColId, sort.Sort))
			}
		} else {
			// If no grouping, allow sorting on any column
			sortParts = append(sortParts, fmt.Sprintf("%s %s", sort.ColId, sort.Sort))
		}
	}

	if len(sortParts) > 0 {
		return strings.Join(sortParts, ", ")
	}

	return ""
}

// 5. Create Limit and Offset SQL
func (r *EntityRepository) createLimitAndOffsetSQLForEntities(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *EntityRepository) createFilterSQLForEntities(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForEntities(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForEntities(colID, filterItem)
// 	}
// }

// func isNumericColumn(colID string) bool {
// 	// You can enhance this logic by checking actual column types if available
// 	// For example, assuming column names like "id", "price", "age", etc. are numeric.
// 	numericColumns := map[string]bool{
// 		"id":     true,
// 		"price":  true,
// 		"age":    true,
// 		"amount": true,
// 	}
// 	return numericColumns[colID]
// }