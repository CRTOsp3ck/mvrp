// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package invoice

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

// CreditNoteView is an object representing the database table.
type CreditNoteView struct {
	ID                null.Int          `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	BaseDocumentID    null.Int          `boil:"base_document_id" json:"base_document_id,omitempty" toml:"base_document_id" yaml:"base_document_id,omitempty"`
	CreditNoteNumber  null.String       `boil:"credit_note_number" json:"credit_note_number,omitempty" toml:"credit_note_number" yaml:"credit_note_number,omitempty"`
	CustomerID        null.Int          `boil:"customer_id" json:"customer_id,omitempty" toml:"customer_id" yaml:"customer_id,omitempty"`
	AdditionalCharges types.NullDecimal `boil:"additional_charges" json:"additional_charges,omitempty" toml:"additional_charges" yaml:"additional_charges,omitempty"`
	TotalValueGen     types.NullDecimal `boil:"total_value_gen" json:"total_value_gen,omitempty" toml:"total_value_gen" yaml:"total_value_gen,omitempty"`
	IssueDate         null.Time         `boil:"issue_date" json:"issue_date,omitempty" toml:"issue_date" yaml:"issue_date,omitempty"`
	ReasonForIssuance null.String       `boil:"reason_for_issuance" json:"reason_for_issuance,omitempty" toml:"reason_for_issuance" yaml:"reason_for_issuance,omitempty"`
	CreatedAt         null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt         null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt         null.Time         `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	BaseDocument      null.JSON         `boil:"base_document" json:"base_document,omitempty" toml:"base_document" yaml:"base_document,omitempty"`
	CustomerInfo      null.JSON         `boil:"customer_info" json:"customer_info,omitempty" toml:"customer_info" yaml:"customer_info,omitempty"`
	CreditNoteItems   null.JSON         `boil:"credit_note_items" json:"credit_note_items,omitempty" toml:"credit_note_items" yaml:"credit_note_items,omitempty"`
}

var CreditNoteViewColumns = struct {
	ID                string
	BaseDocumentID    string
	CreditNoteNumber  string
	CustomerID        string
	AdditionalCharges string
	TotalValueGen     string
	IssueDate         string
	ReasonForIssuance string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	BaseDocument      string
	CustomerInfo      string
	CreditNoteItems   string
}{
	ID:                "id",
	BaseDocumentID:    "base_document_id",
	CreditNoteNumber:  "credit_note_number",
	CustomerID:        "customer_id",
	AdditionalCharges: "additional_charges",
	TotalValueGen:     "total_value_gen",
	IssueDate:         "issue_date",
	ReasonForIssuance: "reason_for_issuance",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
	BaseDocument:      "base_document",
	CustomerInfo:      "customer_info",
	CreditNoteItems:   "credit_note_items",
}

var CreditNoteViewTableColumns = struct {
	ID                string
	BaseDocumentID    string
	CreditNoteNumber  string
	CustomerID        string
	AdditionalCharges string
	TotalValueGen     string
	IssueDate         string
	ReasonForIssuance string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	BaseDocument      string
	CustomerInfo      string
	CreditNoteItems   string
}{
	ID:                "credit_note_view.id",
	BaseDocumentID:    "credit_note_view.base_document_id",
	CreditNoteNumber:  "credit_note_view.credit_note_number",
	CustomerID:        "credit_note_view.customer_id",
	AdditionalCharges: "credit_note_view.additional_charges",
	TotalValueGen:     "credit_note_view.total_value_gen",
	IssueDate:         "credit_note_view.issue_date",
	ReasonForIssuance: "credit_note_view.reason_for_issuance",
	CreatedAt:         "credit_note_view.created_at",
	UpdatedAt:         "credit_note_view.updated_at",
	DeletedAt:         "credit_note_view.deleted_at",
	BaseDocument:      "credit_note_view.base_document",
	CustomerInfo:      "credit_note_view.customer_info",
	CreditNoteItems:   "credit_note_view.credit_note_items",
}

// Generated where

var CreditNoteViewWhere = struct {
	ID                whereHelpernull_Int
	BaseDocumentID    whereHelpernull_Int
	CreditNoteNumber  whereHelpernull_String
	CustomerID        whereHelpernull_Int
	AdditionalCharges whereHelpertypes_NullDecimal
	TotalValueGen     whereHelpertypes_NullDecimal
	IssueDate         whereHelpernull_Time
	ReasonForIssuance whereHelpernull_String
	CreatedAt         whereHelpernull_Time
	UpdatedAt         whereHelpernull_Time
	DeletedAt         whereHelpernull_Time
	BaseDocument      whereHelpernull_JSON
	CustomerInfo      whereHelpernull_JSON
	CreditNoteItems   whereHelpernull_JSON
}{
	ID:                whereHelpernull_Int{field: "\"invoice\".\"credit_note_view\".\"id\""},
	BaseDocumentID:    whereHelpernull_Int{field: "\"invoice\".\"credit_note_view\".\"base_document_id\""},
	CreditNoteNumber:  whereHelpernull_String{field: "\"invoice\".\"credit_note_view\".\"credit_note_number\""},
	CustomerID:        whereHelpernull_Int{field: "\"invoice\".\"credit_note_view\".\"customer_id\""},
	AdditionalCharges: whereHelpertypes_NullDecimal{field: "\"invoice\".\"credit_note_view\".\"additional_charges\""},
	TotalValueGen:     whereHelpertypes_NullDecimal{field: "\"invoice\".\"credit_note_view\".\"total_value_gen\""},
	IssueDate:         whereHelpernull_Time{field: "\"invoice\".\"credit_note_view\".\"issue_date\""},
	ReasonForIssuance: whereHelpernull_String{field: "\"invoice\".\"credit_note_view\".\"reason_for_issuance\""},
	CreatedAt:         whereHelpernull_Time{field: "\"invoice\".\"credit_note_view\".\"created_at\""},
	UpdatedAt:         whereHelpernull_Time{field: "\"invoice\".\"credit_note_view\".\"updated_at\""},
	DeletedAt:         whereHelpernull_Time{field: "\"invoice\".\"credit_note_view\".\"deleted_at\""},
	BaseDocument:      whereHelpernull_JSON{field: "\"invoice\".\"credit_note_view\".\"base_document\""},
	CustomerInfo:      whereHelpernull_JSON{field: "\"invoice\".\"credit_note_view\".\"customer_info\""},
	CreditNoteItems:   whereHelpernull_JSON{field: "\"invoice\".\"credit_note_view\".\"credit_note_items\""},
}

var (
	creditNoteViewAllColumns            = []string{"id", "base_document_id", "credit_note_number", "customer_id", "additional_charges", "total_value_gen", "issue_date", "reason_for_issuance", "created_at", "updated_at", "deleted_at", "base_document", "customer_info", "credit_note_items"}
	creditNoteViewColumnsWithoutDefault = []string{}
	creditNoteViewColumnsWithDefault    = []string{"id", "base_document_id", "credit_note_number", "customer_id", "additional_charges", "total_value_gen", "issue_date", "reason_for_issuance", "created_at", "updated_at", "deleted_at", "base_document", "customer_info", "credit_note_items"}
	creditNoteViewPrimaryKeyColumns     = []string{}
	creditNoteViewGeneratedColumns      = []string{}
)

type (
	// CreditNoteViewSlice is an alias for a slice of pointers to CreditNoteView.
	// This should almost always be used instead of []CreditNoteView.
	CreditNoteViewSlice []*CreditNoteView
	// CreditNoteViewHook is the signature for custom CreditNoteView hook methods
	CreditNoteViewHook func(context.Context, boil.ContextExecutor, *CreditNoteView) error

	creditNoteViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	creditNoteViewType           = reflect.TypeOf(&CreditNoteView{})
	creditNoteViewMapping        = queries.MakeStructMapping(creditNoteViewType)
	creditNoteViewInsertCacheMut sync.RWMutex
	creditNoteViewInsertCache    = make(map[string]insertCache)
	creditNoteViewUpdateCacheMut sync.RWMutex
	creditNoteViewUpdateCache    = make(map[string]updateCache)
	creditNoteViewUpsertCacheMut sync.RWMutex
	creditNoteViewUpsertCache    = make(map[string]insertCache)
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

var creditNoteViewAfterSelectMu sync.Mutex
var creditNoteViewAfterSelectHooks []CreditNoteViewHook

var creditNoteViewBeforeInsertMu sync.Mutex
var creditNoteViewBeforeInsertHooks []CreditNoteViewHook
var creditNoteViewAfterInsertMu sync.Mutex
var creditNoteViewAfterInsertHooks []CreditNoteViewHook

var creditNoteViewBeforeUpsertMu sync.Mutex
var creditNoteViewBeforeUpsertHooks []CreditNoteViewHook
var creditNoteViewAfterUpsertMu sync.Mutex
var creditNoteViewAfterUpsertHooks []CreditNoteViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CreditNoteView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range creditNoteViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CreditNoteView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range creditNoteViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CreditNoteView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range creditNoteViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CreditNoteView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range creditNoteViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CreditNoteView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range creditNoteViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCreditNoteViewHook registers your hook function for all future operations.
func AddCreditNoteViewHook(hookPoint boil.HookPoint, creditNoteViewHook CreditNoteViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		creditNoteViewAfterSelectMu.Lock()
		creditNoteViewAfterSelectHooks = append(creditNoteViewAfterSelectHooks, creditNoteViewHook)
		creditNoteViewAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		creditNoteViewBeforeInsertMu.Lock()
		creditNoteViewBeforeInsertHooks = append(creditNoteViewBeforeInsertHooks, creditNoteViewHook)
		creditNoteViewBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		creditNoteViewAfterInsertMu.Lock()
		creditNoteViewAfterInsertHooks = append(creditNoteViewAfterInsertHooks, creditNoteViewHook)
		creditNoteViewAfterInsertMu.Unlock()
	case boil.BeforeUpsertHook:
		creditNoteViewBeforeUpsertMu.Lock()
		creditNoteViewBeforeUpsertHooks = append(creditNoteViewBeforeUpsertHooks, creditNoteViewHook)
		creditNoteViewBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		creditNoteViewAfterUpsertMu.Lock()
		creditNoteViewAfterUpsertHooks = append(creditNoteViewAfterUpsertHooks, creditNoteViewHook)
		creditNoteViewAfterUpsertMu.Unlock()
	}
}

// One returns a single creditNoteView record from the query.
func (q creditNoteViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CreditNoteView, error) {
	o := &CreditNoteView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "invoice: failed to execute a one query for credit_note_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CreditNoteView records from the query.
func (q creditNoteViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (CreditNoteViewSlice, error) {
	var o []*CreditNoteView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "invoice: failed to assign all query results to CreditNoteView slice")
	}

	if len(creditNoteViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CreditNoteView records in the query.
func (q creditNoteViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "invoice: failed to count credit_note_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q creditNoteViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "invoice: failed to check if credit_note_view exists")
	}

	return count > 0, nil
}

// CreditNoteViews retrieves all the records using an executor.
func CreditNoteViews(mods ...qm.QueryMod) creditNoteViewQuery {
	mods = append(mods, qm.From("\"invoice\".\"credit_note_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invoice\".\"credit_note_view\".*"})
	}

	return creditNoteViewQuery{q}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CreditNoteView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("invoice: no credit_note_view provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(creditNoteViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	creditNoteViewInsertCacheMut.RLock()
	cache, cached := creditNoteViewInsertCache[key]
	creditNoteViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			creditNoteViewAllColumns,
			creditNoteViewColumnsWithDefault,
			creditNoteViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(creditNoteViewType, creditNoteViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(creditNoteViewType, creditNoteViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invoice\".\"credit_note_view\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invoice\".\"credit_note_view\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "invoice: unable to insert into credit_note_view")
	}

	if !cached {
		creditNoteViewInsertCacheMut.Lock()
		creditNoteViewInsertCache[key] = cache
		creditNoteViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CreditNoteView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("invoice: no credit_note_view provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(creditNoteViewColumnsWithDefault, o)

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

	creditNoteViewUpsertCacheMut.RLock()
	cache, cached := creditNoteViewUpsertCache[key]
	creditNoteViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			creditNoteViewAllColumns,
			creditNoteViewColumnsWithDefault,
			creditNoteViewColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			creditNoteViewAllColumns,
			creditNoteViewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("invoice: unable to upsert credit_note_view, could not build update column list")
		}

		ret := strmangle.SetComplement(creditNoteViewAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(creditNoteViewPrimaryKeyColumns) == 0 {
				return errors.New("invoice: unable to upsert credit_note_view, could not build conflict column list")
			}

			conflict = make([]string, len(creditNoteViewPrimaryKeyColumns))
			copy(conflict, creditNoteViewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invoice\".\"credit_note_view\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(creditNoteViewType, creditNoteViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(creditNoteViewType, creditNoteViewMapping, ret)
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
		return errors.Wrap(err, "invoice: unable to upsert credit_note_view")
	}

	if !cached {
		creditNoteViewUpsertCacheMut.Lock()
		creditNoteViewUpsertCache[key] = cache
		creditNoteViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
