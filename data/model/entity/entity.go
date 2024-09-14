// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entity

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

// Entity is an object representing the database table.
type Entity struct {
	ID          int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Code        string       `boil:"code" json:"code" toml:"code" yaml:"code"`
	Name        string       `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description string       `boil:"description" json:"description" toml:"description" yaml:"description"`
	Address     null.String  `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	Phone       null.String  `boil:"phone" json:"phone,omitempty" toml:"phone" yaml:"phone,omitempty"`
	Email       null.String  `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	Website     null.String  `boil:"website" json:"website,omitempty" toml:"website" yaml:"website,omitempty"`
	EntityType  EntityType   `boil:"entity_type" json:"entity_type" toml:"entity_type" yaml:"entity_type"`
	Status      EntityStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt   time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt   null.Time    `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *entityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L entityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EntityColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	Address     string
	Phone       string
	Email       string
	Website     string
	EntityType  string
	Status      string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Description: "description",
	Address:     "address",
	Phone:       "phone",
	Email:       "email",
	Website:     "website",
	EntityType:  "entity_type",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

var EntityTableColumns = struct {
	ID          string
	Code        string
	Name        string
	Description string
	Address     string
	Phone       string
	Email       string
	Website     string
	EntityType  string
	Status      string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "entity.id",
	Code:        "entity.code",
	Name:        "entity.name",
	Description: "entity.description",
	Address:     "entity.address",
	Phone:       "entity.phone",
	Email:       "entity.email",
	Website:     "entity.website",
	EntityType:  "entity.entity_type",
	Status:      "entity.status",
	CreatedAt:   "entity.created_at",
	UpdatedAt:   "entity.updated_at",
	DeletedAt:   "entity.deleted_at",
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

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperEntityType struct{ field string }

func (w whereHelperEntityType) EQ(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperEntityType) NEQ(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperEntityType) LT(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEntityType) LTE(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEntityType) GT(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEntityType) GTE(x EntityType) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEntityType) IN(slice []EntityType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEntityType) NIN(slice []EntityType) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperEntityStatus struct{ field string }

func (w whereHelperEntityStatus) EQ(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperEntityStatus) NEQ(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperEntityStatus) LT(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperEntityStatus) LTE(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperEntityStatus) GT(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperEntityStatus) GTE(x EntityStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperEntityStatus) IN(slice []EntityStatus) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperEntityStatus) NIN(slice []EntityStatus) qm.QueryMod {
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

var EntityWhere = struct {
	ID          whereHelperint
	Code        whereHelperstring
	Name        whereHelperstring
	Description whereHelperstring
	Address     whereHelpernull_String
	Phone       whereHelpernull_String
	Email       whereHelpernull_String
	Website     whereHelpernull_String
	EntityType  whereHelperEntityType
	Status      whereHelperEntityStatus
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
	DeletedAt   whereHelpernull_Time
}{
	ID:          whereHelperint{field: "\"entity\".\"entity\".\"id\""},
	Code:        whereHelperstring{field: "\"entity\".\"entity\".\"code\""},
	Name:        whereHelperstring{field: "\"entity\".\"entity\".\"name\""},
	Description: whereHelperstring{field: "\"entity\".\"entity\".\"description\""},
	Address:     whereHelpernull_String{field: "\"entity\".\"entity\".\"address\""},
	Phone:       whereHelpernull_String{field: "\"entity\".\"entity\".\"phone\""},
	Email:       whereHelpernull_String{field: "\"entity\".\"entity\".\"email\""},
	Website:     whereHelpernull_String{field: "\"entity\".\"entity\".\"website\""},
	EntityType:  whereHelperEntityType{field: "\"entity\".\"entity\".\"entity_type\""},
	Status:      whereHelperEntityStatus{field: "\"entity\".\"entity\".\"status\""},
	CreatedAt:   whereHelpertime_Time{field: "\"entity\".\"entity\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"entity\".\"entity\".\"updated_at\""},
	DeletedAt:   whereHelpernull_Time{field: "\"entity\".\"entity\".\"deleted_at\""},
}

// EntityRels is where relationship names are stored.
var EntityRels = struct {
}{}

// entityR is where relationships are stored.
type entityR struct {
}

// NewStruct creates a new relationship struct
func (*entityR) NewStruct() *entityR {
	return &entityR{}
}

// entityL is where Load methods for each relationship are stored.
type entityL struct{}

var (
	entityAllColumns            = []string{"id", "code", "name", "description", "address", "phone", "email", "website", "entity_type", "status", "created_at", "updated_at", "deleted_at"}
	entityColumnsWithoutDefault = []string{"id", "code", "name", "description", "entity_type", "status", "created_at", "updated_at"}
	entityColumnsWithDefault    = []string{"address", "phone", "email", "website", "deleted_at"}
	entityPrimaryKeyColumns     = []string{"id"}
	entityGeneratedColumns      = []string{}
)

type (
	// EntitySlice is an alias for a slice of pointers to Entity.
	// This should almost always be used instead of []Entity.
	EntitySlice []*Entity
	// EntityHook is the signature for custom Entity hook methods
	EntityHook func(context.Context, boil.ContextExecutor, *Entity) error

	entityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	entityType                 = reflect.TypeOf(&Entity{})
	entityMapping              = queries.MakeStructMapping(entityType)
	entityPrimaryKeyMapping, _ = queries.BindMapping(entityType, entityMapping, entityPrimaryKeyColumns)
	entityInsertCacheMut       sync.RWMutex
	entityInsertCache          = make(map[string]insertCache)
	entityUpdateCacheMut       sync.RWMutex
	entityUpdateCache          = make(map[string]updateCache)
	entityUpsertCacheMut       sync.RWMutex
	entityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var entityAfterSelectMu sync.Mutex
var entityAfterSelectHooks []EntityHook

var entityBeforeInsertMu sync.Mutex
var entityBeforeInsertHooks []EntityHook
var entityAfterInsertMu sync.Mutex
var entityAfterInsertHooks []EntityHook

var entityBeforeUpdateMu sync.Mutex
var entityBeforeUpdateHooks []EntityHook
var entityAfterUpdateMu sync.Mutex
var entityAfterUpdateHooks []EntityHook

var entityBeforeDeleteMu sync.Mutex
var entityBeforeDeleteHooks []EntityHook
var entityAfterDeleteMu sync.Mutex
var entityAfterDeleteHooks []EntityHook

var entityBeforeUpsertMu sync.Mutex
var entityBeforeUpsertHooks []EntityHook
var entityAfterUpsertMu sync.Mutex
var entityAfterUpsertHooks []EntityHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Entity) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Entity) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Entity) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Entity) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Entity) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Entity) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Entity) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Entity) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Entity) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range entityAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEntityHook registers your hook function for all future operations.
func AddEntityHook(hookPoint boil.HookPoint, entityHook EntityHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		entityAfterSelectMu.Lock()
		entityAfterSelectHooks = append(entityAfterSelectHooks, entityHook)
		entityAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		entityBeforeInsertMu.Lock()
		entityBeforeInsertHooks = append(entityBeforeInsertHooks, entityHook)
		entityBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		entityAfterInsertMu.Lock()
		entityAfterInsertHooks = append(entityAfterInsertHooks, entityHook)
		entityAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		entityBeforeUpdateMu.Lock()
		entityBeforeUpdateHooks = append(entityBeforeUpdateHooks, entityHook)
		entityBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		entityAfterUpdateMu.Lock()
		entityAfterUpdateHooks = append(entityAfterUpdateHooks, entityHook)
		entityAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		entityBeforeDeleteMu.Lock()
		entityBeforeDeleteHooks = append(entityBeforeDeleteHooks, entityHook)
		entityBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		entityAfterDeleteMu.Lock()
		entityAfterDeleteHooks = append(entityAfterDeleteHooks, entityHook)
		entityAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		entityBeforeUpsertMu.Lock()
		entityBeforeUpsertHooks = append(entityBeforeUpsertHooks, entityHook)
		entityBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		entityAfterUpsertMu.Lock()
		entityAfterUpsertHooks = append(entityAfterUpsertHooks, entityHook)
		entityAfterUpsertMu.Unlock()
	}
}

// One returns a single entity record from the query.
func (q entityQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Entity, error) {
	o := &Entity{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: failed to execute a one query for entity")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Entity records from the query.
func (q entityQuery) All(ctx context.Context, exec boil.ContextExecutor) (EntitySlice, error) {
	var o []*Entity

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "entity: failed to assign all query results to Entity slice")
	}

	if len(entityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Entity records in the query.
func (q entityQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to count entity rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q entityQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "entity: failed to check if entity exists")
	}

	return count > 0, nil
}

// Entities retrieves all the records using an executor.
func Entities(mods ...qm.QueryMod) entityQuery {
	mods = append(mods, qm.From("\"entity\".\"entity\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"entity\".\"entity\".*"})
	}

	return entityQuery{q}
}

// FindEntity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEntity(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Entity, error) {
	entityObj := &Entity{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"entity\".\"entity\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, entityObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "entity: unable to select from entity")
	}

	if err = entityObj.doAfterSelectHooks(ctx, exec); err != nil {
		return entityObj, err
	}

	return entityObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Entity) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("entity: no entity provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(entityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	entityInsertCacheMut.RLock()
	cache, cached := entityInsertCache[key]
	entityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			entityAllColumns,
			entityColumnsWithDefault,
			entityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(entityType, entityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(entityType, entityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"entity\".\"entity\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"entity\".\"entity\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "entity: unable to insert into entity")
	}

	if !cached {
		entityInsertCacheMut.Lock()
		entityInsertCache[key] = cache
		entityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Entity.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Entity) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	entityUpdateCacheMut.RLock()
	cache, cached := entityUpdateCache[key]
	entityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			entityAllColumns,
			entityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("entity: unable to update entity, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"entity\".\"entity\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, entityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(entityType, entityMapping, append(wl, entityPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "entity: unable to update entity row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by update for entity")
	}

	if !cached {
		entityUpdateCacheMut.Lock()
		entityUpdateCache[key] = cache
		entityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q entityQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all for entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected for entity")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EntitySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("entity: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"entity\".\"entity\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, entityPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to update all in entity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to retrieve rows affected all in update all entity")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Entity) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("entity: no entity provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(entityColumnsWithDefault, o)

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

	entityUpsertCacheMut.RLock()
	cache, cached := entityUpsertCache[key]
	entityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			entityAllColumns,
			entityColumnsWithDefault,
			entityColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			entityAllColumns,
			entityPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("entity: unable to upsert entity, could not build update column list")
		}

		ret := strmangle.SetComplement(entityAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(entityPrimaryKeyColumns) == 0 {
				return errors.New("entity: unable to upsert entity, could not build conflict column list")
			}

			conflict = make([]string, len(entityPrimaryKeyColumns))
			copy(conflict, entityPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"entity\".\"entity\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(entityType, entityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(entityType, entityMapping, ret)
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
		return errors.Wrap(err, "entity: unable to upsert entity")
	}

	if !cached {
		entityUpsertCacheMut.Lock()
		entityUpsertCache[key] = cache
		entityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Entity record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Entity) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("entity: no Entity provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), entityPrimaryKeyMapping)
	sql := "DELETE FROM \"entity\".\"entity\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete from entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by delete for entity")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q entityQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("entity: no entityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from entity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for entity")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EntitySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(entityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"entity\".\"entity\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entityPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "entity: unable to delete all from entity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "entity: failed to get rows affected by deleteall for entity")
	}

	if len(entityAfterDeleteHooks) != 0 {
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
func (o *Entity) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEntity(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EntitySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EntitySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), entityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"entity\".\"entity\".* FROM \"entity\".\"entity\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, entityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "entity: unable to reload all in EntitySlice")
	}

	*o = slice

	return nil
}

// EntityExists checks if the Entity row exists.
func EntityExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"entity\".\"entity\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "entity: unable to check if entity exists")
	}

	return exists, nil
}

// Exists checks if the Entity row exists.
func (o *Entity) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EntityExists(ctx, exec, o.ID)
}
