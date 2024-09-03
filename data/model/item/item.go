// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package item

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
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// Item is an object representing the database table.
type Item struct {
	ID          int               `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code        string            `boil:"code" json:"code" toml:"code" yaml:"code"`
	Name        string            `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string            `boil:"description" json:"description" toml:"description" yaml:"description"`
	Price       types.NullDecimal `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
	Cost        types.NullDecimal `boil:"cost" json:"cost,omitempty" toml:"cost" yaml:"cost,omitempty"`
	Type        ItemType          `boil:"type" json:"type" toml:"type" yaml:"type"`
	CreatedAt   time.Time         `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time         `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt   null.Time         `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *itemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L itemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ItemColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	Price       string
	Cost        string
	Type        string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Description: "description",
	Price:       "price",
	Cost:        "cost",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var ItemTableColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	Price       string
	Cost        string
	Type        string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "item.id",
	Code:        "item.code",
	Name:        "item.name",
	Description: "item.description",
	Price:       "item.price",
	Cost:        "item.cost",
	Type:        "item.type",
	CreatedAt:   "item.created_at",
	UpdatedAt:   "item.updated_at",
	DeletedAt:   "item.deleted_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertypes_NullDecimal struct{ field string }

func (w whereHelpertypes_NullDecimal) EQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_NullDecimal) NEQ(x types.NullDecimal) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_NullDecimal) LT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_NullDecimal) LTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_NullDecimal) GT(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_NullDecimal) GTE(x types.NullDecimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpertypes_NullDecimal) IsNull() qm.QueryMod { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_NullDecimal) IsNotNull() qm.QueryMod {
	return qmhelper.WhereIsNotNull(w.field)
}

type whereHelperItemType struct{ field string }

func (w whereHelperItemType) EQ(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperItemType) NEQ(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperItemType) LT(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperItemType) LTE(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperItemType) GT(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperItemType) GTE(x ItemType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperItemType) IN(slice []ItemType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperItemType) NIN(slice []ItemType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ItemWhere = struct {
	ID          whereHelperint
	Code        whereHelperstring
	Name        whereHelperstring
	Description whereHelperstring
	Price       whereHelpertypes_NullDecimal
	Cost        whereHelpertypes_NullDecimal
	Type        whereHelperItemType
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperint{field: "\"item\".\"item\".\"id\""},
	Code:        whereHelperstring{field: "\"item\".\"item\".\"code\""},
	Name:        whereHelperstring{field: "\"item\".\"item\".\"name\""},
	Description: whereHelperstring{field: "\"item\".\"item\".\"description\""},
	Price:       whereHelpertypes_NullDecimal{field: "\"item\".\"item\".\"price\""},
	Cost:        whereHelpertypes_NullDecimal{field: "\"item\".\"item\".\"cost\""},
	Type:        whereHelperItemType{field: "\"item\".\"item\".\"type\""},
	CreatedAt:   whereHelpertime_Time{field: "\"item\".\"item\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"item\".\"item\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"item\".\"item\".\"deleted_at\""},
}

// ItemRels is where relationship names are stored.
var ItemRels = struct {
}{}

// itemR is where relationships are stored.
type itemR struct {
}

// NewStruct creates a new relationship struct
func (*itemR) NewStruct() *itemR {
	return &itemR{}
}

// itemL is where Load methods for each relationship are stored.
type itemL struct{}

var (
	itemAllColumns            = []string{"id", "code", "name", "description", "price", "cost", "type", "created_at", "updated_at", "deleted_at"}
	itemColumnsWithoutDefault = []string{"id", "code", "name", "description", "type", "created_at", "updated_at"}
	itemColumnsWithDefault    = []string{"price", "cost", "deleted_at"}
	itemPrimaryKeyColumns     = []string{"id"}
	itemGeneratedColumns      = []string{}
)

type (
	// ItemSlice is an alias for a slice of pointers to Item.
	// This should almost always be used instead of []Item.
	ItemSlice []*Item
	// ItemHook is the signature for custom Item hook methods
	ItemHook func(context.Context, boil.ContextExecutor, *Item) error

	itemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	itemType                 = reflect.TypeOf(&Item{})
	itemMapping              = queries.MakeStructMapping(itemType)
	itemPrimaryKeyMapping, _ = queries.BindMapping(itemType, itemMapping, itemPrimaryKeyColumns)
	itemInsertCacheMut       sync.RWMutex
	itemInsertCache          = make(map[string]insertCache)
	itemUpdateCacheMut       sync.RWMutex
	itemUpdateCache          = make(map[string]updateCache)
	itemUpsertCacheMut       sync.RWMutex
	itemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var itemAfterSelectMu sync.Mutex
var itemAfterSelectHooks []ItemHook

var itemBeforeInsertMu sync.Mutex
var itemBeforeInsertHooks []ItemHook
var itemAfterInsertMu sync.Mutex
var itemAfterInsertHooks []ItemHook

var itemBeforeUpdateMu sync.Mutex
var itemBeforeUpdateHooks []ItemHook
var itemAfterUpdateMu sync.Mutex
var itemAfterUpdateHooks []ItemHook

var itemBeforeDeleteMu sync.Mutex
var itemBeforeDeleteHooks []ItemHook
var itemAfterDeleteMu sync.Mutex
var itemAfterDeleteHooks []ItemHook

var itemBeforeUpsertMu sync.Mutex
var itemBeforeUpsertHooks []ItemHook
var itemAfterUpsertMu sync.Mutex
var itemAfterUpsertHooks []ItemHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Item) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Item) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Item) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Item) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Item) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Item) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Item) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Item) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Item) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range itemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddItemHook registers your hook function for all future operations.
func AddItemHook(hookPoint boil.HookPoint, itemHook ItemHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		itemAfterSelectMu.Lock()
		itemAfterSelectHooks = append(itemAfterSelectHooks, itemHook)
		itemAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		itemBeforeInsertMu.Lock()
		itemBeforeInsertHooks = append(itemBeforeInsertHooks, itemHook)
		itemBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		itemAfterInsertMu.Lock()
		itemAfterInsertHooks = append(itemAfterInsertHooks, itemHook)
		itemAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		itemBeforeUpdateMu.Lock()
		itemBeforeUpdateHooks = append(itemBeforeUpdateHooks, itemHook)
		itemBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		itemAfterUpdateMu.Lock()
		itemAfterUpdateHooks = append(itemAfterUpdateHooks, itemHook)
		itemAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		itemBeforeDeleteMu.Lock()
		itemBeforeDeleteHooks = append(itemBeforeDeleteHooks, itemHook)
		itemBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		itemAfterDeleteMu.Lock()
		itemAfterDeleteHooks = append(itemAfterDeleteHooks, itemHook)
		itemAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		itemBeforeUpsertMu.Lock()
		itemBeforeUpsertHooks = append(itemBeforeUpsertHooks, itemHook)
		itemBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		itemAfterUpsertMu.Lock()
		itemAfterUpsertHooks = append(itemAfterUpsertHooks, itemHook)
		itemAfterUpsertMu.Unlock()
	}
}

// One returns a single item record from the query.
func (q itemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Item, error) {
	o := &Item{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "item: failed to execute a one query for item")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Item records from the query.
func (q itemQuery) All(ctx context.Context, exec boil.ContextExecutor) (ItemSlice, error) {
	var o []*Item

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "item: failed to assign all query results to Item slice")
	}

	if len(itemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Item records in the query.
func (q itemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "item: failed to count item rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q itemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "item: failed to check if item exists")
	}

	return count > 0, nil
}

// Items retrieves all the records using an executor.
func Items(mods ...qm.QueryMod) itemQuery {
	mods = append(mods, qm.From("\"item\".\"item\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"item\".\"item\".*"})
	}

	return itemQuery{q}
}

// FindItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Item, error) {
	itemObj := &Item{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"item\".\"item\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, itemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "item: unable to select from item")
	}

	if err = itemObj.doAfterSelectHooks(ctx, exec); err != nil {
		return itemObj, err
	}

	return itemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Item) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("item: no item provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(itemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	itemInsertCacheMut.RLock()
	cache, cached := itemInsertCache[key]
	itemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			itemAllColumns,
			itemColumnsWithDefault,
			itemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(itemType, itemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"item\".\"item\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"item\".\"item\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "item: unable to insert into item")
	}

	if !cached {
		itemInsertCacheMut.Lock()
		itemInsertCache[key] = cache
		itemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Item.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Item) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	itemUpdateCacheMut.RLock()
	cache, cached := itemUpdateCache[key]
	itemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			itemAllColumns,
			itemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("item: unable to update item, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"item\".\"item\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, itemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, append(wl, itemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "item: unable to update item row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: failed to get rows affected by update for item")
	}

	if !cached {
		itemUpdateCacheMut.Lock()
		itemUpdateCache[key] = cache
		itemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q itemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to update all for item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to retrieve rows affected for item")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("item: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"item\".\"item\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, itemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to update all in item slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to retrieve rows affected all in update all item")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Item) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("item: no item provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(itemColumnsWithDefault, o)

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

	itemUpsertCacheMut.RLock()
	cache, cached := itemUpsertCache[key]
	itemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			itemAllColumns,
			itemColumnsWithDefault,
			itemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			itemAllColumns,
			itemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("item: unable to upsert item, could not build update column list")
		}

		ret := strmangle.SetComplement(itemAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(itemPrimaryKeyColumns) == 0 {
				return errors.New("item: unable to upsert item, could not build conflict column list")
			}

			conflict = make([]string, len(itemPrimaryKeyColumns))
			copy(conflict, itemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"item\".\"item\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(itemType, itemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(itemType, itemMapping, ret)
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
		return errors.Wrap(err, "item: unable to upsert item")
	}

	if !cached {
		itemUpsertCacheMut.Lock()
		itemUpsertCache[key] = cache
		itemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Item record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Item) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("item: no Item provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), itemPrimaryKeyMapping)
	sql := "DELETE FROM \"item\".\"item\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to delete from item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: failed to get rows affected by delete for item")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q itemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("item: no itemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to delete all from item")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: failed to get rows affected by deleteall for item")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(itemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"item\".\"item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, itemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "item: unable to delete all from item slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "item: failed to get rows affected by deleteall for item")
	}

	if len(itemAfterDeleteHooks) != 0 {
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
func (o *Item) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"item\".\"item\".* FROM \"item\".\"item\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, itemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "item: unable to reload all in ItemSlice")
	}

	*o = slice

	return nil
}

// ItemExists checks if the Item row exists.
func ItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"item\".\"item\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "item: unable to check if item exists")
	}

	return exists, nil
}

// Exists checks if the Item row exists.
func (o *Item) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ItemExists(ctx, exec, o.ID)
}
