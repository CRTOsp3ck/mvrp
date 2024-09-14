// Code generated by MVRP Codegen Util. DO NOT EDIT.

package item

import (
	"context"
	"database/sql"
	"mvrp/data/model/item"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *ItemRepository) ListAllItems(ctx context.Context, exec boil.ContextExecutor) (item.ItemSlice, error) {
	return item.Items().All(ctx, exec)
}

func (r *ItemRepository) GetItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*item.Item, error) {
	return item.FindItem(ctx, exec, id)
}

func (r *ItemRepository) CreateItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *ItemRepository) UpdateItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *ItemRepository) UpsertItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *ItemRepository) DeleteItem(ctx context.Context, exec boil.ContextExecutor, m *item.Item) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *ItemRepository) ItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return item.ItemExists(ctx, exec, id)
}

func (r *ItemRepository) GetItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := item.Items().Count(ctx, exec)
	return int(count), err
}

func (r *ItemRepository) GetMostRecentItem(ctx context.Context, exec boil.ContextExecutor) (*item.Item, error) {
	return item.Items(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *ItemRepository) GetNextEntryItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := item.Items(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *ItemRepository) GetItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := item.Items().Count(ctx, exec)
	return int(count), err
}

func (r *ItemRepository) SearchItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchItemDTO) (item.ItemSlice, int, error) {
	return r.BuildSearchQueryForItems(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		if dto.Type != "" {
			queryMods = append(queryMods, qm.Where("type = ?", dto.Type))
		}
		queryMods = append(queryMods, qm.And(
			"code ILIKE ? or name ILIKE ? or brand ILIKE ? or description ILIKE ? or category ILIKE ? or origin ILIKE ?",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
			"%" + dto.Keyword + "%",
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
		return item.Items(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *ItemRepository) BuildSearchQueryForItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchItemDTO) (item.ItemSlice, int, error) {
	var queryMods []qm.QueryMod
	if dto.Type != "" {
		queryMods = append(queryMods, qm.Where("type = ?", dto.Type))
	}

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForItems(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForItems(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForItems(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForItems(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForItems(request)
	if limitSQL > 0 {
		queryMods = append(queryMods, qm.Limit(limitSQL))
	}
	if offsetSQL > 0 {
		queryMods = append(queryMods, qm.Offset(offsetSQL))
	}

	countQueryMods := []qm.QueryMod{
		qm.Where("entity_type = ?", "customer"),
	}

	if whereSQL != "" {
		countQueryMods = append(countQueryMods, qm.Where(whereSQL))
	}

	count, err := item.Items(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := item.Items(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *ItemRepository) createSelectSQLForItems(request query.IServerSideGetRowsRequest) string {
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
func (r *ItemRepository) createWhereSQLForItems(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForItems(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}

func (r *ItemRepository) createFilterSQLForItems(colID string, filterItem query.FilterItem) string {
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForItems(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForItems(colID, filterItem)
	default:
		return "true"
	}
}

func (r *ItemRepository) createNumberFilterSQLForItems(colID string, filterItem query.FilterItem) string {
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

func (r *ItemRepository) createTextFilterSQLForItems(colID string, filterModel query.FilterItem) string {
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
func (r *ItemRepository) createGroupBySQLForItems(request query.IServerSideGetRowsRequest) string {
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
func (r *ItemRepository) createOrderBySQLForItems(request query.IServerSideGetRowsRequest) string {
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
func (r *ItemRepository) createLimitAndOffsetSQLForItems(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *ItemRepository) createFilterSQLForItems(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForItems(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForItems(colID, filterItem)
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