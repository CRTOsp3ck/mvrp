// Code generated by MVRP Codegen Util. DO NOT EDIT.

package invoice

import (
	"context"
	"database/sql"
	"mvrp/data/model/invoice"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) ListAllInvoiceItems(ctx context.Context, exec boil.ContextExecutor) (invoice.InvoiceItemSlice, error) {
	return invoice.InvoiceItems().All(ctx, exec)
}

func (r *InvoiceRepository) GetInvoiceItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.InvoiceItem, error) {
	return invoice.FindInvoiceItem(ctx, exec, id)
}

func (r *InvoiceRepository) CreateInvoiceItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.InvoiceItem) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InvoiceRepository) UpdateInvoiceItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.InvoiceItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InvoiceRepository) UpsertInvoiceItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.InvoiceItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InvoiceRepository) DeleteInvoiceItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.InvoiceItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InvoiceRepository) InvoiceItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return invoice.InvoiceItemExists(ctx, exec, id)
}

func (r *InvoiceRepository) GetInvoiceItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.InvoiceItems().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) GetMostRecentInvoiceItem(ctx context.Context, exec boil.ContextExecutor) (*invoice.InvoiceItem, error) {
	return invoice.InvoiceItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InvoiceRepository) GetNextEntryInvoiceItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := invoice.InvoiceItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *InvoiceRepository) GetInvoiceItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.InvoiceItems().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) SearchInvoiceItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInvoiceItemDTO) (invoice.InvoiceItemSlice, int, error) {
	return r.BuildSearchQueryForInvoiceItems(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return invoice.InvoiceItems(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InvoiceRepository) BuildSearchQueryForInvoiceItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInvoiceItemDTO) (invoice.InvoiceItemSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForInvoiceItems(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForInvoiceItems(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForInvoiceItems(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForInvoiceItems(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForInvoiceItems(request)
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

	count, err := invoice.InvoiceItems(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := invoice.InvoiceItems(queryMods...).All(ctx, exec)
	return res, int(count), err
}

/*
// 1. Create Select SQL
func (r *InvoiceRepository) createSelectSQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createSelectSQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForInvoiceItems(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsForInvoiceItems(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForInvoiceItems(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

/*
// 2. Create Where SQL (handle group keys and filters)
func (r *InvoiceRepository) createWhereSQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForInvoiceItems(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}
*/

// 2. Create Where SQL (handle group keys and filters)
func (r *InvoiceRepository) createWhereSQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsForInvoiceItems(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLForInvoiceItems(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *InvoiceRepository) handleNestedFieldsForInvoiceItems(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *InvoiceRepository) createFilterSQLForInvoiceItems(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForInvoiceItems(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForInvoiceItems(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InvoiceRepository) createNumberFilterSQLForInvoiceItems(colID string, filterItem query.FilterItem) string {
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

func (r *InvoiceRepository) createTextFilterSQLForInvoiceItems(colID string, filterModel query.FilterItem) string {
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
func (r *InvoiceRepository) createGroupBySQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createOrderBySQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createOrderBySQLForInvoiceItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createLimitAndOffsetSQLForInvoiceItems(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InvoiceRepository) createFilterSQLForInvoiceItems(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForInvoiceItems(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForInvoiceItems(colID, filterItem)
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