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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PurchaseOrderItem is an object representing the database table.
type PurchaseOrderItem struct {
	ID                 int `boil:"id" json:"id" toml:"id" yaml:"id"`
	BaseDocumentItemID int `boil:"base_document_item_id" json:"base_document_item_id" toml:"base_document_item_id" yaml:"base_document_item_id"`
	PurchaseOrderID    int `boil:"purchase_order_id" json:"purchase_order_id" toml:"purchase_order_id" yaml:"purchase_order_id"`

	R *purchaseOrderItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L purchaseOrderItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PurchaseOrderItemColumns = struct {
	ID                 string
	BaseDocumentItemID string
	PurchaseOrderID    string
}{
	ID:                 "id",
	BaseDocumentItemID: "base_document_item_id",
	PurchaseOrderID:    "purchase_order_id",
}

var PurchaseOrderItemTableColumns = struct {
	ID                 string
	BaseDocumentItemID string
	PurchaseOrderID    string
}{
	ID:                 "purchase_order_item.id",
	BaseDocumentItemID: "purchase_order_item.base_document_item_id",
	PurchaseOrderID:    "purchase_order_item.purchase_order_id",
}

// Generated where

var PurchaseOrderItemWhere = struct {
	ID                 whereHelperint
	BaseDocumentItemID whereHelperint
	PurchaseOrderID    whereHelperint
}{
	ID:                 whereHelperint{field: "\"purchase\".\"purchase_order_item\".\"id\""},
	BaseDocumentItemID: whereHelperint{field: "\"purchase\".\"purchase_order_item\".\"base_document_item_id\""},
	PurchaseOrderID:    whereHelperint{field: "\"purchase\".\"purchase_order_item\".\"purchase_order_id\""},
}

// PurchaseOrderItemRels is where relationship names are stored.
var PurchaseOrderItemRels = struct {
	PurchaseOrder string
}{
	PurchaseOrder: "PurchaseOrder",
}

// purchaseOrderItemR is where relationships are stored.
type purchaseOrderItemR struct {
	PurchaseOrder *PurchaseOrder `boil:"PurchaseOrder" json:"PurchaseOrder" toml:"PurchaseOrder" yaml:"PurchaseOrder"`
}

// NewStruct creates a new relationship struct
func (*purchaseOrderItemR) NewStruct() *purchaseOrderItemR {
	return &purchaseOrderItemR{}
}

func (r *purchaseOrderItemR) GetPurchaseOrder() *PurchaseOrder {
	if r == nil {
		return nil
	}
	return r.PurchaseOrder
}

// purchaseOrderItemL is where Load methods for each relationship are stored.
type purchaseOrderItemL struct{}

var (
	purchaseOrderItemAllColumns            = []string{"id", "base_document_item_id", "purchase_order_id"}
	purchaseOrderItemColumnsWithoutDefault = []string{"id", "base_document_item_id", "purchase_order_id"}
	purchaseOrderItemColumnsWithDefault    = []string{}
	purchaseOrderItemPrimaryKeyColumns     = []string{"id"}
	purchaseOrderItemGeneratedColumns      = []string{}
)

type (
	// PurchaseOrderItemSlice is an alias for a slice of pointers to PurchaseOrderItem.
	// This should almost always be used instead of []PurchaseOrderItem.
	PurchaseOrderItemSlice []*PurchaseOrderItem
	// PurchaseOrderItemHook is the signature for custom PurchaseOrderItem hook methods
	PurchaseOrderItemHook func(context.Context, boil.ContextExecutor, *PurchaseOrderItem) error

	purchaseOrderItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	purchaseOrderItemType                 = reflect.TypeOf(&PurchaseOrderItem{})
	purchaseOrderItemMapping              = queries.MakeStructMapping(purchaseOrderItemType)
	purchaseOrderItemPrimaryKeyMapping, _ = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, purchaseOrderItemPrimaryKeyColumns)
	purchaseOrderItemInsertCacheMut       sync.RWMutex
	purchaseOrderItemInsertCache          = make(map[string]insertCache)
	purchaseOrderItemUpdateCacheMut       sync.RWMutex
	purchaseOrderItemUpdateCache          = make(map[string]updateCache)
	purchaseOrderItemUpsertCacheMut       sync.RWMutex
	purchaseOrderItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var purchaseOrderItemAfterSelectMu sync.Mutex
var purchaseOrderItemAfterSelectHooks []PurchaseOrderItemHook

var purchaseOrderItemBeforeInsertMu sync.Mutex
var purchaseOrderItemBeforeInsertHooks []PurchaseOrderItemHook
var purchaseOrderItemAfterInsertMu sync.Mutex
var purchaseOrderItemAfterInsertHooks []PurchaseOrderItemHook

var purchaseOrderItemBeforeUpdateMu sync.Mutex
var purchaseOrderItemBeforeUpdateHooks []PurchaseOrderItemHook
var purchaseOrderItemAfterUpdateMu sync.Mutex
var purchaseOrderItemAfterUpdateHooks []PurchaseOrderItemHook

var purchaseOrderItemBeforeDeleteMu sync.Mutex
var purchaseOrderItemBeforeDeleteHooks []PurchaseOrderItemHook
var purchaseOrderItemAfterDeleteMu sync.Mutex
var purchaseOrderItemAfterDeleteHooks []PurchaseOrderItemHook

var purchaseOrderItemBeforeUpsertMu sync.Mutex
var purchaseOrderItemBeforeUpsertHooks []PurchaseOrderItemHook
var purchaseOrderItemAfterUpsertMu sync.Mutex
var purchaseOrderItemAfterUpsertHooks []PurchaseOrderItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PurchaseOrderItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PurchaseOrderItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PurchaseOrderItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PurchaseOrderItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PurchaseOrderItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PurchaseOrderItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PurchaseOrderItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PurchaseOrderItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PurchaseOrderItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range purchaseOrderItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPurchaseOrderItemHook registers your hook function for all future operations.
func AddPurchaseOrderItemHook(hookPoint boil.HookPoint, purchaseOrderItemHook PurchaseOrderItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		purchaseOrderItemAfterSelectMu.Lock()
		purchaseOrderItemAfterSelectHooks = append(purchaseOrderItemAfterSelectHooks, purchaseOrderItemHook)
		purchaseOrderItemAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		purchaseOrderItemBeforeInsertMu.Lock()
		purchaseOrderItemBeforeInsertHooks = append(purchaseOrderItemBeforeInsertHooks, purchaseOrderItemHook)
		purchaseOrderItemBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		purchaseOrderItemAfterInsertMu.Lock()
		purchaseOrderItemAfterInsertHooks = append(purchaseOrderItemAfterInsertHooks, purchaseOrderItemHook)
		purchaseOrderItemAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		purchaseOrderItemBeforeUpdateMu.Lock()
		purchaseOrderItemBeforeUpdateHooks = append(purchaseOrderItemBeforeUpdateHooks, purchaseOrderItemHook)
		purchaseOrderItemBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		purchaseOrderItemAfterUpdateMu.Lock()
		purchaseOrderItemAfterUpdateHooks = append(purchaseOrderItemAfterUpdateHooks, purchaseOrderItemHook)
		purchaseOrderItemAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		purchaseOrderItemBeforeDeleteMu.Lock()
		purchaseOrderItemBeforeDeleteHooks = append(purchaseOrderItemBeforeDeleteHooks, purchaseOrderItemHook)
		purchaseOrderItemBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		purchaseOrderItemAfterDeleteMu.Lock()
		purchaseOrderItemAfterDeleteHooks = append(purchaseOrderItemAfterDeleteHooks, purchaseOrderItemHook)
		purchaseOrderItemAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		purchaseOrderItemBeforeUpsertMu.Lock()
		purchaseOrderItemBeforeUpsertHooks = append(purchaseOrderItemBeforeUpsertHooks, purchaseOrderItemHook)
		purchaseOrderItemBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		purchaseOrderItemAfterUpsertMu.Lock()
		purchaseOrderItemAfterUpsertHooks = append(purchaseOrderItemAfterUpsertHooks, purchaseOrderItemHook)
		purchaseOrderItemAfterUpsertMu.Unlock()
	}
}

// One returns a single purchaseOrderItem record from the query.
func (q purchaseOrderItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PurchaseOrderItem, error) {
	o := &PurchaseOrderItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "purchase: failed to execute a one query for purchase_order_item")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PurchaseOrderItem records from the query.
func (q purchaseOrderItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (PurchaseOrderItemSlice, error) {
	var o []*PurchaseOrderItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "purchase: failed to assign all query results to PurchaseOrderItem slice")
	}

	if len(purchaseOrderItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PurchaseOrderItem records in the query.
func (q purchaseOrderItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to count purchase_order_item rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q purchaseOrderItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "purchase: failed to check if purchase_order_item exists")
	}

	return count > 0, nil
}

// PurchaseOrder pointed to by the foreign key.
func (o *PurchaseOrderItem) PurchaseOrder(mods ...qm.QueryMod) purchaseOrderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PurchaseOrderID),
	}

	queryMods = append(queryMods, mods...)

	return PurchaseOrders(queryMods...)
}

// LoadPurchaseOrder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (purchaseOrderItemL) LoadPurchaseOrder(ctx context.Context, e boil.ContextExecutor, singular bool, maybePurchaseOrderItem interface{}, mods queries.Applicator) error {
	var slice []*PurchaseOrderItem
	var object *PurchaseOrderItem

	if singular {
		var ok bool
		object, ok = maybePurchaseOrderItem.(*PurchaseOrderItem)
		if !ok {
			object = new(PurchaseOrderItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePurchaseOrderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePurchaseOrderItem))
			}
		}
	} else {
		s, ok := maybePurchaseOrderItem.(*[]*PurchaseOrderItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePurchaseOrderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePurchaseOrderItem))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &purchaseOrderItemR{}
		}
		args[object.PurchaseOrderID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &purchaseOrderItemR{}
			}

			args[obj.PurchaseOrderID] = struct{}{}

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
		qm.From(`purchase.purchase_order`),
		qm.WhereIn(`purchase.purchase_order.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PurchaseOrder")
	}

	var resultSlice []*PurchaseOrder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PurchaseOrder")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for purchase_order")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for purchase_order")
	}

	if len(purchaseOrderAfterSelectHooks) != 0 {
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
		object.R.PurchaseOrder = foreign
		if foreign.R == nil {
			foreign.R = &purchaseOrderR{}
		}
		foreign.R.PurchaseOrderItems = append(foreign.R.PurchaseOrderItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PurchaseOrderID == foreign.ID {
				local.R.PurchaseOrder = foreign
				if foreign.R == nil {
					foreign.R = &purchaseOrderR{}
				}
				foreign.R.PurchaseOrderItems = append(foreign.R.PurchaseOrderItems, local)
				break
			}
		}
	}

	return nil
}

// SetPurchaseOrder of the purchaseOrderItem to the related item.
// Sets o.R.PurchaseOrder to related.
// Adds o to related.R.PurchaseOrderItems.
func (o *PurchaseOrderItem) SetPurchaseOrder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PurchaseOrder) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"purchase\".\"purchase_order_item\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"purchase_order_id"}),
		strmangle.WhereClause("\"", "\"", 2, purchaseOrderItemPrimaryKeyColumns),
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

	o.PurchaseOrderID = related.ID
	if o.R == nil {
		o.R = &purchaseOrderItemR{
			PurchaseOrder: related,
		}
	} else {
		o.R.PurchaseOrder = related
	}

	if related.R == nil {
		related.R = &purchaseOrderR{
			PurchaseOrderItems: PurchaseOrderItemSlice{o},
		}
	} else {
		related.R.PurchaseOrderItems = append(related.R.PurchaseOrderItems, o)
	}

	return nil
}

// PurchaseOrderItems retrieves all the records using an executor.
func PurchaseOrderItems(mods ...qm.QueryMod) purchaseOrderItemQuery {
	mods = append(mods, qm.From("\"purchase\".\"purchase_order_item\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"purchase\".\"purchase_order_item\".*"})
	}

	return purchaseOrderItemQuery{q}
}

// FindPurchaseOrderItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPurchaseOrderItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*PurchaseOrderItem, error) {
	purchaseOrderItemObj := &PurchaseOrderItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"purchase\".\"purchase_order_item\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, purchaseOrderItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "purchase: unable to select from purchase_order_item")
	}

	if err = purchaseOrderItemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return purchaseOrderItemObj, err
	}

	return purchaseOrderItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PurchaseOrderItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("purchase: no purchase_order_item provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(purchaseOrderItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	purchaseOrderItemInsertCacheMut.RLock()
	cache, cached := purchaseOrderItemInsertCache[key]
	purchaseOrderItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			purchaseOrderItemAllColumns,
			purchaseOrderItemColumnsWithDefault,
			purchaseOrderItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"purchase\".\"purchase_order_item\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"purchase\".\"purchase_order_item\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "purchase: unable to insert into purchase_order_item")
	}

	if !cached {
		purchaseOrderItemInsertCacheMut.Lock()
		purchaseOrderItemInsertCache[key] = cache
		purchaseOrderItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PurchaseOrderItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PurchaseOrderItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	purchaseOrderItemUpdateCacheMut.RLock()
	cache, cached := purchaseOrderItemUpdateCache[key]
	purchaseOrderItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			purchaseOrderItemAllColumns,
			purchaseOrderItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("purchase: unable to update purchase_order_item, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"purchase\".\"purchase_order_item\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, purchaseOrderItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, append(wl, purchaseOrderItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "purchase: unable to update purchase_order_item row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to get rows affected by update for purchase_order_item")
	}

	if !cached {
		purchaseOrderItemUpdateCacheMut.Lock()
		purchaseOrderItemUpdateCache[key] = cache
		purchaseOrderItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q purchaseOrderItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to update all for purchase_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to retrieve rows affected for purchase_order_item")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PurchaseOrderItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("purchase: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchaseOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"purchase\".\"purchase_order_item\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, purchaseOrderItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to update all in purchaseOrderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to retrieve rows affected all in update all purchaseOrderItem")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PurchaseOrderItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("purchase: no purchase_order_item provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(purchaseOrderItemColumnsWithDefault, o)

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

	purchaseOrderItemUpsertCacheMut.RLock()
	cache, cached := purchaseOrderItemUpsertCache[key]
	purchaseOrderItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			purchaseOrderItemAllColumns,
			purchaseOrderItemColumnsWithDefault,
			purchaseOrderItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			purchaseOrderItemAllColumns,
			purchaseOrderItemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("purchase: unable to upsert purchase_order_item, could not build update column list")
		}

		ret := strmangle.SetComplement(purchaseOrderItemAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(purchaseOrderItemPrimaryKeyColumns) == 0 {
				return errors.New("purchase: unable to upsert purchase_order_item, could not build conflict column list")
			}

			conflict = make([]string, len(purchaseOrderItemPrimaryKeyColumns))
			copy(conflict, purchaseOrderItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"purchase\".\"purchase_order_item\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(purchaseOrderItemType, purchaseOrderItemMapping, ret)
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
		return errors.Wrap(err, "purchase: unable to upsert purchase_order_item")
	}

	if !cached {
		purchaseOrderItemUpsertCacheMut.Lock()
		purchaseOrderItemUpsertCache[key] = cache
		purchaseOrderItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PurchaseOrderItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PurchaseOrderItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("purchase: no PurchaseOrderItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), purchaseOrderItemPrimaryKeyMapping)
	sql := "DELETE FROM \"purchase\".\"purchase_order_item\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to delete from purchase_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to get rows affected by delete for purchase_order_item")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q purchaseOrderItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("purchase: no purchaseOrderItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to delete all from purchase_order_item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to get rows affected by deleteall for purchase_order_item")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PurchaseOrderItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(purchaseOrderItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchaseOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"purchase\".\"purchase_order_item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, purchaseOrderItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "purchase: unable to delete all from purchaseOrderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "purchase: failed to get rows affected by deleteall for purchase_order_item")
	}

	if len(purchaseOrderItemAfterDeleteHooks) != 0 {
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
func (o *PurchaseOrderItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPurchaseOrderItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PurchaseOrderItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PurchaseOrderItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), purchaseOrderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"purchase\".\"purchase_order_item\".* FROM \"purchase\".\"purchase_order_item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, purchaseOrderItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "purchase: unable to reload all in PurchaseOrderItemSlice")
	}

	*o = slice

	return nil
}

// PurchaseOrderItemExists checks if the PurchaseOrderItem row exists.
func PurchaseOrderItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"purchase\".\"purchase_order_item\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "purchase: unable to check if purchase_order_item exists")
	}

	return exists, nil
}

// Exists checks if the PurchaseOrderItem row exists.
func (o *PurchaseOrderItem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PurchaseOrderItemExists(ctx, exec, o.ID)
}
