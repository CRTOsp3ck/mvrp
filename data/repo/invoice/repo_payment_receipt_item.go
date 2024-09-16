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

func (r *InvoiceRepository) ListAllPaymentReceiptItems(ctx context.Context, exec boil.ContextExecutor) (invoice.PaymentReceiptItemSlice, error) {
	return invoice.PaymentReceiptItems().All(ctx, exec)
}

func (r *InvoiceRepository) GetPaymentReceiptItemByID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.PaymentReceiptItem, error) {
	return invoice.FindPaymentReceiptItem(ctx, exec, id)
}

func (r *InvoiceRepository) CreatePaymentReceiptItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceiptItem) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InvoiceRepository) UpdatePaymentReceiptItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceiptItem) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InvoiceRepository) UpsertPaymentReceiptItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceiptItem) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InvoiceRepository) DeletePaymentReceiptItem(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceiptItem) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InvoiceRepository) PaymentReceiptItemExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return invoice.PaymentReceiptItemExists(ctx, exec, id)
}

func (r *InvoiceRepository) GetPaymentReceiptItemRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.PaymentReceiptItems().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) GetMostRecentPaymentReceiptItem(ctx context.Context, exec boil.ContextExecutor) (*invoice.PaymentReceiptItem, error) {
	return invoice.PaymentReceiptItems(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InvoiceRepository) GetNextEntryPaymentReceiptItemID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := invoice.PaymentReceiptItems(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *InvoiceRepository) GetPaymentReceiptItemTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.PaymentReceiptItems().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) SearchPaymentReceiptItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchPaymentReceiptItemDTO) (invoice.PaymentReceiptItemSlice, int, error) {
	return r.BuildSearchQueryForPaymentReceiptItems(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return invoice.PaymentReceiptItems(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InvoiceRepository) BuildSearchQueryForPaymentReceiptItems(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchPaymentReceiptItemDTO) (invoice.PaymentReceiptItemSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForPaymentReceiptItems(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForPaymentReceiptItems(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForPaymentReceiptItems(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForPaymentReceiptItems(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForPaymentReceiptItems(request)
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

	count, err := invoice.PaymentReceiptItems(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := invoice.PaymentReceiptItems(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *InvoiceRepository) createSelectSQLForPaymentReceiptItems(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForPaymentReceiptItems(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsForPaymentReceiptItems(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsForPaymentReceiptItems(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

// 2. Create Where SQL (handle group keys and filters)
func (r *InvoiceRepository) createWhereSQLForPaymentReceiptItems(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsForPaymentReceiptItems(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLForPaymentReceiptItems(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *InvoiceRepository) handleNestedFieldsForPaymentReceiptItems(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *InvoiceRepository) createFilterSQLForPaymentReceiptItems(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForPaymentReceiptItems(colID, filterItem)
	case "number":
		// Cast the JSONB text value to a number
		if strings.Contains(colID, "->>'") {
			colID = fmt.Sprintf("(%s)::numeric", colID)
		}
		return r.createNumberFilterSQLForPaymentReceiptItems(colID, filterItem)
	case "date":
		return r.createDateFilterSQLForPaymentReceiptItems(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InvoiceRepository) createNumberFilterSQLForPaymentReceiptItems(colID string, filterItem query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterItem.Operator != "" {
		conditions := filterItem.Conditions
		switch filterItem.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createNumberFilterSQLForPaymentReceiptItems(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createNumberFilterSQLForPaymentReceiptItems(colID, condition))
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

func (r *InvoiceRepository) createTextFilterSQLForPaymentReceiptItems(colID string, filterModel query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterModel.Operator != "" {
		conditions := filterModel.Conditions
		switch filterModel.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createTextFilterSQLForPaymentReceiptItems(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createTextFilterSQLForPaymentReceiptItems(colID, condition))
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

func (r *InvoiceRepository) createDateFilterSQLForPaymentReceiptItems(colID string, filterModel query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterModel.Operator != "" {
		conditions := filterModel.Conditions
		switch filterModel.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createDateFilterSQLForPaymentReceiptItems(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createDateFilterSQLForPaymentReceiptItems(colID, condition))
			}
			return strings.Join(orParts, " OR ")
		default:
			return "false"
		}
	}

	// Basic filter handling
	switch filterModel.Type {
	case "equals":
		return fmt.Sprintf("DATE(%s) = DATE('%s')", colID, filterModel.DateFrom)
	case "notEqual":
		return fmt.Sprintf("DATE(%s) != DATE('%s')", colID, filterModel.DateFrom)
	case "greaterThan":
		return fmt.Sprintf("DATE(%s) > DATE('%s')", colID, filterModel.DateFrom)
	// case "greaterThanOrEqual":
	// 	return fmt.Sprintf("DATE(%s) >= DATE('%s')", colID, filterModel.DateFrom)
	case "lessThan":
		return fmt.Sprintf("DATE(%s) < DATE('%s')", colID, filterModel.DateFrom)
	// case "lessThanOrEqual":
	// 	return fmt.Sprintf("DATE(%s) <= DATE('%s')", colID, filterModel.DateFrom)
	case "inRange":
		return fmt.Sprintf("DATE(%s) BETWEEN DATE('%s') AND DATE('%s')", colID, filterModel.DateFrom, filterModel.DateTo)
	case "blank":
		return fmt.Sprintf("%s IS NULL", colID)
	case "notBlank":
		return fmt.Sprintf("%s IS NOT NULL", colID)
	default:
		return "false"
	}
}

// 3. Create Group By SQL
func (r *InvoiceRepository) createGroupBySQLForPaymentReceiptItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createOrderBySQLForPaymentReceiptItems(request query.IServerSideGetRowsRequest) string {
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
func (r *InvoiceRepository) createLimitAndOffsetSQLForPaymentReceiptItems(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InvoiceRepository) createFilterSQLForPaymentReceiptItems(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForPaymentReceiptItems(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForPaymentReceiptItems(colID, filterItem)
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