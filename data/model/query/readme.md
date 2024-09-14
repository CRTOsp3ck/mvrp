Sure! Based on the schema you provided earlier, I will convert the relevant parts into Golang structs. These structs will represent the request and filter models that you can use in your query-building logic.

### Go Structs for AG Grid Request

#### `IServerSideGetRowsRequest`
This struct will capture the data that is sent from AG Grid, including row grouping, sorting, filtering, and pagination.

```go
type IServerSideGetRowsRequest struct {
    RowGroupCols []RowGroupCol `json:"rowGroupCols"`
    ValueCols    []ValueCol    `json:"valueCols"`
    GroupKeys    []string      `json:"groupKeys"`
    FilterModel  FilterModel   `json:"filterModel"`
    SortModel    []SortModel   `json:"sortModel"`
    StartRow     int           `json:"startRow"`
    EndRow       int           `json:"endRow"`
}
```

#### `RowGroupCol`
Represents the row group columns, which are used for grouping the data.

```go
type RowGroupCol struct {
    ID    string `json:"id"`
    Field string `json:"field"`
}
```

#### `ValueCol`
Represents the value columns for aggregation functions (e.g., sum, count, etc.).

```go
type ValueCol struct {
    Field   string `json:"field"`
    AggFunc string `json:"aggFunc"`
}
```

#### `FilterModel`
Represents the filter model, which contains the filters applied to each column.

```go
type FilterModel map[string]FilterItem

type FilterItem struct {
    FilterType string `json:"filterType"`
    Type       string `json:"type"`
    Filter     string `json:"filter"`
    FilterTo   string `json:"filterTo,omitempty"`
}
```

The `FilterModel` is a map where the key is the column name and the value is a `FilterItem`, which represents the filter applied to that column.

#### `SortModel`
Represents the sort model, which defines how data is sorted.

```go
type SortModel struct {
    ColId string `json:"colId"`
    Sort  string `json:"sort"` // "asc" or "desc"
}
```

### Example Struct Usage

With these structs, you can now deserialize JSON requests from AG Grid into Go structs. Here's an example of how you could deserialize an AG Grid request:

```go
import (
    "encoding/json"
    "net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    var request IServerSideGetRowsRequest

    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Now you can pass `request` to your query builder
    entities, err := repository.BuildSearchQuery(context.Background(), request)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send the result back
    json.NewEncoder(w).Encode(entities)
}
```

### Full Example of All Structs

```go
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
    FilterType string `json:"filterType"`
    Type       string `json:"type"`
    Filter     string `json:"filter"`
    FilterTo   string `json:"filterTo,omitempty"`
}

type SortModel struct {
    ColId string `json:"colId"`
    Sort  string `json:"sort"` // "asc" or "desc"
}
```

This should give you a good structure to handle the AG Grid server-side request in your Go backend, covering grouping, filtering, sorting, and pagination. You can easily expand on these structs if AG Grid requires more fields in the future.