// Code generated by MVRP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) ListAllReturnMerchandiseAuthorizationViews(ctx context.Context, exec boil.ContextExecutor) (inventory.ReturnMerchandiseAuthorizationViewSlice, error) {
	return inventory.ReturnMerchandiseAuthorizationViews().All(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationViewByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.ReturnMerchandiseAuthorizationView, error) {
	return inventory.ReturnMerchandiseAuthorizationViews(qm.Where(inventory.ReturnMerchandiseAuthorizationViewColumns.ID+"=?", id)).One(ctx, exec)
}

func (r *InventoryRepository) ReturnMerchandiseAuthorizationViewExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.ReturnMerchandiseAuthorizationViews(qm.Where(inventory.ReturnMerchandiseAuthorizationViewColumns.ID+"=?", id)).Exists(ctx, exec)
}

func (r *InventoryRepository) GetReturnMerchandiseAuthorizationViewRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.ReturnMerchandiseAuthorizationViews().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentReturnMerchandiseAuthorizationView(ctx context.Context, exec boil.ContextExecutor) (*inventory.ReturnMerchandiseAuthorizationView, error) {
	return inventory.ReturnMerchandiseAuthorizationViews(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) SearchReturnMerchandiseAuthorizationViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchReturnMerchandiseAuthorizationDTO) (inventory.ReturnMerchandiseAuthorizationViewSlice, int, error) {
	return r.BuildSearchQueryForReturnMerchandiseAuthorizationViews(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return inventory.ReturnMerchandiseAuthorizationViews(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InventoryRepository) BuildSearchQueryForReturnMerchandiseAuthorizationViews(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchReturnMerchandiseAuthorizationDTO) (inventory.ReturnMerchandiseAuthorizationViewSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForReturnMerchandiseAuthorizationViews(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForReturnMerchandiseAuthorizationViews(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForReturnMerchandiseAuthorizationViews(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForReturnMerchandiseAuthorizationViews(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForReturnMerchandiseAuthorizationViews(request)
	if limitSQL > 0 {
		queryMods = append(queryMods, qm.Limit(limitSQL))
	}
	if offsetSQL > 0 {
		queryMods = append(queryMods, qm.Offset(offsetSQL))
	}

	// ---------------- Pagination Count Query Mods ---------------- 
	var countQueryMods []qm.QueryMod
	// --------------------------------------------------------------

	if whereSQL != "" {
		countQueryMods = append(countQueryMods, qm.Where(whereSQL))
	}

	count, err := inventory.ReturnMerchandiseAuthorizationViews(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := inventory.ReturnMerchandiseAuthorizationViews(queryMods...).All(ctx, exec)
	return res, int(count), err
}

/*
// 1. Create Select SQL
func (r *InventoryRepository) createSelectSQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
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
*/

// 1. Create Select SQL
func (r *InventoryRepository) createSelectSQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForReturnMerchandiseAuthorizationViews(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsForReturnMerchandiseAuthorizationViews(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForReturnMerchandiseAuthorizationViews(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

/*
// 2. Create Where SQL (handle group keys and filters)
func (r *InventoryRepository) createWhereSQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}
*/

// 2. Create Where SQL (handle group keys and filters)
func (r *InventoryRepository) createWhereSQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsForReturnMerchandiseAuthorizationViews(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *InventoryRepository) handleNestedFieldsForReturnMerchandiseAuthorizationViews(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *InventoryRepository) createFilterSQLForReturnMerchandiseAuthorizationViews(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InventoryRepository) createNumberFilterSQLForReturnMerchandiseAuthorizationViews(colID string, filterItem query.FilterItem) string {
    // Handle nested JSON fields
    if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }

    // Properly format the filter value
    filterValue := filterItem.Filter
    if filterItem.FilterType == "string" {
        filterValue = fmt.Sprintf("'%s'", filterItem.Filter)
    }

    switch filterItem.Type {
    case "equals":
        return fmt.Sprintf("%s = %s", colID, filterValue)
    case "greaterThan":
        return fmt.Sprintf("%s > %s", colID, filterValue)
    case "lessThan":
        return fmt.Sprintf("%s < %s", colID, filterValue)
    default:
        return "true"
    }
}

func (r *InventoryRepository) createTextFilterSQLForReturnMerchandiseAuthorizationViews(colID string, filterModel query.FilterItem) string {
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
func (r *InventoryRepository) createGroupBySQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
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

/*
// 4. Create Order By SQL
func (r *InventoryRepository) createOrderBySQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
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
		colID := sort.ColId
		// Handle nested JSON fields
        if strings.Contains(colID, ".") {
            parts := strings.Split(colID, ".")
            jsonField := parts[0]
            nestedField := parts[1]
            colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
        }

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
*/

// 4. Create Order By SQL
func (r *InventoryRepository) createOrderBySQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) string {
	sortModel := request.SortModel
	rowGroupCols := request.RowGroupCols
	groupKeys := request.GroupKeys

	var sortParts []string

	// Determine if we are doing grouping
	grouping := len(rowGroupCols) > len(groupKeys)
	groupColIds := make(map[string]string)

	// Create a map of grouped columns with dynamic field handling
	for i := 0; i < len(groupKeys)+1 && i < len(rowGroupCols); i++ {
		colID := rowGroupCols[i].Field
		if strings.Contains(colID, ".") {
			// Handle nested JSON fields for grouped columns
			parts := strings.Split(colID, ".")
			jsonField := parts[0]
			nestedField := parts[1]
			groupColIds[colID] = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
		} else {
			// Non-nested fields
			groupColIds[colID] = colID
		}
	}

	// Iterate over the sort model and construct the ORDER BY clause
	for _, sort := range sortModel {
		colID := sort.ColId

		// Handle nested JSON fields for sorting
		if strings.Contains(colID, ".") {
			parts := strings.Split(colID, ".")
			jsonField := parts[0]
			nestedField := parts[1]
			colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
		}

		if grouping {
			// Allow sorting only on grouped columns (with dynamic field handling)
			if groupCol, exists := groupColIds[sort.ColId]; exists {
				sortParts = append(sortParts, fmt.Sprintf("%s %s", groupCol, sort.Sort))
			}
		} else {
			// If not grouping, allow sorting on any column
			sortParts = append(sortParts, fmt.Sprintf("%s %s", colID, sort.Sort))
		}
	}

	if len(sortParts) > 0 {
		return strings.Join(sortParts, ", ")
	}

	return ""
}

// 5. Create Limit and Offset SQL
func (r *InventoryRepository) createLimitAndOffsetSQLForReturnMerchandiseAuthorizationViews(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InventoryRepository) createFilterSQLForReturnMerchandiseAuthorizationViews(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForReturnMerchandiseAuthorizationViews(colID, filterItem)
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