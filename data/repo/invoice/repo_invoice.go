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

func (r *InvoiceRepository) ListAllInvoices(ctx context.Context, exec boil.ContextExecutor) (invoice.InvoiceSlice, error) {
	return invoice.Invoices().All(ctx, exec)
}

func (r *InvoiceRepository) GetInvoiceByID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.Invoice, error) {
	return invoice.FindInvoice(ctx, exec, id)
}

func (r *InvoiceRepository) CreateInvoice(ctx context.Context, exec boil.ContextExecutor, m *invoice.Invoice) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InvoiceRepository) UpdateInvoice(ctx context.Context, exec boil.ContextExecutor, m *invoice.Invoice) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InvoiceRepository) UpsertInvoice(ctx context.Context, exec boil.ContextExecutor, m *invoice.Invoice) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InvoiceRepository) DeleteInvoice(ctx context.Context, exec boil.ContextExecutor, m *invoice.Invoice) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InvoiceRepository) InvoiceExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return invoice.InvoiceExists(ctx, exec, id)
}

func (r *InvoiceRepository) GetInvoiceRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.Invoices().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) GetMostRecentInvoice(ctx context.Context, exec boil.ContextExecutor) (*invoice.Invoice, error) {
	return invoice.Invoices(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InvoiceRepository) GetNextEntryInvoiceID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := invoice.Invoices(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *InvoiceRepository) GetInvoiceTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.Invoices().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) SearchInvoices(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInvoiceDTO) (invoice.InvoiceSlice, int, error) {
	return r.BuildSearchQueryForInvoices(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return invoice.Invoices(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InvoiceRepository) BuildSearchQueryForInvoices(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchInvoiceDTO) (invoice.InvoiceSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForInvoices(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForInvoices(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForInvoices(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForInvoices(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForInvoices(request)
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

	count, err := invoice.Invoices(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := invoice.Invoices(queryMods...).All(ctx, exec)
	return res, int(count), err
}

/*
// 1. Create Select SQL
func (r *InvoiceRepository) createSelectSQLForInvoices(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createSelectSQLForInvoices(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForInvoices(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsForInvoices(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForInvoices(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

/*
// 2. Create Where SQL (handle group keys and filters)
func (r *InvoiceRepository) createWhereSQLForInvoices(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForInvoices(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}
*/

// 2. Create Where SQL (handle group keys and filters)
func (r *InvoiceRepository) createWhereSQLForInvoices(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsForInvoices(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLForInvoices(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *InvoiceRepository) handleNestedFieldsForInvoices(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *InvoiceRepository) createFilterSQLForInvoices(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForInvoices(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForInvoices(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InvoiceRepository) createNumberFilterSQLForInvoices(colID string, filterItem query.FilterItem) string {
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
				andParts = append(andParts, r.createNumberFilterSQLForInvoices(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createNumberFilterSQLForInvoices(colID, condition))
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

func (r *InvoiceRepository) createTextFilterSQLForInvoices(colID string, filterModel query.FilterItem) string {
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
				andParts = append(andParts, r.createTextFilterSQLForInvoices(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createTextFilterSQLForInvoices(colID, condition))
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
func (r *InvoiceRepository) createGroupBySQLForInvoices(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createOrderBySQLForInvoices(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createOrderBySQLForInvoices(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createLimitAndOffsetSQLForInvoices(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InvoiceRepository) createFilterSQLForInvoices(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForInvoices(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForInvoices(colID, filterItem)
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