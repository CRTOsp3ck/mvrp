// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package purchase

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RequestForQuotationView is an object representing the database table.
type RequestForQuotationView struct {
	ID                        null.Int    `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	BaseDocumentID            null.Int    `boil:"base_document_id" json:"base_document_id,omitempty" toml:"base_document_id" yaml:"base_document_id,omitempty"`
	RequestForQuotationNumber null.String `boil:"request_for_quotation_number" json:"request_for_quotation_number,omitempty" toml:"request_for_quotation_number" yaml:"request_for_quotation_number,omitempty"`
	ValidUntilDate            null.Time   `boil:"valid_until_date" json:"valid_until_date,omitempty" toml:"valid_until_date" yaml:"valid_until_date,omitempty"`
	VendorID                  null.Int    `boil:"vendor_id" json:"vendor_id,omitempty" toml:"vendor_id" yaml:"vendor_id,omitempty"`
	CustomerID                null.Int    `boil:"customer_id" json:"customer_id,omitempty" toml:"customer_id" yaml:"customer_id,omitempty"`
	ShipToInformation         null.String `boil:"ship_to_information" json:"ship_to_information,omitempty" toml:"ship_to_information" yaml:"ship_to_information,omitempty"`
	RequestedBy               null.String `boil:"requested_by" json:"requested_by,omitempty" toml:"requested_by" yaml:"requested_by,omitempty"`
	BaseDocument              null.JSON   `boil:"base_document" json:"base_document,omitempty" toml:"base_document" yaml:"base_document,omitempty"`
	RequestForQuotationItems  null.JSON   `boil:"request_for_quotation_items" json:"request_for_quotation_items,omitempty" toml:"request_for_quotation_items" yaml:"request_for_quotation_items,omitempty"`
}

var RequestForQuotationViewColumns = struct {
	ID                        string
	BaseDocumentID            string
	RequestForQuotationNumber string
	ValidUntilDate            string
	VendorID                  string
	CustomerID                string
	ShipToInformation         string
	RequestedBy               string
	BaseDocument              string
	RequestForQuotationItems  string
}{
	ID:                        "id",
	BaseDocumentID:            "base_document_id",
	RequestForQuotationNumber: "request_for_quotation_number",
	ValidUntilDate:            "valid_until_date",
	VendorID:                  "vendor_id",
	CustomerID:                "customer_id",
	ShipToInformation:         "ship_to_information",
	RequestedBy:               "requested_by",
	BaseDocument:              "base_document",
	RequestForQuotationItems:  "request_for_quotation_items",
}

var RequestForQuotationViewTableColumns = struct {
	ID                        string
	BaseDocumentID            string
	RequestForQuotationNumber string
	ValidUntilDate            string
	VendorID                  string
	CustomerID                string
	ShipToInformation         string
	RequestedBy               string
	BaseDocument              string
	RequestForQuotationItems  string
}{
	ID:                        "request_for_quotation_view.id",
	BaseDocumentID:            "request_for_quotation_view.base_document_id",
	RequestForQuotationNumber: "request_for_quotation_view.request_for_quotation_number",
	ValidUntilDate:            "request_for_quotation_view.valid_until_date",
	VendorID:                  "request_for_quotation_view.vendor_id",
	CustomerID:                "request_for_quotation_view.customer_id",
	ShipToInformation:         "request_for_quotation_view.ship_to_information",
	RequestedBy:               "request_for_quotation_view.requested_by",
	BaseDocument:              "request_for_quotation_view.base_document",
	RequestForQuotationItems:  "request_for_quotation_view.request_for_quotation_items",
}

// Generated where

var RequestForQuotationViewWhere = struct {
	ID                        whereHelpernull_Int
	BaseDocumentID            whereHelpernull_Int
	RequestForQuotationNumber whereHelpernull_String
	ValidUntilDate            whereHelpernull_Time
	VendorID                  whereHelpernull_Int
	CustomerID                whereHelpernull_Int
	ShipToInformation         whereHelpernull_String
	RequestedBy               whereHelpernull_String
	BaseDocument              whereHelpernull_JSON
	RequestForQuotationItems  whereHelpernull_JSON
}{
	ID:                        whereHelpernull_Int{field: "\"purchase\".\"request_for_quotation_view\".\"id\""},
	BaseDocumentID:            whereHelpernull_Int{field: "\"purchase\".\"request_for_quotation_view\".\"base_document_id\""},
	RequestForQuotationNumber: whereHelpernull_String{field: "\"purchase\".\"request_for_quotation_view\".\"request_for_quotation_number\""},
	ValidUntilDate:            whereHelpernull_Time{field: "\"purchase\".\"request_for_quotation_view\".\"valid_until_date\""},
	VendorID:                  whereHelpernull_Int{field: "\"purchase\".\"request_for_quotation_view\".\"vendor_id\""},
	CustomerID:                whereHelpernull_Int{field: "\"purchase\".\"request_for_quotation_view\".\"customer_id\""},
	ShipToInformation:         whereHelpernull_String{field: "\"purchase\".\"request_for_quotation_view\".\"ship_to_information\""},
	RequestedBy:               whereHelpernull_String{field: "\"purchase\".\"request_for_quotation_view\".\"requested_by\""},
	BaseDocument:              whereHelpernull_JSON{field: "\"purchase\".\"request_for_quotation_view\".\"base_document\""},
	RequestForQuotationItems:  whereHelpernull_JSON{field: "\"purchase\".\"request_for_quotation_view\".\"request_for_quotation_items\""},
}

var (
	requestForQuotationViewAllColumns            = []string{"id", "base_document_id", "request_for_quotation_number", "valid_until_date", "vendor_id", "customer_id", "ship_to_information", "requested_by", "base_document", "request_for_quotation_items"}
	requestForQuotationViewColumnsWithoutDefault = []string{}
	requestForQuotationViewColumnsWithDefault    = []string{"id", "base_document_id", "request_for_quotation_number", "valid_until_date", "vendor_id", "customer_id", "ship_to_information", "requested_by", "base_document", "request_for_quotation_items"}
	requestForQuotationViewPrimaryKeyColumns     = []string{}
	requestForQuotationViewGeneratedColumns      = []string{}
)

type (
	// RequestForQuotationViewSlice is an alias for a slice of pointers to RequestForQuotationView.
	// This should almost always be used instead of []RequestForQuotationView.
	RequestForQuotationViewSlice []*RequestForQuotationView
	// RequestForQuotationViewHook is the signature for custom RequestForQuotationView hook methods
	RequestForQuotationViewHook func(context.Context, boil.ContextExecutor, *RequestForQuotationView) error

	requestForQuotationViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	requestForQuotationViewType           = reflect.TypeOf(&RequestForQuotationView{})
	requestForQuotationViewMapping        = queries.MakeStructMapping(requestForQuotationViewType)
	requestForQuotationViewInsertCacheMut sync.RWMutex
	requestForQuotationViewInsertCache    = make(map[string]insertCache)
	requestForQuotationViewUpdateCacheMut sync.RWMutex
	requestForQuotationViewUpdateCache    = make(map[string]updateCache)
	requestForQuotationViewUpsertCacheMut sync.RWMutex
	requestForQuotationViewUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
	// These are used in some views
	_ = fmt.Sprintln("")
	_ = reflect.Int
	_ = strings.Builder{}
	_ = sync.Mutex{}
	_ = strmangle.Plural("")
	_ = strconv.IntSize
)

var requestForQuotationViewAfterSelectMu sync.Mutex
var requestForQuotationViewAfterSelectHooks []RequestForQuotationViewHook

var requestForQuotationViewBeforeInsertMu sync.Mutex
var requestForQuotationViewBeforeInsertHooks []RequestForQuotationViewHook
var requestForQuotationViewAfterInsertMu sync.Mutex
var requestForQuotationViewAfterInsertHooks []RequestForQuotationViewHook

var requestForQuotationViewBeforeUpsertMu sync.Mutex
var requestForQuotationViewBeforeUpsertHooks []RequestForQuotationViewHook
var requestForQuotationViewAfterUpsertMu sync.Mutex
var requestForQuotationViewAfterUpsertHooks []RequestForQuotationViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RequestForQuotationView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestForQuotationViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RequestForQuotationView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestForQuotationViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RequestForQuotationView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestForQuotationViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RequestForQuotationView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestForQuotationViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RequestForQuotationView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestForQuotationViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRequestForQuotationViewHook registers your hook function for all future operations.
func AddRequestForQuotationViewHook(hookPoint boil.HookPoint, requestForQuotationViewHook RequestForQuotationViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		requestForQuotationViewAfterSelectMu.Lock()
		requestForQuotationViewAfterSelectHooks = append(requestForQuotationViewAfterSelectHooks, requestForQuotationViewHook)
		requestForQuotationViewAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		requestForQuotationViewBeforeInsertMu.Lock()
		requestForQuotationViewBeforeInsertHooks = append(requestForQuotationViewBeforeInsertHooks, requestForQuotationViewHook)
		requestForQuotationViewBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		requestForQuotationViewAfterInsertMu.Lock()
		requestForQuotationViewAfterInsertHooks = append(requestForQuotationViewAfterInsertHooks, requestForQuotationViewHook)
		requestForQuotationViewAfterInsertMu.Unlock()
	case boil.BeforeUpsertHook:
		requestForQuotationViewBeforeUpsertMu.Lock()
		requestForQuotationViewBeforeUpsertHooks = append(requestForQuotationViewBeforeUpsertHooks, requestForQuotationViewHook)
		requestForQuotationViewBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		requestForQuotationViewAfterUpsertMu.Lock()
		requestForQuotationViewAfterUpsertHooks = append(requestForQuotationViewAfterUpsertHooks, requestForQuotationViewHook)
		requestForQuotationViewAfterUpsertMu.Unlock()
	}
}

// One returns a single requestForQuotationView record from the query.
func (q requestForQuotationViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RequestForQuotationView, error) {
	o := &RequestForQuotationView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "purchase: failed to execute a one query for request_for_quotation_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RequestForQuotationView records from the query.
func (q requestForQuotationViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (RequestForQuotationViewSlice, error) {
	var o []*RequestForQuotationView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "purchase: failed to assign all query results to RequestForQuotationView slice")
	}

	if len(requestForQuotationViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RequestForQuotationView records in the query.
func (q requestForQuotationViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to count request_for_quotation_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q requestForQuotationViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "purchase: failed to check if request_for_quotation_view exists")
	}

	return count > 0, nil
}

// RequestForQuotationViews retrieves all the records using an executor.
func RequestForQuotationViews(mods ...qm.QueryMod) requestForQuotationViewQuery {
	mods = append(mods, qm.From("\"purchase\".\"request_for_quotation_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"purchase\".\"request_for_quotation_view\".*"})
	}

	return requestForQuotationViewQuery{q}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RequestForQuotationView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("purchase: no request_for_quotation_view provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestForQuotationViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	requestForQuotationViewInsertCacheMut.RLock()
	cache, cached := requestForQuotationViewInsertCache[key]
	requestForQuotationViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			requestForQuotationViewAllColumns,
			requestForQuotationViewColumnsWithDefault,
			requestForQuotationViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(requestForQuotationViewType, requestForQuotationViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(requestForQuotationViewType, requestForQuotationViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"purchase\".\"request_for_quotation_view\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"purchase\".\"request_for_quotation_view\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "purchase: unable to insert into request_for_quotation_view")
	}

	if !cached {
		requestForQuotationViewInsertCacheMut.Lock()
		requestForQuotationViewInsertCache[key] = cache
		requestForQuotationViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RequestForQuotationView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("purchase: no request_for_quotation_view provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestForQuotationViewColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	requestForQuotationViewUpsertCacheMut.RLock()
	cache, cached := requestForQuotationViewUpsertCache[key]
	requestForQuotationViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			requestForQuotationViewAllColumns,
			requestForQuotationViewColumnsWithDefault,
			requestForQuotationViewColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			requestForQuotationViewAllColumns,
			requestForQuotationViewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("purchase: unable to upsert request_for_quotation_view, could not build update column list")
		}

		ret := strmangle.SetComplement(requestForQuotationViewAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(requestForQuotationViewPrimaryKeyColumns) == 0 {
				return errors.New("purchase: unable to upsert request_for_quotation_view, could not build conflict column list")
			}

			conflict = make([]string, len(requestForQuotationViewPrimaryKeyColumns))
			copy(conflict, requestForQuotationViewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"purchase\".\"request_for_quotation_view\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(requestForQuotationViewType, requestForQuotationViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(requestForQuotationViewType, requestForQuotationViewMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "purchase: unable to upsert request_for_quotation_view")
	}

	if !cached {
		requestForQuotationViewUpsertCacheMut.Lock()
		requestForQuotationViewUpsertCache[key] = cache
		requestForQuotationViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
