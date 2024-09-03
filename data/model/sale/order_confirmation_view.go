// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sale

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

// OrderConfirmationView is an object representing the database table.
type OrderConfirmationView struct {
	ID                      null.Int    `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	BaseDocumentID          null.Int    `boil:"base_document_id" json:"base_document_id,omitempty" toml:"base_document_id" yaml:"base_document_id,omitempty"`
	OrderConfirmationNumber null.String `boil:"order_confirmation_number" json:"order_confirmation_number,omitempty" toml:"order_confirmation_number" yaml:"order_confirmation_number,omitempty"`
	SalesOrderID            null.Int    `boil:"sales_order_id" json:"sales_order_id,omitempty" toml:"sales_order_id" yaml:"sales_order_id,omitempty"`
	CustomerID              null.Int    `boil:"customer_id" json:"customer_id,omitempty" toml:"customer_id" yaml:"customer_id,omitempty"`
	ShipToInformation       null.String `boil:"ship_to_information" json:"ship_to_information,omitempty" toml:"ship_to_information" yaml:"ship_to_information,omitempty"`
	BaseDocument            null.JSON   `boil:"base_document" json:"base_document,omitempty" toml:"base_document" yaml:"base_document,omitempty"`
	OrderConfirmationItems  null.JSON   `boil:"order_confirmation_items" json:"order_confirmation_items,omitempty" toml:"order_confirmation_items" yaml:"order_confirmation_items,omitempty"`
}

var OrderConfirmationViewColumns = struct {
	ID                      string
	BaseDocumentID          string
	OrderConfirmationNumber string
	SalesOrderID            string
	CustomerID              string
	ShipToInformation       string
	BaseDocument            string
	OrderConfirmationItems  string
}{
	ID:                      "id",
	BaseDocumentID:          "base_document_id",
	OrderConfirmationNumber: "order_confirmation_number",
	SalesOrderID:            "sales_order_id",
	CustomerID:              "customer_id",
	ShipToInformation:       "ship_to_information",
	BaseDocument:            "base_document",
	OrderConfirmationItems:  "order_confirmation_items",
}

var OrderConfirmationViewTableColumns = struct {
	ID                      string
	BaseDocumentID          string
	OrderConfirmationNumber string
	SalesOrderID            string
	CustomerID              string
	ShipToInformation       string
	BaseDocument            string
	OrderConfirmationItems  string
}{
	ID:                      "order_confirmation_view.id",
	BaseDocumentID:          "order_confirmation_view.base_document_id",
	OrderConfirmationNumber: "order_confirmation_view.order_confirmation_number",
	SalesOrderID:            "order_confirmation_view.sales_order_id",
	CustomerID:              "order_confirmation_view.customer_id",
	ShipToInformation:       "order_confirmation_view.ship_to_information",
	BaseDocument:            "order_confirmation_view.base_document",
	OrderConfirmationItems:  "order_confirmation_view.order_confirmation_items",
}

// Generated where

var OrderConfirmationViewWhere = struct {
	ID                      whereHelpernull_Int
	BaseDocumentID          whereHelpernull_Int
	OrderConfirmationNumber whereHelpernull_String
	SalesOrderID            whereHelpernull_Int
	CustomerID              whereHelpernull_Int
	ShipToInformation       whereHelpernull_String
	BaseDocument            whereHelpernull_JSON
	OrderConfirmationItems  whereHelpernull_JSON
}{
	ID:                      whereHelpernull_Int{field: "\"sale\".\"order_confirmation_view\".\"id\""},
	BaseDocumentID:          whereHelpernull_Int{field: "\"sale\".\"order_confirmation_view\".\"base_document_id\""},
	OrderConfirmationNumber: whereHelpernull_String{field: "\"sale\".\"order_confirmation_view\".\"order_confirmation_number\""},
	SalesOrderID:            whereHelpernull_Int{field: "\"sale\".\"order_confirmation_view\".\"sales_order_id\""},
	CustomerID:              whereHelpernull_Int{field: "\"sale\".\"order_confirmation_view\".\"customer_id\""},
	ShipToInformation:       whereHelpernull_String{field: "\"sale\".\"order_confirmation_view\".\"ship_to_information\""},
	BaseDocument:            whereHelpernull_JSON{field: "\"sale\".\"order_confirmation_view\".\"base_document\""},
	OrderConfirmationItems:  whereHelpernull_JSON{field: "\"sale\".\"order_confirmation_view\".\"order_confirmation_items\""},
}

var (
	orderConfirmationViewAllColumns            = []string{"id", "base_document_id", "order_confirmation_number", "sales_order_id", "customer_id", "ship_to_information", "base_document", "order_confirmation_items"}
	orderConfirmationViewColumnsWithoutDefault = []string{}
	orderConfirmationViewColumnsWithDefault    = []string{"id", "base_document_id", "order_confirmation_number", "sales_order_id", "customer_id", "ship_to_information", "base_document", "order_confirmation_items"}
	orderConfirmationViewPrimaryKeyColumns     = []string{}
	orderConfirmationViewGeneratedColumns      = []string{}
)

type (
	// OrderConfirmationViewSlice is an alias for a slice of pointers to OrderConfirmationView.
	// This should almost always be used instead of []OrderConfirmationView.
	OrderConfirmationViewSlice []*OrderConfirmationView
	// OrderConfirmationViewHook is the signature for custom OrderConfirmationView hook methods
	OrderConfirmationViewHook func(context.Context, boil.ContextExecutor, *OrderConfirmationView) error

	orderConfirmationViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderConfirmationViewType           = reflect.TypeOf(&OrderConfirmationView{})
	orderConfirmationViewMapping        = queries.MakeStructMapping(orderConfirmationViewType)
	orderConfirmationViewInsertCacheMut sync.RWMutex
	orderConfirmationViewInsertCache    = make(map[string]insertCache)
	orderConfirmationViewUpdateCacheMut sync.RWMutex
	orderConfirmationViewUpdateCache    = make(map[string]updateCache)
	orderConfirmationViewUpsertCacheMut sync.RWMutex
	orderConfirmationViewUpsertCache    = make(map[string]insertCache)
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

var orderConfirmationViewAfterSelectMu sync.Mutex
var orderConfirmationViewAfterSelectHooks []OrderConfirmationViewHook

var orderConfirmationViewBeforeInsertMu sync.Mutex
var orderConfirmationViewBeforeInsertHooks []OrderConfirmationViewHook
var orderConfirmationViewAfterInsertMu sync.Mutex
var orderConfirmationViewAfterInsertHooks []OrderConfirmationViewHook

var orderConfirmationViewBeforeUpsertMu sync.Mutex
var orderConfirmationViewBeforeUpsertHooks []OrderConfirmationViewHook
var orderConfirmationViewAfterUpsertMu sync.Mutex
var orderConfirmationViewAfterUpsertHooks []OrderConfirmationViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrderConfirmationView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderConfirmationViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrderConfirmationView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderConfirmationViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrderConfirmationView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderConfirmationViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrderConfirmationView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderConfirmationViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrderConfirmationView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orderConfirmationViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrderConfirmationViewHook registers your hook function for all future operations.
func AddOrderConfirmationViewHook(hookPoint boil.HookPoint, orderConfirmationViewHook OrderConfirmationViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orderConfirmationViewAfterSelectMu.Lock()
		orderConfirmationViewAfterSelectHooks = append(orderConfirmationViewAfterSelectHooks, orderConfirmationViewHook)
		orderConfirmationViewAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		orderConfirmationViewBeforeInsertMu.Lock()
		orderConfirmationViewBeforeInsertHooks = append(orderConfirmationViewBeforeInsertHooks, orderConfirmationViewHook)
		orderConfirmationViewBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		orderConfirmationViewAfterInsertMu.Lock()
		orderConfirmationViewAfterInsertHooks = append(orderConfirmationViewAfterInsertHooks, orderConfirmationViewHook)
		orderConfirmationViewAfterInsertMu.Unlock()
	case boil.BeforeUpsertHook:
		orderConfirmationViewBeforeUpsertMu.Lock()
		orderConfirmationViewBeforeUpsertHooks = append(orderConfirmationViewBeforeUpsertHooks, orderConfirmationViewHook)
		orderConfirmationViewBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		orderConfirmationViewAfterUpsertMu.Lock()
		orderConfirmationViewAfterUpsertHooks = append(orderConfirmationViewAfterUpsertHooks, orderConfirmationViewHook)
		orderConfirmationViewAfterUpsertMu.Unlock()
	}
}

// One returns a single orderConfirmationView record from the query.
func (q orderConfirmationViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderConfirmationView, error) {
	o := &OrderConfirmationView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: failed to execute a one query for order_confirmation_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrderConfirmationView records from the query.
func (q orderConfirmationViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderConfirmationViewSlice, error) {
	var o []*OrderConfirmationView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sale: failed to assign all query results to OrderConfirmationView slice")
	}

	if len(orderConfirmationViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrderConfirmationView records in the query.
func (q orderConfirmationViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to count order_confirmation_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderConfirmationViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sale: failed to check if order_confirmation_view exists")
	}

	return count > 0, nil
}

// OrderConfirmationViews retrieves all the records using an executor.
func OrderConfirmationViews(mods ...qm.QueryMod) orderConfirmationViewQuery {
	mods = append(mods, qm.From("\"sale\".\"order_confirmation_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sale\".\"order_confirmation_view\".*"})
	}

	return orderConfirmationViewQuery{q}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderConfirmationView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sale: no order_confirmation_view provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderConfirmationViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderConfirmationViewInsertCacheMut.RLock()
	cache, cached := orderConfirmationViewInsertCache[key]
	orderConfirmationViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderConfirmationViewAllColumns,
			orderConfirmationViewColumnsWithDefault,
			orderConfirmationViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderConfirmationViewType, orderConfirmationViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderConfirmationViewType, orderConfirmationViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sale\".\"order_confirmation_view\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sale\".\"order_confirmation_view\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sale: unable to insert into order_confirmation_view")
	}

	if !cached {
		orderConfirmationViewInsertCacheMut.Lock()
		orderConfirmationViewInsertCache[key] = cache
		orderConfirmationViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderConfirmationView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("sale: no order_confirmation_view provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orderConfirmationViewColumnsWithDefault, o)

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

	orderConfirmationViewUpsertCacheMut.RLock()
	cache, cached := orderConfirmationViewUpsertCache[key]
	orderConfirmationViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			orderConfirmationViewAllColumns,
			orderConfirmationViewColumnsWithDefault,
			orderConfirmationViewColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orderConfirmationViewAllColumns,
			orderConfirmationViewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sale: unable to upsert order_confirmation_view, could not build update column list")
		}

		ret := strmangle.SetComplement(orderConfirmationViewAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(orderConfirmationViewPrimaryKeyColumns) == 0 {
				return errors.New("sale: unable to upsert order_confirmation_view, could not build conflict column list")
			}

			conflict = make([]string, len(orderConfirmationViewPrimaryKeyColumns))
			copy(conflict, orderConfirmationViewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"sale\".\"order_confirmation_view\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(orderConfirmationViewType, orderConfirmationViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderConfirmationViewType, orderConfirmationViewMapping, ret)
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
		return errors.Wrap(err, "sale: unable to upsert order_confirmation_view")
	}

	if !cached {
		orderConfirmationViewUpsertCacheMut.Lock()
		orderConfirmationViewUpsertCache[key] = cache
		orderConfirmationViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
