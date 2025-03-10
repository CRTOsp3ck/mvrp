// Code generated by MVRP Codegen Util. DO NOT EDIT.

package {{ .Package }}

import (
	"context"
	"database/sql"
	"mvrp/data/model/{{ .Package }}"
	{{- if .HasSearchDTO }}
	"mvrp/data/model/query"
	"mvrp/domain/dto"
	"strings"
	"fmt"
	{{- end }}
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *{{ .Package | ToPascalCase }}Repository) ListAll{{ .PluralModelName }}(ctx context.Context, exec boil.ContextExecutor) ({{ .Package }}.{{ .ModelName }}Slice, error) {
	return {{ .Package }}.{{ .PluralModelName }}().All(ctx, exec)
}

func (r *{{ .Package | ToPascalCase }}Repository) Get{{ .ModelName }}ByID(ctx context.Context, exec boil.ContextExecutor, id int) (*{{ .Package }}.{{ .ModelName }}, error) {
	return {{ .Package }}.Find{{ .ModelName }}(ctx, exec, id)
}

func (r *{{ .Package | ToPascalCase }}Repository) Create{{ .ModelName }}(ctx context.Context, exec boil.ContextExecutor, m *{{ .Package }}.{{ .ModelName }}) error {
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *{{ .Package | ToPascalCase }}Repository) Update{{ .ModelName }}(ctx context.Context, exec boil.ContextExecutor, m *{{ .Package }}.{{ .ModelName }}) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *{{ .Package | ToPascalCase }}Repository) Upsert{{ .ModelName }}(ctx context.Context, exec boil.ContextExecutor, m *{{ .Package }}.{{ .ModelName }}) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *{{ .Package | ToPascalCase }}Repository) Delete{{ .ModelName }}(ctx context.Context, exec boil.ContextExecutor, m *{{ .Package }}.{{ .ModelName }}) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *{{ .Package | ToPascalCase }}Repository) {{ .ModelName }}Exists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return {{ .Package }}.{{ .ModelName }}Exists(ctx, exec, id)
}

func (r *{{ .Package | ToPascalCase }}Repository) Get{{ .ModelName }}RowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := {{ .Package }}.{{ .PluralModelName }}().Count(ctx, exec)
	return int(count), err
}

func (r *{{ .Package | ToPascalCase }}Repository) GetMostRecent{{ .ModelName }}(ctx context.Context, exec boil.ContextExecutor) (*{{ .Package }}.{{ .ModelName }}, error) {
	return {{ .Package }}.{{ .PluralModelName }}(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *{{ .Package | ToPascalCase }}Repository) GetNextEntry{{ .ModelName }}ID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := {{ .Package }}.{{ .PluralModelName }}(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil
}

func (r *{{ .Package | ToPascalCase }}Repository) Get{{ .ModelName }}TotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := {{ .Package }}.{{ .PluralModelName }}().Count(ctx, exec)
	return int(count), err
}

{{- if .HasSearchDTO }}

func (r *{{ .Package | ToPascalCase }}Repository) Search{{ .PluralModelName }}(ctx context.Context, exec boil.ContextExecutor, dto dto.Search{{ .ModelName }}DTO) ({{ .Package }}.{{ .ModelName }}Slice, int, error) {
	return r.BuildSearchQueryFor{{ .PluralModelName }}(ctx, exec, dto)

	/*
		var queryMods []qm.QueryMod
		{{- range .GroupQueryFields }}
		if dto.{{ .Name | ToPascalCase }} != "" {
			queryMods = append(queryMods, qm.Where("{{ .Name }} = ?", dto.{{ .Name | ToPascalCase }}))
		}
		{{- end }}
		{{- if .SearchQueryStatement }}
		queryMods = append(queryMods, qm.{{.SearchQueryStatement}})
		{{- end }}
		queryMods = append(queryMods,
			qm.Limit(dto.ItemsPerPage),
			qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
			// qm.GroupBy("id"),
			qm.OrderBy(dto.OrderBy+" "+"ASC"),
		)
		return {{ .Package }}.{{ .PluralModelName }}(queryMods...).All(ctx, exec)
	*/
}

/*
	AG-Grid Server-Side Row Model Queries
*/

func (r *{{ .Package | ToPascalCase }}Repository) BuildSearchQueryFor{{ .PluralModelName }}(ctx context.Context, exec boil.ContextExecutor, dto dto.Search{{ .ModelName }}DTO) ({{ .Package }}.{{ .ModelName }}Slice, int, error) {
	var queryMods []qm.QueryMod

	{{- range .GroupQueryFields }}
	if dto.{{ .Name | ToPascalCase }} != "" {
		queryMods = append(queryMods, qm.Where("{{ .Name }} = ?", dto.{{ .Name | ToPascalCase }}))
	}
	{{- end }}

	request := dto.IServerSideGetRowsRequest

	selectSQL := r.createSelectSQLFor{{ .PluralModelName }}(request)
	if selectSQL != "" {
		queryMods = append(queryMods, qm.Select(selectSQL))
	}

	whereSQL := r.createWhereSQLFor{{ .PluralModelName }}(request)
	if whereSQL != "" {
		queryMods = append(queryMods, qm.Where(whereSQL))
	}

	groupBySQL := r.createGroupBySQLFor{{ .PluralModelName }}(request)
	if groupBySQL != "" {
		queryMods = append(queryMods, qm.GroupBy(groupBySQL))
	}

	orderBySQL := r.createOrderBySQLFor{{ .PluralModelName }}(request)
	if orderBySQL != "" {
		queryMods = append(queryMods, qm.OrderBy(orderBySQL))
	}

	limitSQL, offsetSQL := r.createLimitAndOffsetSQLFor{{ .PluralModelName }}(request)
	if limitSQL > 0 {
		queryMods = append(queryMods, qm.Limit(limitSQL))
	}
	if offsetSQL > 0 {
		queryMods = append(queryMods, qm.Offset(offsetSQL))
	}

	// ---------------- Pagination Count Query Mods ---------------- 
	var countQueryMods []qm.QueryMod
	{{- range .GroupQueryFields }}
	if dto.{{ .Name | ToPascalCase }} != "" {
		countQueryMods = append(countQueryMods, qm.Where("{{ .Name }} = ?", dto.{{ .Name | ToPascalCase }}))
	}
	{{- end }}
	// --------------------------------------------------------------

	if whereSQL != "" {
		countQueryMods = append(countQueryMods, qm.Where(whereSQL))
	}

	count, err := {{ .Package }}.{{ .PluralModelName }}(countQueryMods...).Count(ctx, exec)
	if err != nil {
		return nil, 0, err
	}

	res, err := {{ .Package }}.{{ .PluralModelName }}(queryMods...).All(ctx, exec)
	return res, int(count), err
}

// 1. Create Select SQL
func (r *{{ .Package | ToPascalCase }}Repository) createSelectSQLFor{{ .PluralModelName }}(request query.IServerSideGetRowsRequest) string {
    rowGroupCols := request.RowGroupCols
    valueCols := request.ValueCols
    groupKeys := request.GroupKeys

    var colsToSelect []string

    if len(rowGroupCols) > len(groupKeys) {
        // Include group columns in SELECT
        for i := 0; i <= len(groupKeys) && i < len(rowGroupCols); i++ {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsFor{{ .PluralModelName }}(rowGroupCols[i].Field))
        }

        // Include aggregated columns in SELECT
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, fmt.Sprintf("SUM(%s) AS %s", r.handleNestedFieldsFor{{ .PluralModelName }}(valueCol.Field), valueCol.Field))
        }
    } else {
        // If not grouping, select all columns directly
        for _, valueCol := range valueCols {
            colsToSelect = append(colsToSelect, r.handleNestedFieldsFor{{ .PluralModelName }}(valueCol.Field))
        }
    }

    return strings.Join(colsToSelect, ", ")
}

// 2. Create Where SQL (handle group keys and filters)
func (r *{{ .Package | ToPascalCase }}Repository) createWhereSQLFor{{ .PluralModelName }}(request query.IServerSideGetRowsRequest) string {
    var whereParts []string

    // Handle group keys (if any)
    for i, groupKey := range request.GroupKeys {
        colName := request.RowGroupCols[i].Field
        whereParts = append(whereParts, fmt.Sprintf("%s = '%s'", r.handleNestedFieldsFor{{ .PluralModelName }}(colName), groupKey))
    }

    // Handle filter model (apply filters to each column)
    for colID, filterItem := range request.FilterModel {
        whereParts = append(whereParts, r.createFilterSQLFor{{ .PluralModelName }}(colID, filterItem))
    }

    if len(whereParts) > 0 {
        return strings.Join(whereParts, " AND ")
    }
    return ""
}

// Helper function to handle nested fields (splitting JSON fields)
func (r *{{ .Package | ToPascalCase }}Repository) handleNestedFieldsFor{{ .PluralModelName }}(field string) string {
    parts := strings.Split(field, ".")
    
    if len(parts) == 1 {
        // Not a nested field, return as is
        return parts[0]
    }
    
    // If it's a nested field, treat it as JSONB (e.g., receipient.name -> receipient->>'name')
    return fmt.Sprintf("%s->>'%s'", parts[0], parts[1])
}


func (r *{{ .Package | ToPascalCase }}Repository) createFilterSQLFor{{ .PluralModelName }}(colID string, filterItem query.FilterItem) string {
	// Handle nested JSON fields
	if strings.Contains(colID, ".") {
        parts := strings.Split(colID, ".")
        jsonField := parts[0]
        nestedField := parts[1]
        colID = fmt.Sprintf("%s->>'%s'", jsonField, nestedField)
    }
	
	switch filterItem.FilterType {
	case "text":
		return r.createTextFilterSQLFor{{ .PluralModelName }}(colID, filterItem)
	case "number":
		// Cast the JSONB text value to a number
		if strings.Contains(colID, "->>'") {
			colID = fmt.Sprintf("(%s)::numeric", colID)
		}
		return r.createNumberFilterSQLFor{{ .PluralModelName }}(colID, filterItem)
	case "date":
		return r.createDateFilterSQLFor{{ .PluralModelName }}(colID, filterItem)
	default:
		return "true"
	}
}

func (r *{{ .Package | ToPascalCase }}Repository) createNumberFilterSQLFor{{ .PluralModelName }}(colID string, filterItem query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterItem.Operator != "" {
		conditions := filterItem.Conditions
		switch filterItem.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createNumberFilterSQLFor{{ .PluralModelName }}(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createNumberFilterSQLFor{{ .PluralModelName }}(colID, condition))
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

func (r *{{ .Package | ToPascalCase }}Repository) createTextFilterSQLFor{{ .PluralModelName }}(colID string, filterModel query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterModel.Operator != "" {
		conditions := filterModel.Conditions
		switch filterModel.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createTextFilterSQLFor{{ .PluralModelName }}(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createTextFilterSQLFor{{ .PluralModelName }}(colID, condition))
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

func (r *{{ .Package | ToPascalCase }}Repository) createDateFilterSQLFor{{ .PluralModelName }}(colID string, filterModel query.FilterItem) string {
	// Handle filter operator and conditions (recursive)
	if filterModel.Operator != "" {
		conditions := filterModel.Conditions
		switch filterModel.Operator {
		case "AND":
			var andParts []string
			for _, condition := range conditions {
				andParts = append(andParts, r.createDateFilterSQLFor{{ .PluralModelName }}(colID, condition))
			}
			return strings.Join(andParts, " AND ")
		case "OR":
			var orParts []string
			for _, condition := range conditions {
				orParts = append(orParts, r.createDateFilterSQLFor{{ .PluralModelName }}(colID, condition))
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
func (r *{{ .Package | ToPascalCase }}Repository) createGroupBySQLFor{{ .PluralModelName }}(request query.IServerSideGetRowsRequest) string {
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
func (r *{{ .Package | ToPascalCase }}Repository) createOrderBySQLFor{{ .PluralModelName }}(request query.IServerSideGetRowsRequest) string {
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
func (r *{{ .Package | ToPascalCase }}Repository) createLimitAndOffsetSQLFor{{ .PluralModelName }}(request query.IServerSideGetRowsRequest) (limit int, offset int) {
	limit = request.EndRow - request.StartRow
	offset = request.StartRow
	return
}

// func (r *{{ .Package | ToPascalCase }}Repository) createFilterSQLFor{{ .PluralModelName }}(colID string, filterItem query.FilterItem) string {
// 	// Check if the filter column is a number or text
// 	if isNumericColumn(colID) {
// 		return r.createNumberFilterSQLFor{{ .PluralModelName }}(colID, filterItem)
// 	} else {
// 		return r.createTextFilterSQLFor{{ .PluralModelName }}(colID, filterItem)
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

{{- end }}