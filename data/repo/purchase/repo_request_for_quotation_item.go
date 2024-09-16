// Code generated by MVRP Codegen Util. DO NOT EDIT.

package purchase

import (
	"context"
	"database/sql"
	"mvrp/data/model/purchase"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *PurchaseRepository) ListAllRequestForQuotationItems(ctx context.Context, exec boil.ContextExecutor) (purchase.RequestForQuotationItemSlice, error) {
	return purchase.RequestForQuotationItems().All(ctx, exec)
}

func (r *PurchaseRepository) GetRequestForQuotationItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*purchase.RequestForQuotationItem, error) {
	return purchase.FindRequestForQuotationItem(ctx, exec, id)
}

func (r *PurchaseRepository) CreateRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *PurchaseRepository) UpdateRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *PurchaseRepository) UpsertRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *PurchaseRepository) DeleteRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor, m *purchase.RequestForQuotationItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *PurchaseRepository) RequestForQuotationItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return purchase.RequestForQuotationItemExists(ctx, exec, id)
}

func (r *PurchaseRepository) GetRequestForQuotationItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotationItems().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) GetMostRecentRequestForQuotationItem(ctx context.Context, exec boil.ContextExecutor) (*purchase.RequestForQuotationItem, error) {
	return purchase.RequestForQuotationItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *PurchaseRepository) GetNextEntryRequestForQuotationItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := purchase.RequestForQuotationItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *PurchaseRepository) GetRequestForQuotationItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := purchase.RequestForQuotationItems().Count(ctx, exec)
	return int(count), err
}

func (r *PurchaseRepository) SearchRequestForQuotationItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchRequestForQuotationItemDTO) (purchase.RequestForQuotationItemSlice, int, error) {
	return r.BuildSearchQueryForRequestForQuotationItems(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return purchase.RequestForQuotationItems(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *PurchaseRepository) BuildSearchQueryForRequestForQuotationItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchRequestForQuotationItemDTO) (purchase.RequestForQuotationItemSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForRequestForQuotationItems(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForRequestForQuotationItems(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForRequestForQuotationItems(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForRequestForQuotationItems(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForRequestForQuotationItems(request)
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

	count, err := purchase.RequestForQuotationItems(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := purchase.RequestForQuotationItems(queryMods...).All(ctx, exec)
	return res, int(count), err
}

/*
// 1. Create Select SQL
func (r *PurchaseRepository) createSelectSQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
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
func (r *PurchaseRepository) createSelectSQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForRequestForQuotationItems(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsForRequestForQuotationItems(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForRequestForQuotationItems(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

/*
// 2. Create Where SQL (handle group keys and filters)
func (r *PurchaseRepository) createWhereSQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForRequestForQuotationItems(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}
*/

// 2. Create Where SQL (handle group keys and filters)
func (r *PurchaseRepository) createWhereSQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsForRequestForQuotationItems(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLForRequestForQuotationItems(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *PurchaseRepository) handleNestedFieldsForRequestForQuotationItems(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *PurchaseRepository) createFilterSQLForRequestForQuotationItems(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForRequestForQuotationItems(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForRequestForQuotationItems(colID, filterItem)
	default:
		return "true"
	}
}

func (r *PurchaseRepository) createNumberFilterSQLForRequestForQuotationItems(colID string, filterItem query.FilterItem) string {
    // Handle nested JSON fields
    // if strings.Contains(colID, ".") {
    //     parts := strings.Split(colID, ".")
    //     jsonField := parts[0]
    //     nestedField := parts[1]
    //     colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    // }

	/*
		// Properly format the filter value
		filterValue := filterItem.Filter
		if filterItem.FilterType == "string" {
			filterValue = fmt.Sprintf("'%s'", filterItem.Filter)
		}
	*/

	// Handle filter operator and conditions (recursive)
	if filterItem.Operator != "" {
		conditions := filterItem.Conditions
		switch filterItem.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createNumberFilterSQLForRequestForQuotationItems(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createNumberFilterSQLForRequestForQuotationItems(colID, condition))
			}
			return strings.Join(orParts, " OR ")
		default:
			return "false"
		}
	}

	// Basic filter handling
    switch filterItem.Type {
    case "equals":
        return fmt.Sprintf("%s = %v", colID, filterItem.Filter)
	case "notEqual":
		return fmt.Sprintf("%s != %v", colID, filterItem.Filter)
    case "greaterThan":
        return fmt.Sprintf("%s > %v", colID, filterItem.Filter)
	case "greaterThanOrEqual":
		return fmt.Sprintf("%s >= %v", colID, filterItem.Filter)
    case "lessThan":
        return fmt.Sprintf("%s < %v", colID, filterItem.Filter)
	case "lessThanOrEqual":
		return fmt.Sprintf("%s <= %v", colID, filterItem.Filter)
	case "inRange":
		return fmt.Sprintf("%s BETWEEN %v AND %v", colID, filterItem.Filter, filterItem.FilterTo)
	case "blank":
		return fmt.Sprintf("%s IS NULL", colID)
	case "notBlank":
		return fmt.Sprintf("%s IS NOT NULL", colID)
    default:
        return "false"
    }
}

func (r *PurchaseRepository) createTextFilterSQLForRequestForQuotationItems(colID string, filterModel query.FilterItem) string {
	// Handle nested JSON fields
    // if strings.Contains(colID, ".") {
    //     parts := strings.Split(colID, ".")
    //     jsonField := parts[0]
    //     nestedField := parts[1]
    //     colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    // }

	// Handle filter operator and conditions (recursive)
	if filterModel.Operator != "" {
		conditions := filterModel.Conditions
		switch filterModel.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createTextFilterSQLForRequestForQuotationItems(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createTextFilterSQLForRequestForQuotationItems(colID, condition))
			}
			return strings.Join(orParts, " OR ")
		default:
			return "false"
		}
	}

	// Basic filter handling
	switch filterModel.Type {
	case "contains":
		return fmt.Sprintf("%s ILIKE '%%%s%%'", colID, filterModel.Filter)
	case "notContains":
		return fmt.Sprintf("%s NOT ILIKE '%%%s%%'", colID, filterModel.Filter)
	case "equals":
		return fmt.Sprintf("%s = '%s'", colID, filterModel.Filter)
	case "notEqual":
		return fmt.Sprintf("%s != '%s'", colID, filterModel.Filter)
	case "startsWith":
		return fmt.Sprintf("%s ILIKE '%s%%'", colID, filterModel.Filter)
	case "endsWith":
		return fmt.Sprintf("%s ILIKE '%%%s'", colID, filterModel.Filter)
	case "blank":
		return fmt.Sprintf("%s IS NULL", colID)
	case "notBlank":
		return fmt.Sprintf("%s IS NOT NULL", colID)
	default:
		return "false"
	}
}

// 3. Create Group By SQL
func (r *PurchaseRepository) createGroupBySQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
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
func (r *PurchaseRepository) createOrderBySQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
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
func (r *PurchaseRepository) createOrderBySQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) string {
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
func (r *PurchaseRepository) createLimitAndOffsetSQLForRequestForQuotationItems(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *PurchaseRepository) createFilterSQLForRequestForQuotationItems(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForRequestForQuotationItems(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForRequestForQuotationItems(colID, filterItem)
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