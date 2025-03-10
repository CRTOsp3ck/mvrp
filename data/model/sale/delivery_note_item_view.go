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

// DeliveryNoteItemView is an object representing the database table.
type DeliveryNoteItemView struct {
	ID                 null.Int    `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	BaseDocumentItemID null.Int    `boil:"base_document_item_id" json:"base_document_item_id,omitempty" toml:"base_document_item_id" yaml:"base_document_item_id,omitempty"`
	DeliveryNoteID     null.Int    `boil:"delivery_note_id" json:"delivery_note_id,omitempty" toml:"delivery_note_id" yaml:"delivery_note_id,omitempty"`
	GoodsCondition     null.String `boil:"goods_condition" json:"goods_condition,omitempty" toml:"goods_condition" yaml:"goods_condition,omitempty"`
	CreatedAt          null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt          null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt          null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	BaseDocumentItem   null.JSON   `boil:"base_document_item" json:"base_document_item,omitempty" toml:"base_document_item" yaml:"base_document_item,omitempty"`
	InventoryInfo      null.JSON   `boil:"inventory_info" json:"inventory_info,omitempty" toml:"inventory_info" yaml:"inventory_info,omitempty"`
}

var DeliveryNoteItemViewColumns = struct {
	ID                 string
	BaseDocumentItemID string
	DeliveryNoteID     string
	GoodsCondition     string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
	BaseDocumentItem   string
	InventoryInfo      string
}{
	ID:                 "id",
	BaseDocumentItemID: "base_document_item_id",
	DeliveryNoteID:     "delivery_note_id",
	GoodsCondition:     "goods_condition",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
	BaseDocumentItem:   "base_document_item",
	InventoryInfo:      "inventory_info",
}

var DeliveryNoteItemViewTableColumns = struct {
	ID                 string
	BaseDocumentItemID string
	DeliveryNoteID     string
	GoodsCondition     string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
	BaseDocumentItem   string
	InventoryInfo      string
}{
	ID:                 "delivery_note_item_view.id",
	BaseDocumentItemID: "delivery_note_item_view.base_document_item_id",
	DeliveryNoteID:     "delivery_note_item_view.delivery_note_id",
	GoodsCondition:     "delivery_note_item_view.goods_condition",
	CreatedAt:          "delivery_note_item_view.created_at",
	UpdatedAt:          "delivery_note_item_view.updated_at",
	DeletedAt:          "delivery_note_item_view.deleted_at",
	BaseDocumentItem:   "delivery_note_item_view.base_document_item",
	InventoryInfo:      "delivery_note_item_view.inventory_info",
}

// Generated where

var DeliveryNoteItemViewWhere = struct {
	ID                 whereHelpernull_Int
	BaseDocumentItemID whereHelpernull_Int
	DeliveryNoteID     whereHelpernull_Int
	GoodsCondition     whereHelpernull_String
	CreatedAt          whereHelpernull_Time
	UpdatedAt          whereHelpernull_Time
	DeletedAt          whereHelpernull_Time
	BaseDocumentItem   whereHelpernull_JSON
	InventoryInfo      whereHelpernull_JSON
}{
	ID:                 whereHelpernull_Int{field: "\"sale\".\"delivery_note_item_view\".\"id\""},
	BaseDocumentItemID: whereHelpernull_Int{field: "\"sale\".\"delivery_note_item_view\".\"base_document_item_id\""},
	DeliveryNoteID:     whereHelpernull_Int{field: "\"sale\".\"delivery_note_item_view\".\"delivery_note_id\""},
	GoodsCondition:     whereHelpernull_String{field: "\"sale\".\"delivery_note_item_view\".\"goods_condition\""},
	CreatedAt:          whereHelpernull_Time{field: "\"sale\".\"delivery_note_item_view\".\"created_at\""},
	UpdatedAt:          whereHelpernull_Time{field: "\"sale\".\"delivery_note_item_view\".\"updated_at\""},
	DeletedAt:          whereHelpernull_Time{field: "\"sale\".\"delivery_note_item_view\".\"deleted_at\""},
	BaseDocumentItem:   whereHelpernull_JSON{field: "\"sale\".\"delivery_note_item_view\".\"base_document_item\""},
	InventoryInfo:      whereHelpernull_JSON{field: "\"sale\".\"delivery_note_item_view\".\"inventory_info\""},
}

var (
	deliveryNoteItemViewAllColumns            = []string{"id", "base_document_item_id", "delivery_note_id", "goods_condition", "created_at", "updated_at", "deleted_at", "base_document_item", "inventory_info"}
	deliveryNoteItemViewColumnsWithoutDefault = []string{}
	deliveryNoteItemViewColumnsWithDefault    = []string{"id", "base_document_item_id", "delivery_note_id", "goods_condition", "created_at", "updated_at", "deleted_at", "base_document_item", "inventory_info"}
	deliveryNoteItemViewPrimaryKeyColumns     = []string{}
	deliveryNoteItemViewGeneratedColumns      = []string{}
)

type (
	// DeliveryNoteItemViewSlice is an alias for a slice of pointers to DeliveryNoteItemView.
	// This should almost always be used instead of []DeliveryNoteItemView.
	DeliveryNoteItemViewSlice []*DeliveryNoteItemView
	// DeliveryNoteItemViewHook is the signature for custom DeliveryNoteItemView hook methods
	DeliveryNoteItemViewHook func(context.Context, boil.ContextExecutor, *DeliveryNoteItemView) error

	deliveryNoteItemViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deliveryNoteItemViewType           = reflect.TypeOf(&DeliveryNoteItemView{})
	deliveryNoteItemViewMapping        = queries.MakeStructMapping(deliveryNoteItemViewType)
	deliveryNoteItemViewInsertCacheMut sync.RWMutex
	deliveryNoteItemViewInsertCache    = make(map[string]insertCache)
	deliveryNoteItemViewUpdateCacheMut sync.RWMutex
	deliveryNoteItemViewUpdateCache    = make(map[string]updateCache)
	deliveryNoteItemViewUpsertCacheMut sync.RWMutex
	deliveryNoteItemViewUpsertCache    = make(map[string]insertCache)
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

var deliveryNoteItemViewAfterSelectMu sync.Mutex
var deliveryNoteItemViewAfterSelectHooks []DeliveryNoteItemViewHook

var deliveryNoteItemViewBeforeInsertMu sync.Mutex
var deliveryNoteItemViewBeforeInsertHooks []DeliveryNoteItemViewHook
var deliveryNoteItemViewAfterInsertMu sync.Mutex
var deliveryNoteItemViewAfterInsertHooks []DeliveryNoteItemViewHook

var deliveryNoteItemViewBeforeUpsertMu sync.Mutex
var deliveryNoteItemViewBeforeUpsertHooks []DeliveryNoteItemViewHook
var deliveryNoteItemViewAfterUpsertMu sync.Mutex
var deliveryNoteItemViewAfterUpsertHooks []DeliveryNoteItemViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DeliveryNoteItemView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryNoteItemViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DeliveryNoteItemView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryNoteItemViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DeliveryNoteItemView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryNoteItemViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DeliveryNoteItemView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryNoteItemViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DeliveryNoteItemView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deliveryNoteItemViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeliveryNoteItemViewHook registers your hook function for all future operations.
func AddDeliveryNoteItemViewHook(hookPoint boil.HookPoint, deliveryNoteItemViewHook DeliveryNoteItemViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		deliveryNoteItemViewAfterSelectMu.Lock()
		deliveryNoteItemViewAfterSelectHooks = append(deliveryNoteItemViewAfterSelectHooks, deliveryNoteItemViewHook)
		deliveryNoteItemViewAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		deliveryNoteItemViewBeforeInsertMu.Lock()
		deliveryNoteItemViewBeforeInsertHooks = append(deliveryNoteItemViewBeforeInsertHooks, deliveryNoteItemViewHook)
		deliveryNoteItemViewBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		deliveryNoteItemViewAfterInsertMu.Lock()
		deliveryNoteItemViewAfterInsertHooks = append(deliveryNoteItemViewAfterInsertHooks, deliveryNoteItemViewHook)
		deliveryNoteItemViewAfterInsertMu.Unlock()
	case boil.BeforeUpsertHook:
		deliveryNoteItemViewBeforeUpsertMu.Lock()
		deliveryNoteItemViewBeforeUpsertHooks = append(deliveryNoteItemViewBeforeUpsertHooks, deliveryNoteItemViewHook)
		deliveryNoteItemViewBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		deliveryNoteItemViewAfterUpsertMu.Lock()
		deliveryNoteItemViewAfterUpsertHooks = append(deliveryNoteItemViewAfterUpsertHooks, deliveryNoteItemViewHook)
		deliveryNoteItemViewAfterUpsertMu.Unlock()
	}
}

// One returns a single deliveryNoteItemView record from the query.
func (q deliveryNoteItemViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeliveryNoteItemView, error) {
	o := &DeliveryNoteItemView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: failed to execute a one query for delivery_note_item_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DeliveryNoteItemView records from the query.
func (q deliveryNoteItemViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeliveryNoteItemViewSlice, error) {
	var o []*DeliveryNoteItemView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sale: failed to assign all query results to DeliveryNoteItemView slice")
	}

	if len(deliveryNoteItemViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DeliveryNoteItemView records in the query.
func (q deliveryNoteItemViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to count delivery_note_item_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deliveryNoteItemViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sale: failed to check if delivery_note_item_view exists")
	}

	return count > 0, nil
}

// DeliveryNoteItemViews retrieves all the records using an executor.
func DeliveryNoteItemViews(mods ...qm.QueryMod) deliveryNoteItemViewQuery {
	mods = append(mods, qm.From("\"sale\".\"delivery_note_item_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sale\".\"delivery_note_item_view\".*"})
	}

	return deliveryNoteItemViewQuery{q}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DeliveryNoteItemView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sale: no delivery_note_item_view provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deliveryNoteItemViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deliveryNoteItemViewInsertCacheMut.RLock()
	cache, cached := deliveryNoteItemViewInsertCache[key]
	deliveryNoteItemViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deliveryNoteItemViewAllColumns,
			deliveryNoteItemViewColumnsWithDefault,
			deliveryNoteItemViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deliveryNoteItemViewType, deliveryNoteItemViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deliveryNoteItemViewType, deliveryNoteItemViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sale\".\"delivery_note_item_view\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sale\".\"delivery_note_item_view\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sale: unable to insert into delivery_note_item_view")
	}

	if !cached {
		deliveryNoteItemViewInsertCacheMut.Lock()
		deliveryNoteItemViewInsertCache[key] = cache
		deliveryNoteItemViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DeliveryNoteItemView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("sale: no delivery_note_item_view provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(deliveryNoteItemViewColumnsWithDefault, o)

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

	deliveryNoteItemViewUpsertCacheMut.RLock()
	cache, cached := deliveryNoteItemViewUpsertCache[key]
	deliveryNoteItemViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			deliveryNoteItemViewAllColumns,
			deliveryNoteItemViewColumnsWithDefault,
			deliveryNoteItemViewColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			deliveryNoteItemViewAllColumns,
			deliveryNoteItemViewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sale: unable to upsert delivery_note_item_view, could not build update column list")
		}

		ret := strmangle.SetComplement(deliveryNoteItemViewAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(deliveryNoteItemViewPrimaryKeyColumns) == 0 {
				return errors.New("sale: unable to upsert delivery_note_item_view, could not build conflict column list")
			}

			conflict = make([]string, len(deliveryNoteItemViewPrimaryKeyColumns))
			copy(conflict, deliveryNoteItemViewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"sale\".\"delivery_note_item_view\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(deliveryNoteItemViewType, deliveryNoteItemViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deliveryNoteItemViewType, deliveryNoteItemViewMapping, ret)
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
		return errors.Wrap(err, "sale: unable to upsert delivery_note_item_view")
	}

	if !cached {
		deliveryNoteItemViewUpsertCacheMut.Lock()
		deliveryNoteItemViewUpsertCache[key] = cache
		deliveryNoteItemViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
