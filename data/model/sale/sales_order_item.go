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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SalesOrderItem is an object representing the database table.
type SalesOrderItem struct {
	ID                 int `boil:"id" json:"id" toml:"id" yaml:"id"`
	BaseDocumentItemID int `boil:"base_document_item_id" json:"base_document_item_id" toml:"base_document_item_id" yaml:"base_document_item_id"`
	SalesOrderID       int `boil:"sales_order_id" json:"sales_order_id" toml:"sales_order_id" yaml:"sales_order_id"`

	R *salesOrderItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L salesOrderItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SalesOrderItemColumns = struct {
	ID                 string
	BaseDocumentItemID string
	SalesOrderID       string
}{
	ID:                 "id",
	BaseDocumentItemID: "base_document_item_id",
	SalesOrderID:       "sales_order_id",
}

var SalesOrderItemTableColumns = struct {
	ID                 string
	BaseDocumentItemID string
	SalesOrderID       string
}{
	ID:                 "sales_order_item.id",
	BaseDocumentItemID: "sales_order_item.base_document_item_id",
	SalesOrderID:       "sales_order_item.sales_order_id",
}

// Generated where

var SalesOrderItemWhere = struct {
	ID                 whereHelperint
	BaseDocumentItemID whereHelperint
	SalesOrderID       whereHelperint
}{
	ID:                 whereHelperint{field: "\"sale\".\"sales_order_item\".\"id\""},
	BaseDocumentItemID: whereHelperint{field: "\"sale\".\"sales_order_item\".\"base_document_item_id\""},
	SalesOrderID:       whereHelperint{field: "\"sale\".\"sales_order_item\".\"sales_order_id\""},
}

// SalesOrderItemRels is where relationship names are stored.
var SalesOrderItemRels = struct {
	SalesOrder string
}{
	SalesOrder: "SalesOrder",
}

// salesOrderItemR is where relationships are stored.
type salesOrderItemR struct {
	SalesOrder *SalesOrder `boil:"SalesOrder" json:"SalesOrder" toml:"SalesOrder" yaml:"SalesOrder"`
}

// NewStruct creates a new relationship struct
func (*salesOrderItemR) NewStruct() *salesOrderItemR {
	return &salesOrderItemR{}
}

func (r *salesOrderItemR) GetSalesOrder() *SalesOrder {
	if r == nil {
		return nil
	}
	return r.SalesOrder
}

// salesOrderItemL is where Load methods for each relationship are stored.
type salesOrderItemL struct{}

var (
	salesOrderItemAllColumns            = []string{"id", "base_document_item_id", "sales_order_id"}
	salesOrderItemColumnsWithoutDefault = []string{"id", "base_document_item_id", "sales_order_id"}
	salesOrderItemColumnsWithDefault    = []string{}
	salesOrderItemPrimaryKeyColumns     = []string{"id"}
	salesOrderItemGeneratedColumns      = []string{}
)

type (
	// SalesOrderItemSlice is an alias for a slice of pointers to SalesOrderItem.
	// This should almost always be used instead of []SalesOrderItem.
	SalesOrderItemSlice []*SalesOrderItem
	// SalesOrderItemHook is the signature for custom SalesOrderItem hook methods
	SalesOrderItemHook func(context.Context, boil.ContextExecutor, *SalesOrderItem) error

	salesOrderItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	salesOrderItemType                 = reflect.TypeOf(&SalesOrderItem{})
	salesOrderItemMapping              = queries.MakeStructMapping(salesOrderItemType)
	salesOrderItemPrimaryKeyMapping, _ = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, salesOrderItemPrimaryKeyColumns)
	salesOrderItemInsertCacheMut       sync.RWMutex
	salesOrderItemInsertCache          = make(map[string]insertCache)
	salesOrderItemUpdateCacheMut       sync.RWMutex
	salesOrderItemUpdateCache          = make(map[string]updateCache)
	salesOrderItemUpsertCacheMut       sync.RWMutex
	salesOrderItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var salesOrderItemAfterSelectMu sync.Mutex
var salesOrderItemAfterSelectHooks []SalesOrderItemHook

var salesOrderItemBeforeInsertMu sync.Mutex
var salesOrderItemBeforeInsertHooks []SalesOrderItemHook
var salesOrderItemAfterInsertMu sync.Mutex
var salesOrderItemAfterInsertHooks []SalesOrderItemHook

var salesOrderItemBeforeUpdateMu sync.Mutex
var salesOrderItemBeforeUpdateHooks []SalesOrderItemHook
var salesOrderItemAfterUpdateMu sync.Mutex
var salesOrderItemAfterUpdateHooks []SalesOrderItemHook

var salesOrderItemBeforeDeleteMu sync.Mutex
var salesOrderItemBeforeDeleteHooks []SalesOrderItemHook
var salesOrderItemAfterDeleteMu sync.Mutex
var salesOrderItemAfterDeleteHooks []SalesOrderItemHook

var salesOrderItemBeforeUpsertMu sync.Mutex
var salesOrderItemBeforeUpsertHooks []SalesOrderItemHook
var salesOrderItemAfterUpsertMu sync.Mutex
var salesOrderItemAfterUpsertHooks []SalesOrderItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SalesOrderItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SalesOrderItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SalesOrderItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SalesOrderItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SalesOrderItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SalesOrderItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SalesOrderItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SalesOrderItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SalesOrderItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesOrderItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSalesOrderItemHook registers your hook function for all future operations.
func AddSalesOrderItemHook(hookPoint boil.HookPoint, salesOrderItemHook SalesOrderItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		salesOrderItemAfterSelectMu.Lock()
		salesOrderItemAfterSelectHooks = append(salesOrderItemAfterSelectHooks, salesOrderItemHook)
		salesOrderItemAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		salesOrderItemBeforeInsertMu.Lock()
		salesOrderItemBeforeInsertHooks = append(salesOrderItemBeforeInsertHooks, salesOrderItemHook)
		salesOrderItemBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		salesOrderItemAfterInsertMu.Lock()
		salesOrderItemAfterInsertHooks = append(salesOrderItemAfterInsertHooks, salesOrderItemHook)
		salesOrderItemAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		salesOrderItemBeforeUpdateMu.Lock()
		salesOrderItemBeforeUpdateHooks = append(salesOrderItemBeforeUpdateHooks, salesOrderItemHook)
		salesOrderItemBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		salesOrderItemAfterUpdateMu.Lock()
		salesOrderItemAfterUpdateHooks = append(salesOrderItemAfterUpdateHooks, salesOrderItemHook)
		salesOrderItemAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		salesOrderItemBeforeDeleteMu.Lock()
		salesOrderItemBeforeDeleteHooks = append(salesOrderItemBeforeDeleteHooks, salesOrderItemHook)
		salesOrderItemBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		salesOrderItemAfterDeleteMu.Lock()
		salesOrderItemAfterDeleteHooks = append(salesOrderItemAfterDeleteHooks, salesOrderItemHook)
		salesOrderItemAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		salesOrderItemBeforeUpsertMu.Lock()
		salesOrderItemBeforeUpsertHooks = append(salesOrderItemBeforeUpsertHooks, salesOrderItemHook)
		salesOrderItemBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		salesOrderItemAfterUpsertMu.Lock()
		salesOrderItemAfterUpsertHooks = append(salesOrderItemAfterUpsertHooks, salesOrderItemHook)
		salesOrderItemAfterUpsertMu.Unlock()
	}
}

// One returns a single salesOrderItem record from the query.
func (q salesOrderItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SalesOrderItem, error) {
	o := &SalesOrderItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: failed to execute a one query for sales_order_item")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SalesOrderItem records from the query.
func (q salesOrderItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (SalesOrderItemSlice, error) {
	var o []*SalesOrderItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sale: failed to assign all query results to SalesOrderItem slice")
	}

	if len(salesOrderItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SalesOrderItem records in the query.
func (q salesOrderItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to count sales_order_item rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q salesOrderItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sale: failed to check if sales_order_item exists")
	}

	return count > 0, nil
}

// SalesOrder pointed to by the foreign key.
func (o *SalesOrderItem) SalesOrder(mods ...qm.QueryMod) salesOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SalesOrderID),
	}

	queryMods = append(queryMods, mods...)

	return SalesOrders(queryMods...)
}

// LoadSalesOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (salesOrderItemL) LoadSalesOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSalesOrderItem interface{}, mods queries.Applicator) error {
	var slice []*SalesOrderItem
	var object *SalesOrderItem

	if singular {
		var ok bool
		object, ok = maybeSalesOrderItem.(*SalesOrderItem)
		if !ok {
			object = new(SalesOrderItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSalesOrderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSalesOrderItem))
			}
		}
	} else {
		s, ok := maybeSalesOrderItem.(*[]*SalesOrderItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSalesOrderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSalesOrderItem))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &salesOrderItemR{}
		}
		args[object.SalesOrderID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &salesOrderItemR{}
			}

			args[obj.SalesOrderID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`sale.sales_order`),
		qm.WhereIn(`sale.sales_order.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SalesOrder")
	}

	var resultSlice []*SalesOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SalesOrder")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for sales_order")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sales_order")
	}

	if len(salesOrderAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.SalesOrder = foreign
		if foreign.R == nil {
			foreign.R = &salesOrderR{}
		}
		foreign.R.SalesOrderItems = append(foreign.R.SalesOrderItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SalesOrderID == foreign.ID {
				local.R.SalesOrder = foreign
				if foreign.R == nil {
					foreign.R = &salesOrderR{}
				}
				foreign.R.SalesOrderItems = append(foreign.R.SalesOrderItems, local)
				break
			}
		}
	}

	return nil
}

// SetSalesOrder of the salesOrderItem to the related item.
// Sets o.R.SalesOrder to related.
// Adds o to related.R.SalesOrderItems.
func (o *SalesOrderItem) SetSalesOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SalesOrder) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"sale\".\"sales_order_item\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"sales_order_id"}),
		strmangle.WhereClause("\"", "\"", 2, salesOrderItemPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SalesOrderID = related.ID
	if o.R == nil {
		o.R = &salesOrderItemR{
			SalesOrder: related,
		}
	} else {
		o.R.SalesOrder = related
	}

	if related.R == nil {
		related.R = &salesOrderR{
			SalesOrderItems: SalesOrderItemSlice{o},
		}
	} else {
		related.R.SalesOrderItems = append(related.R.SalesOrderItems, o)
	}

	return nil
}

// SalesOrderItems retrieves all the records using an executor.
func SalesOrderItems(mods ...qm.QueryMod) salesOrderItemQuery {
	mods = append(mods, qm.From("\"sale\".\"sales_order_item\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sale\".\"sales_order_item\".*"})
	}

	return salesOrderItemQuery{q}
}

// FindSalesOrderItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSalesOrderItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*SalesOrderItem, error) {
	salesOrderItemObj := &SalesOrderItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"sale\".\"sales_order_item\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, salesOrderItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: unable to select from sales_order_item")
	}

	if err = salesOrderItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return salesOrderItemObj, err
	}

	return salesOrderItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SalesOrderItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sale: no sales_order_item provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(salesOrderItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	salesOrderItemInsertCacheMut.RLock()
	cache, cached := salesOrderItemInsertCache[key]
	salesOrderItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			salesOrderItemAllColumns,
			salesOrderItemColumnsWithDefault,
			salesOrderItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sale\".\"sales_order_item\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sale\".\"sales_order_item\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sale: unable to insert into sales_order_item")
	}

	if !cached {
		salesOrderItemInsertCacheMut.Lock()
		salesOrderItemInsertCache[key] = cache
		salesOrderItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SalesOrderItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SalesOrderItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	salesOrderItemUpdateCacheMut.RLock()
	cache, cached := salesOrderItemUpdateCache[key]
	salesOrderItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			salesOrderItemAllColumns,
			salesOrderItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sale: unable to update sales_order_item, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"sale\".\"sales_order_item\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, salesOrderItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, append(wl, salesOrderItemPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to update sales_order_item row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by update for sales_order_item")
	}

	if !cached {
		salesOrderItemUpdateCacheMut.Lock()
		salesOrderItemUpdateCache[key] = cache
		salesOrderItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q salesOrderItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to update all for sales_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to retrieve rows affected for sales_order_item")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SalesOrderItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sale: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"sale\".\"sales_order_item\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, salesOrderItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to update all in salesOrderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to retrieve rows affected all in update all salesOrderItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SalesOrderItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("sale: no sales_order_item provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(salesOrderItemColumnsWithDefault, o)

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

	salesOrderItemUpsertCacheMut.RLock()
	cache, cached := salesOrderItemUpsertCache[key]
	salesOrderItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			salesOrderItemAllColumns,
			salesOrderItemColumnsWithDefault,
			salesOrderItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			salesOrderItemAllColumns,
			salesOrderItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sale: unable to upsert sales_order_item, could not build update column list")
		}

		ret := strmangle.SetComplement(salesOrderItemAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(salesOrderItemPrimaryKeyColumns) == 0 {
				return errors.New("sale: unable to upsert sales_order_item, could not build conflict column list")
			}

			conflict = make([]string, len(salesOrderItemPrimaryKeyColumns))
			copy(conflict, salesOrderItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"sale\".\"sales_order_item\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(salesOrderItemType, salesOrderItemMapping, ret)
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
		return errors.Wrap(err, "sale: unable to upsert sales_order_item")
	}

	if !cached {
		salesOrderItemUpsertCacheMut.Lock()
		salesOrderItemUpsertCache[key] = cache
		salesOrderItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SalesOrderItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SalesOrderItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sale: no SalesOrderItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), salesOrderItemPrimaryKeyMapping)
	sql := "DELETE FROM \"sale\".\"sales_order_item\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete from sales_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by delete for sales_order_item")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q salesOrderItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sale: no salesOrderItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete all from sales_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by deleteall for sales_order_item")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SalesOrderItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(salesOrderItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"sale\".\"sales_order_item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesOrderItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete all from salesOrderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by deleteall for sales_order_item")
	}

	if len(salesOrderItemAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SalesOrderItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSalesOrderItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SalesOrderItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SalesOrderItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"sale\".\"sales_order_item\".* FROM \"sale\".\"sales_order_item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesOrderItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sale: unable to reload all in SalesOrderItemSlice")
	}

	*o = slice

	return nil
}

// SalesOrderItemExists checks if the SalesOrderItem row exists.
func SalesOrderItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"sale\".\"sales_order_item\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sale: unable to check if sales_order_item exists")
	}

	return exists, nil
}

// Exists checks if the SalesOrderItem row exists.
func (o *SalesOrderItem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SalesOrderItemExists(ctx, exec, o.ID)
}
