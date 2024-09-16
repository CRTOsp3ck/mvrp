package inventory

import (
	"context"
	"fmt"
	"mvrp/data/model/inventory"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) GetInventoryTransactionTotalCountByInventoryID(ctx context.Context, exec boil.ContextExecutor, inventoryID string) (int, error) {
	count, err := inventory.InventoryTransactions(qm.Where("inventory_id = ?", inventoryID)).Count(ctx, exec)
	return int(count), err
}

// func (r *InventoryRepository) SearchAllAllInventoryTransactions(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryTransactionDTO) (inventory.InventoryTransactionSlice, error) {
// 	return inventory.AllInventoryTransactions(
// 		// qm.Where("inventory_id = ?", dto.InventoryId),
// 		qm.Limit(dto.ItemsPerPage),
// 		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
// 		// qm.GroupBy("id"),
// 		qm.OrderBy(dto.OrderBy+" "+"ASC"),
// 	).All(ctx, exec)
// }

func (r *InventoryRepository) SearchAllInventoryTransactions(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryTransactionDTO) (inventory.InventoryTransactionSlice, int, error) {
	return r.BuildSearchQueryForAllInventoryTransactions(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		if dto.InventoryId != "" {
			queryMods = append(queryMods, qm.Where("inventory_id = ?", dto.InventoryId))
		}
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return inventory.AllInventoryTransactions(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InventoryRepository) BuildSearchQueryForAllInventoryTransactions(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInventoryTransactionDTO) (inventory.InventoryTransactionSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForAllInventoryTransactions(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForAllInventoryTransactions(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForAllInventoryTransactions(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForAllInventoryTransactions(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForAllInventoryTransactions(request)
	if limitSQL > 0 {
		queryMods = append(queryMods, qm.Limit(limitSQL))
	}
	if offsetSQL > 0 {
		queryMods = append(queryMods, qm.Offset(offsetSQL))
	}

	// countQueryMods := []qm.QueryMod{
	// 	// qm.Where("inventory_id = ?", dto.InventoryId),
	// }

	// if whereSQL != "" {
	// 	countQueryMods = append(countQueryMods, qm.Where(whereSQL))
	// }

	count, err := inventory.InventoryTransactions().Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := inventory.InventoryTransactions(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *InventoryRepository) createSelectSQLForAllInventoryTransactions(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createWhereSQLForAllInventoryTransactions(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForAllInventoryTransactions(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}

func (r *InventoryRepository) createFilterSQLForAllInventoryTransactions(colID string, filterItem query.FilterItem) string {
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForAllInventoryTransactions(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForAllInventoryTransactions(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InventoryRepository) createNumberFilterSQLForAllInventoryTransactions(colID string, filterItem query.FilterItem) string {
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

func (r *InventoryRepository) createTextFilterSQLForAllInventoryTransactions(colID string, filterModel query.FilterItem) string {
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
func (r *InventoryRepository) createGroupBySQLForAllInventoryTransactions(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createOrderBySQLForAllInventoryTransactions(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createLimitAndOffsetSQLForAllInventoryTransactions(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InventoryRepository) createFilterSQLForAllInventoryTransactions(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForAllInventoryTransactions(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForAllInventoryTransactions(colID, filterItem)
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
