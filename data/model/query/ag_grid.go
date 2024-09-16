package query

type IServerSideGetRowsRequest struct {
	RowGroupCols []RowGroupCol `json:"rowGroupCols"`
	ValueCols    []ValueCol    `json:"valueCols"`
	GroupKeys    []string      `json:"groupKeys"`
	FilterModel  FilterModel   `json:"filterModel"`
	SortModel    []SortModel   `json:"sortModel"`
	StartRow     int           `json:"startRow"`
	EndRow       int           `json:"endRow"`
}

type RowGroupCol struct {
	ID    string `json:"id"`
	Field string `json:"field"`
}

type ValueCol struct {
	Field   string `json:"field"`
	AggFunc string `json:"aggFunc"`
}

type FilterModel map[string]FilterItem

type FilterItem struct {
	FilterType string      `json:"filterType"`
	Type       string      `json:"type"`
	Filter     interface{} `json:"filter"`
	FilterTo   interface{} `json:"filterTo,omitempty"`

	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`

	Operator   string       `json:"operator,omitempty"`
	Conditions []FilterItem `json:"conditions,omitempty"`
}

type SortModel struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"` // "asc" or "desc"
}
