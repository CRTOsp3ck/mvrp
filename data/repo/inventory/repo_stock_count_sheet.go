// Code generated by MVRP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"database/sql"
	"mvrp/data/model/inventory"
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InventoryRepository) ListAllStockCountSheets(ctx context.Context, exec boil.ContextExecutor) (inventory.StockCountSheetSlice, error) {
	return inventory.StockCountSheets().All(ctx, exec)
}

func (r *InventoryRepository) GetStockCountSheetByID(ctx context.Context, exec boil.ContextExecutor, id int) (*inventory.StockCountSheet, error) {
	return inventory.FindStockCountSheet(ctx, exec, id)
}

func (r *InventoryRepository) CreateStockCountSheet(ctx context.Context, exec boil.ContextExecutor, m *inventory.StockCountSheet) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InventoryRepository) UpdateStockCountSheet(ctx context.Context, exec boil.ContextExecutor, m *inventory.StockCountSheet) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InventoryRepository) UpsertStockCountSheet(ctx context.Context, exec boil.ContextExecutor, m *inventory.StockCountSheet) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InventoryRepository) DeleteStockCountSheet(ctx context.Context, exec boil.ContextExecutor, m *inventory.StockCountSheet) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InventoryRepository) StockCountSheetExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return inventory.StockCountSheetExists(ctx, exec, id)
}

func (r *InventoryRepository) GetStockCountSheetRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.StockCountSheets().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) GetMostRecentStockCountSheet(ctx context.Context, exec boil.ContextExecutor) (*inventory.StockCountSheet, error) {
	return inventory.StockCountSheets(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InventoryRepository) GetNextEntryStockCountSheetID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := inventory.StockCountSheets(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *InventoryRepository) GetStockCountSheetTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := inventory.StockCountSheets().Count(ctx, exec)
	return int(count), err
}

func (r *InventoryRepository) SearchStockCountSheets(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchStockCountSheetDTO) (inventory.StockCountSheetSlice, int, error) {
	return r.BuildSearchQueryForStockCountSheets(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return inventory.StockCountSheets(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *InventoryRepository) BuildSearchQueryForStockCountSheets(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchStockCountSheetDTO) (inventory.StockCountSheetSlice, int, error) {
	var queryMods []qm.QueryMod

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLForStockCountSheets(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLForStockCountSheets(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLForStockCountSheets(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLForStockCountSheets(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLForStockCountSheets(request)
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

	count, err := inventory.StockCountSheets(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := inventory.StockCountSheets(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *InventoryRepository) createSelectSQLForStockCountSheets(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createWhereSQLForStockCountSheets(request query.IServerSideGetRowsRequest) string {
	var whereParts []string

	// Handle group keys (if any)
	for i, groupKey := range request.GroupKeys {
		colName := request.RowGroupCols[i].Field
		whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", colName, groupKey))
	}

	// Handle filter model (apply filters to each column)
	for colID, filterItem := range request.FilterModel {
		whereParts = append(whereParts, r.createFilterSQLForStockCountSheets(colID, filterItem))
	}

	if len(whereParts) > 0 {
		return strings.Join(whereParts, " AND ")
	}
	return ""
}

func (r *InventoryRepository) createFilterSQLForStockCountSheets(colID string, filterItem query.FilterItem) string {
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLForStockCountSheets(colID, filterItem)
	case "number":
		return r.createNumberFilterSQLForStockCountSheets(colID, filterItem)
	default:
		return "true"
	}
}

func (r *InventoryRepository) createNumberFilterSQLForStockCountSheets(colID string, filterItem query.FilterItem) string {
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

func (r *InventoryRepository) createTextFilterSQLForStockCountSheets(colID string, filterModel query.FilterItem) string {
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
func (r *InventoryRepository) createGroupBySQLForStockCountSheets(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createOrderBySQLForStockCountSheets(request query.IServerSideGetRowsRequest) string {
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
func (r *InventoryRepository) createLimitAndOffsetSQLForStockCountSheets(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *InventoryRepository) createFilterSQLForStockCountSheets(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLForStockCountSheets(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLForStockCountSheets(colID, filterItem)
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