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

// PaymentReceiptView is an object representing the database table.
type PaymentReceiptView struct {
	ID                   null.Int          `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	PaymentReceiptNumber null.String       `boil:"payment_receipt_number" json:"payment_receipt_number,omitempty" toml:"payment_receipt_number" yaml:"payment_receipt_number,omitempty"`
	InvoiceID            null.Int          `boil:"invoice_id" json:"invoice_id,omitempty" toml:"invoice_id" yaml:"invoice_id,omitempty"`
	DateOfPayment        null.Time         `boil:"date_of_payment" json:"date_of_payment,omitempty" toml:"date_of_payment" yaml:"date_of_payment,omitempty"`
	PayerID              null.Int          `boil:"payer_id" json:"payer_id,omitempty" toml:"payer_id" yaml:"payer_id,omitempty"`
	PayeeID              null.Int          `boil:"payee_id" json:"payee_id,omitempty" toml:"payee_id" yaml:"payee_id,omitempty"`
	TotalValueGen        types.NullDecimal `boil:"total_value_gen" json:"total_value_gen,omitempty" toml:"total_value_gen" yaml:"total_value_gen,omitempty"`
	CreatedAt            null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt            null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt            null.Time         `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	PayerInfo            null.JSON         `boil:"payer_info" json:"payer_info,omitempty" toml:"payer_info" yaml:"payer_info,omitempty"`
	PayeeInfo            null.JSON         `boil:"payee_info" json:"payee_info,omitempty" toml:"payee_info" yaml:"payee_info,omitempty"`
	PaymentReceiptItems  null.JSON         `boil:"payment_receipt_items" json:"payment_receipt_items,omitempty" toml:"payment_receipt_items" yaml:"payment_receipt_items,omitempty"`
}

var PaymentReceiptViewColumns = struct {
	ID                   string
	PaymentReceiptNumber string
	InvoiceID            string
	DateOfPayment        string
	PayerID              string
	PayeeID              string
	TotalValueGen        string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	PayerInfo            string
	PayeeInfo            string
	PaymentReceiptItems  string
}{
	ID:                   "id",
	PaymentReceiptNumber: "payment_receipt_number",
	InvoiceID:            "invoice_id",
	DateOfPayment:        "date_of_payment",
	PayerID:              "payer_id",
	PayeeID:              "payee_id",
	TotalValueGen:        "total_value_gen",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	DeletedAt:            "deleted_at",
	PayerInfo:            "payer_info",
	PayeeInfo:            "payee_info",
	PaymentReceiptItems:  "payment_receipt_items",
}

var PaymentReceiptViewTableColumns = struct {
	ID                   string
	PaymentReceiptNumber string
	InvoiceID            string
	DateOfPayment        string
	PayerID              string
	PayeeID              string
	TotalValueGen        string
	CreatedAt            string
	UpdatedAt            string
	DeletedAt            string
	PayerInfo            string
	PayeeInfo            string
	PaymentReceiptItems  string
}{
	ID:                   "payment_receipt_view.id",
	PaymentReceiptNumber: "payment_receipt_view.payment_receipt_number",
	InvoiceID:            "payment_receipt_view.invoice_id",
	DateOfPayment:        "payment_receipt_view.date_of_payment",
	PayerID:              "payment_receipt_view.payer_id",
	PayeeID:              "payment_receipt_view.payee_id",
	TotalValueGen:        "payment_receipt_view.total_value_gen",
	CreatedAt:            "payment_receipt_view.created_at",
	UpdatedAt:            "payment_receipt_view.updated_at",
	DeletedAt:            "payment_receipt_view.deleted_at",
	PayerInfo:            "payment_receipt_view.payer_info",
	PayeeInfo:            "payment_receipt_view.payee_info",
	PaymentReceiptItems:  "payment_receipt_view.payment_receipt_items",
}

// Generated where

var PaymentReceiptViewWhere = struct {
	ID                   whereHelpernull_Int
	PaymentReceiptNumber whereHelpernull_String
	InvoiceID            whereHelpernull_Int
	DateOfPayment        whereHelpernull_Time
	PayerID              whereHelpernull_Int
	PayeeID              whereHelpernull_Int
	TotalValueGen        whereHelpertypes_NullDecimal
	CreatedAt            whereHelpernull_Time
	UpdatedAt            whereHelpernull_Time
	DeletedAt            whereHelpernull_Time
	PayerInfo            whereHelpernull_JSON
	PayeeInfo            whereHelpernull_JSON
	PaymentReceiptItems  whereHelpernull_JSON
}{
	ID:                   whereHelpernull_Int{field: "\"invoice\".\"payment_receipt_view\".\"id\""},
	PaymentReceiptNumber: whereHelpernull_String{field: "\"invoice\".\"payment_receipt_view\".\"payment_receipt_number\""},
	InvoiceID:            whereHelpernull_Int{field: "\"invoice\".\"payment_receipt_view\".\"invoice_id\""},
	DateOfPayment:        whereHelpernull_Time{field: "\"invoice\".\"payment_receipt_view\".\"date_of_payment\""},
	PayerID:              whereHelpernull_Int{field: "\"invoice\".\"payment_receipt_view\".\"payer_id\""},
	PayeeID:              whereHelpernull_Int{field: "\"invoice\".\"payment_receipt_view\".\"payee_id\""},
	TotalValueGen:        whereHelpertypes_NullDecimal{field: "\"invoice\".\"payment_receipt_view\".\"total_value_gen\""},
	CreatedAt:            whereHelpernull_Time{field: "\"invoice\".\"payment_receipt_view\".\"created_at\""},
	UpdatedAt:            whereHelpernull_Time{field: "\"invoice\".\"payment_receipt_view\".\"updated_at\""},
	DeletedAt:            whereHelpernull_Time{field: "\"invoice\".\"payment_receipt_view\".\"deleted_at\""},
	PayerInfo:            whereHelpernull_JSON{field: "\"invoice\".\"payment_receipt_view\".\"payer_info\""},
	PayeeInfo:            whereHelpernull_JSON{field: "\"invoice\".\"payment_receipt_view\".\"payee_info\""},
	PaymentReceiptItems:  whereHelpernull_JSON{field: "\"invoice\".\"payment_receipt_view\".\"payment_receipt_items\""},
}

var (
	paymentReceiptViewAllColumns            = []string{"id", "payment_receipt_number", "invoice_id", "date_of_payment", "payer_id", "payee_id", "total_value_gen", "created_at", "updated_at", "deleted_at", "payer_info", "payee_info", "payment_receipt_items"}
	paymentReceiptViewColumnsWithoutDefault = []string{}
	paymentReceiptViewColumnsWithDefault    = []string{"id", "payment_receipt_number", "invoice_id", "date_of_payment", "payer_id", "payee_id", "total_value_gen", "created_at", "updated_at", "deleted_at", "payer_info", "payee_info", "payment_receipt_items"}
	paymentReceiptViewPrimaryKeyColumns     = []string{}
	paymentReceiptViewGeneratedColumns      = []string{}
)

type (
	// PaymentReceiptViewSlice is an alias for a slice of pointers to PaymentReceiptView.
	// This should almost always be used instead of []PaymentReceiptView.
	PaymentReceiptViewSlice []*PaymentReceiptView
	// PaymentReceiptViewHook is the signature for custom PaymentReceiptView hook methods
	PaymentReceiptViewHook func(context.Context, boil.ContextExecutor, *PaymentReceiptView) error

	paymentReceiptViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	paymentReceiptViewType           = reflect.TypeOf(&PaymentReceiptView{})
	paymentReceiptViewMapping        = queries.MakeStructMapping(paymentReceiptViewType)
	paymentReceiptViewInsertCacheMut sync.RWMutex
	paymentReceiptViewInsertCache    = make(map[string]insertCache)
	paymentReceiptViewUpdateCacheMut sync.RWMutex
	paymentReceiptViewUpdateCache    = make(map[string]updateCache)
	paymentReceiptViewUpsertCacheMut sync.RWMutex
	paymentReceiptViewUpsertCache    = make(map[string]insertCache)
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

var paymentReceiptViewAfterSelectMu sync.Mutex
var paymentReceiptViewAfterSelectHooks []PaymentReceiptViewHook

var paymentReceiptViewBeforeInsertMu sync.Mutex
var paymentReceiptViewBeforeInsertHooks []PaymentReceiptViewHook
var paymentReceiptViewAfterInsertMu sync.Mutex
var paymentReceiptViewAfterInsertHooks []PaymentReceiptViewHook

var paymentReceiptViewBeforeUpsertMu sync.Mutex
var paymentReceiptViewBeforeUpsertHooks []PaymentReceiptViewHook
var paymentReceiptViewAfterUpsertMu sync.Mutex
var paymentReceiptViewAfterUpsertHooks []PaymentReceiptViewHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PaymentReceiptView) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentReceiptViewAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PaymentReceiptView) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentReceiptViewBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PaymentReceiptView) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentReceiptViewAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PaymentReceiptView) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentReceiptViewBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PaymentReceiptView) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentReceiptViewAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPaymentReceiptViewHook registers your hook function for all future operations.
func AddPaymentReceiptViewHook(hookPoint boil.HookPoint, paymentReceiptViewHook PaymentReceiptViewHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		paymentReceiptViewAfterSelectMu.Lock()
		paymentReceiptViewAfterSelectHooks = append(paymentReceiptViewAfterSelectHooks, paymentReceiptViewHook)
		paymentReceiptViewAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		paymentReceiptViewBeforeInsertMu.Lock()
		paymentReceiptViewBeforeInsertHooks = append(paymentReceiptViewBeforeInsertHooks, paymentReceiptViewHook)
		paymentReceiptViewBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		paymentReceiptViewAfterInsertMu.Lock()
		paymentReceiptViewAfterInsertHooks = append(paymentReceiptViewAfterInsertHooks, paymentReceiptViewHook)
		paymentReceiptViewAfterInsertMu.Unlock()
	case boil.BeforeUpsertHook:
		paymentReceiptViewBeforeUpsertMu.Lock()
		paymentReceiptViewBeforeUpsertHooks = append(paymentReceiptViewBeforeUpsertHooks, paymentReceiptViewHook)
		paymentReceiptViewBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		paymentReceiptViewAfterUpsertMu.Lock()
		paymentReceiptViewAfterUpsertHooks = append(paymentReceiptViewAfterUpsertHooks, paymentReceiptViewHook)
		paymentReceiptViewAfterUpsertMu.Unlock()
	}
}

// One returns a single paymentReceiptView record from the query.
func (q paymentReceiptViewQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PaymentReceiptView, error) {
	o := &PaymentReceiptView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "invoice: failed to execute a one query for payment_receipt_view")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PaymentReceiptView records from the query.
func (q paymentReceiptViewQuery) All(ctx context.Context, exec boil.ContextExecutor) (PaymentReceiptViewSlice, error) {
	var o []*PaymentReceiptView

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "invoice: failed to assign all query results to PaymentReceiptView slice")
	}

	if len(paymentReceiptViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PaymentReceiptView records in the query.
func (q paymentReceiptViewQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "invoice: failed to count payment_receipt_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q paymentReceiptViewQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "invoice: failed to check if payment_receipt_view exists")
	}

	return count > 0, nil
}

// PaymentReceiptViews retrieves all the records using an executor.
func PaymentReceiptViews(mods ...qm.QueryMod) paymentReceiptViewQuery {
	mods = append(mods, qm.From("\"invoice\".\"payment_receipt_view\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"invoice\".\"payment_receipt_view\".*"})
	}

	return paymentReceiptViewQuery{q}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PaymentReceiptView) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("invoice: no payment_receipt_view provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(paymentReceiptViewColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	paymentReceiptViewInsertCacheMut.RLock()
	cache, cached := paymentReceiptViewInsertCache[key]
	paymentReceiptViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			paymentReceiptViewAllColumns,
			paymentReceiptViewColumnsWithDefault,
			paymentReceiptViewColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(paymentReceiptViewType, paymentReceiptViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(paymentReceiptViewType, paymentReceiptViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"invoice\".\"payment_receipt_view\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"invoice\".\"payment_receipt_view\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "invoice: unable to insert into payment_receipt_view")
	}

	if !cached {
		paymentReceiptViewInsertCacheMut.Lock()
		paymentReceiptViewInsertCache[key] = cache
		paymentReceiptViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PaymentReceiptView) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("invoice: no payment_receipt_view provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(paymentReceiptViewColumnsWithDefault, o)

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

	paymentReceiptViewUpsertCacheMut.RLock()
	cache, cached := paymentReceiptViewUpsertCache[key]
	paymentReceiptViewUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			paymentReceiptViewAllColumns,
			paymentReceiptViewColumnsWithDefault,
			paymentReceiptViewColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			paymentReceiptViewAllColumns,
			paymentReceiptViewPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("invoice: unable to upsert payment_receipt_view, could not build update column list")
		}

		ret := strmangle.SetComplement(paymentReceiptViewAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(paymentReceiptViewPrimaryKeyColumns) == 0 {
				return errors.New("invoice: unable to upsert payment_receipt_view, could not build conflict column list")
			}

			conflict = make([]string, len(paymentReceiptViewPrimaryKeyColumns))
			copy(conflict, paymentReceiptViewPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"invoice\".\"payment_receipt_view\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(paymentReceiptViewType, paymentReceiptViewMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(paymentReceiptViewType, paymentReceiptViewMapping, ret)
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
		return errors.Wrap(err, "invoice: unable to upsert payment_receipt_view")
	}

	if !cached {
		paymentReceiptViewUpsertCacheMut.Lock()
		paymentReceiptViewUpsertCache[key] = cache
		paymentReceiptViewUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
