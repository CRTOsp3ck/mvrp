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

// SalesQuotation is an object representing the database table.
type SalesQuotation struct {
	ID                   int                  `boil:"id" json:"id" toml:"id" yaml:"id"`
	BaseDocumentID       int                  `boil:"base_document_id" json:"base_document_id" toml:"base_document_id" yaml:"base_document_id"`
	SalesQuotationNumber string               `boil:"sales_quotation_number" json:"sales_quotation_number" toml:"sales_quotation_number" yaml:"sales_quotation_number"`
	ValidUntilDate       null.Time            `boil:"valid_until_date" json:"valid_until_date,omitempty" toml:"valid_until_date" yaml:"valid_until_date,omitempty"`
	VendorID             null.Int             `boil:"vendor_id" json:"vendor_id,omitempty" toml:"vendor_id" yaml:"vendor_id,omitempty"`
	CustomerID           null.Int             `boil:"customer_id" json:"customer_id,omitempty" toml:"customer_id" yaml:"customer_id,omitempty"`
	ShipToInformation    null.String          `boil:"ship_to_information" json:"ship_to_information,omitempty" toml:"ship_to_information" yaml:"ship_to_information,omitempty"`
	RequestedBy          null.String          `boil:"requested_by" json:"requested_by,omitempty" toml:"requested_by" yaml:"requested_by,omitempty"`
	PreparedBy           null.String          `boil:"prepared_by" json:"prepared_by,omitempty" toml:"prepared_by" yaml:"prepared_by,omitempty"`
	QuotationStatus      SalesQuotationStatus `boil:"quotation_status" json:"quotation_status" toml:"quotation_status" yaml:"quotation_status"`

	R *salesQuotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L salesQuotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SalesQuotationColumns = struct {
	ID                   string
	BaseDocumentID       string
	SalesQuotationNumber string
	ValidUntilDate       string
	VendorID             string
	CustomerID           string
	ShipToInformation    string
	RequestedBy          string
	PreparedBy           string
	QuotationStatus      string
}{
	ID:                   "id",
	BaseDocumentID:       "base_document_id",
	SalesQuotationNumber: "sales_quotation_number",
	ValidUntilDate:       "valid_until_date",
	VendorID:             "vendor_id",
	CustomerID:           "customer_id",
	ShipToInformation:    "ship_to_information",
	RequestedBy:          "requested_by",
	PreparedBy:           "prepared_by",
	QuotationStatus:      "quotation_status",
}

var SalesQuotationTableColumns = struct {
	ID                   string
	BaseDocumentID       string
	SalesQuotationNumber string
	ValidUntilDate       string
	VendorID             string
	CustomerID           string
	ShipToInformation    string
	RequestedBy          string
	PreparedBy           string
	QuotationStatus      string
}{
	ID:                   "sales_quotation.id",
	BaseDocumentID:       "sales_quotation.base_document_id",
	SalesQuotationNumber: "sales_quotation.sales_quotation_number",
	ValidUntilDate:       "sales_quotation.valid_until_date",
	VendorID:             "sales_quotation.vendor_id",
	CustomerID:           "sales_quotation.customer_id",
	ShipToInformation:    "sales_quotation.ship_to_information",
	RequestedBy:          "sales_quotation.requested_by",
	PreparedBy:           "sales_quotation.prepared_by",
	QuotationStatus:      "sales_quotation.quotation_status",
}

// Generated where

type whereHelperSalesQuotationStatus struct{ field string }

func (w whereHelperSalesQuotationStatus) EQ(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperSalesQuotationStatus) NEQ(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperSalesQuotationStatus) LT(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperSalesQuotationStatus) LTE(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperSalesQuotationStatus) GT(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperSalesQuotationStatus) GTE(x SalesQuotationStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperSalesQuotationStatus) IN(slice []SalesQuotationStatus) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperSalesQuotationStatus) NIN(slice []SalesQuotationStatus) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var SalesQuotationWhere = struct {
	ID                   whereHelperint
	BaseDocumentID       whereHelperint
	SalesQuotationNumber whereHelperstring
	ValidUntilDate       whereHelpernull_Time
	VendorID             whereHelpernull_Int
	CustomerID           whereHelpernull_Int
	ShipToInformation    whereHelpernull_String
	RequestedBy          whereHelpernull_String
	PreparedBy           whereHelpernull_String
	QuotationStatus      whereHelperSalesQuotationStatus
}{
	ID:                   whereHelperint{field: "\"sale\".\"sales_quotation\".\"id\""},
	BaseDocumentID:       whereHelperint{field: "\"sale\".\"sales_quotation\".\"base_document_id\""},
	SalesQuotationNumber: whereHelperstring{field: "\"sale\".\"sales_quotation\".\"sales_quotation_number\""},
	ValidUntilDate:       whereHelpernull_Time{field: "\"sale\".\"sales_quotation\".\"valid_until_date\""},
	VendorID:             whereHelpernull_Int{field: "\"sale\".\"sales_quotation\".\"vendor_id\""},
	CustomerID:           whereHelpernull_Int{field: "\"sale\".\"sales_quotation\".\"customer_id\""},
	ShipToInformation:    whereHelpernull_String{field: "\"sale\".\"sales_quotation\".\"ship_to_information\""},
	RequestedBy:          whereHelpernull_String{field: "\"sale\".\"sales_quotation\".\"requested_by\""},
	PreparedBy:           whereHelpernull_String{field: "\"sale\".\"sales_quotation\".\"prepared_by\""},
	QuotationStatus:      whereHelperSalesQuotationStatus{field: "\"sale\".\"sales_quotation\".\"quotation_status\""},
}

// SalesQuotationRels is where relationship names are stored.
var SalesQuotationRels = struct {
	SalesQuotationItems string
}{
	SalesQuotationItems: "SalesQuotationItems",
}

// salesQuotationR is where relationships are stored.
type salesQuotationR struct {
	SalesQuotationItems SalesQuotationItemSlice `boil:"SalesQuotationItems" json:"SalesQuotationItems" toml:"SalesQuotationItems" yaml:"SalesQuotationItems"`
}

// NewStruct creates a new relationship struct
func (*salesQuotationR) NewStruct() *salesQuotationR {
	return &salesQuotationR{}
}

func (r *salesQuotationR) GetSalesQuotationItems() SalesQuotationItemSlice {
	if r == nil {
		return nil
	}
	return r.SalesQuotationItems
}

// salesQuotationL is where Load methods for each relationship are stored.
type salesQuotationL struct{}

var (
	salesQuotationAllColumns            = []string{"id", "base_document_id", "sales_quotation_number", "valid_until_date", "vendor_id", "customer_id", "ship_to_information", "requested_by", "prepared_by", "quotation_status"}
	salesQuotationColumnsWithoutDefault = []string{"id", "base_document_id", "sales_quotation_number", "quotation_status"}
	salesQuotationColumnsWithDefault    = []string{"valid_until_date", "vendor_id", "customer_id", "ship_to_information", "requested_by", "prepared_by"}
	salesQuotationPrimaryKeyColumns     = []string{"id"}
	salesQuotationGeneratedColumns      = []string{}
)

type (
	// SalesQuotationSlice is an alias for a slice of pointers to SalesQuotation.
	// This should almost always be used instead of []SalesQuotation.
	SalesQuotationSlice []*SalesQuotation
	// SalesQuotationHook is the signature for custom SalesQuotation hook methods
	SalesQuotationHook func(context.Context, boil.ContextExecutor, *SalesQuotation) error

	salesQuotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	salesQuotationType                 = reflect.TypeOf(&SalesQuotation{})
	salesQuotationMapping              = queries.MakeStructMapping(salesQuotationType)
	salesQuotationPrimaryKeyMapping, _ = queries.BindMapping(salesQuotationType, salesQuotationMapping, salesQuotationPrimaryKeyColumns)
	salesQuotationInsertCacheMut       sync.RWMutex
	salesQuotationInsertCache          = make(map[string]insertCache)
	salesQuotationUpdateCacheMut       sync.RWMutex
	salesQuotationUpdateCache          = make(map[string]updateCache)
	salesQuotationUpsertCacheMut       sync.RWMutex
	salesQuotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var salesQuotationAfterSelectMu sync.Mutex
var salesQuotationAfterSelectHooks []SalesQuotationHook

var salesQuotationBeforeInsertMu sync.Mutex
var salesQuotationBeforeInsertHooks []SalesQuotationHook
var salesQuotationAfterInsertMu sync.Mutex
var salesQuotationAfterInsertHooks []SalesQuotationHook

var salesQuotationBeforeUpdateMu sync.Mutex
var salesQuotationBeforeUpdateHooks []SalesQuotationHook
var salesQuotationAfterUpdateMu sync.Mutex
var salesQuotationAfterUpdateHooks []SalesQuotationHook

var salesQuotationBeforeDeleteMu sync.Mutex
var salesQuotationBeforeDeleteHooks []SalesQuotationHook
var salesQuotationAfterDeleteMu sync.Mutex
var salesQuotationAfterDeleteHooks []SalesQuotationHook

var salesQuotationBeforeUpsertMu sync.Mutex
var salesQuotationBeforeUpsertHooks []SalesQuotationHook
var salesQuotationAfterUpsertMu sync.Mutex
var salesQuotationAfterUpsertHooks []SalesQuotationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SalesQuotation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SalesQuotation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SalesQuotation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SalesQuotation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SalesQuotation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SalesQuotation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SalesQuotation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SalesQuotation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SalesQuotation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range salesQuotationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSalesQuotationHook registers your hook function for all future operations.
func AddSalesQuotationHook(hookPoint boil.HookPoint, salesQuotationHook SalesQuotationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		salesQuotationAfterSelectMu.Lock()
		salesQuotationAfterSelectHooks = append(salesQuotationAfterSelectHooks, salesQuotationHook)
		salesQuotationAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		salesQuotationBeforeInsertMu.Lock()
		salesQuotationBeforeInsertHooks = append(salesQuotationBeforeInsertHooks, salesQuotationHook)
		salesQuotationBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		salesQuotationAfterInsertMu.Lock()
		salesQuotationAfterInsertHooks = append(salesQuotationAfterInsertHooks, salesQuotationHook)
		salesQuotationAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		salesQuotationBeforeUpdateMu.Lock()
		salesQuotationBeforeUpdateHooks = append(salesQuotationBeforeUpdateHooks, salesQuotationHook)
		salesQuotationBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		salesQuotationAfterUpdateMu.Lock()
		salesQuotationAfterUpdateHooks = append(salesQuotationAfterUpdateHooks, salesQuotationHook)
		salesQuotationAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		salesQuotationBeforeDeleteMu.Lock()
		salesQuotationBeforeDeleteHooks = append(salesQuotationBeforeDeleteHooks, salesQuotationHook)
		salesQuotationBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		salesQuotationAfterDeleteMu.Lock()
		salesQuotationAfterDeleteHooks = append(salesQuotationAfterDeleteHooks, salesQuotationHook)
		salesQuotationAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		salesQuotationBeforeUpsertMu.Lock()
		salesQuotationBeforeUpsertHooks = append(salesQuotationBeforeUpsertHooks, salesQuotationHook)
		salesQuotationBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		salesQuotationAfterUpsertMu.Lock()
		salesQuotationAfterUpsertHooks = append(salesQuotationAfterUpsertHooks, salesQuotationHook)
		salesQuotationAfterUpsertMu.Unlock()
	}
}

// One returns a single salesQuotation record from the query.
func (q salesQuotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SalesQuotation, error) {
	o := &SalesQuotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: failed to execute a one query for sales_quotation")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SalesQuotation records from the query.
func (q salesQuotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (SalesQuotationSlice, error) {
	var o []*SalesQuotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sale: failed to assign all query results to SalesQuotation slice")
	}

	if len(salesQuotationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SalesQuotation records in the query.
func (q salesQuotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to count sales_quotation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q salesQuotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sale: failed to check if sales_quotation exists")
	}

	return count > 0, nil
}

// SalesQuotationItems retrieves all the sales_quotation_item's SalesQuotationItems with an executor.
func (o *SalesQuotation) SalesQuotationItems(mods ...qm.QueryMod) salesQuotationItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sale\".\"sales_quotation_item\".\"sales_quotation_id\"=?", o.ID),
	)

	return SalesQuotationItems(queryMods...)
}

// LoadSalesQuotationItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (salesQuotationL) LoadSalesQuotationItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSalesQuotation interface{}, mods queries.Applicator) error {
	var slice []*SalesQuotation
	var object *SalesQuotation

	if singular {
		var ok bool
		object, ok = maybeSalesQuotation.(*SalesQuotation)
		if !ok {
			object = new(SalesQuotation)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSalesQuotation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSalesQuotation))
			}
		}
	} else {
		s, ok := maybeSalesQuotation.(*[]*SalesQuotation)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSalesQuotation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSalesQuotation))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &salesQuotationR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &salesQuotationR{}
			}
			args[obj.ID] = struct{}{}
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
		qm.From(`sale.sales_quotation_item`),
		qm.WhereIn(`sale.sales_quotation_item.sales_quotation_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sales_quotation_item")
	}

	var resultSlice []*SalesQuotationItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sales_quotation_item")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sales_quotation_item")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sales_quotation_item")
	}

	if len(salesQuotationItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SalesQuotationItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &salesQuotationItemR{}
			}
			foreign.R.SalesQuotation = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SalesQuotationID {
				local.R.SalesQuotationItems = append(local.R.SalesQuotationItems, foreign)
				if foreign.R == nil {
					foreign.R = &salesQuotationItemR{}
				}
				foreign.R.SalesQuotation = local
				break
			}
		}
	}

	return nil
}

// AddSalesQuotationItems adds the given related objects to the existing relationships
// of the sales_quotation, optionally inserting them as new records.
// Appends related to o.R.SalesQuotationItems.
// Sets related.R.SalesQuotation appropriately.
func (o *SalesQuotation) AddSalesQuotationItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SalesQuotationItem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SalesQuotationID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sale\".\"sales_quotation_item\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"sales_quotation_id"}),
				strmangle.WhereClause("\"", "\"", 2, salesQuotationItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SalesQuotationID = o.ID
		}
	}

	if o.R == nil {
		o.R = &salesQuotationR{
			SalesQuotationItems: related,
		}
	} else {
		o.R.SalesQuotationItems = append(o.R.SalesQuotationItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &salesQuotationItemR{
				SalesQuotation: o,
			}
		} else {
			rel.R.SalesQuotation = o
		}
	}
	return nil
}

// SalesQuotations retrieves all the records using an executor.
func SalesQuotations(mods ...qm.QueryMod) salesQuotationQuery {
	mods = append(mods, qm.From("\"sale\".\"sales_quotation\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"sale\".\"sales_quotation\".*"})
	}

	return salesQuotationQuery{q}
}

// FindSalesQuotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSalesQuotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*SalesQuotation, error) {
	salesQuotationObj := &SalesQuotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"sale\".\"sales_quotation\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, salesQuotationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sale: unable to select from sales_quotation")
	}

	if err = salesQuotationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return salesQuotationObj, err
	}

	return salesQuotationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SalesQuotation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sale: no sales_quotation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(salesQuotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	salesQuotationInsertCacheMut.RLock()
	cache, cached := salesQuotationInsertCache[key]
	salesQuotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			salesQuotationAllColumns,
			salesQuotationColumnsWithDefault,
			salesQuotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(salesQuotationType, salesQuotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(salesQuotationType, salesQuotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"sale\".\"sales_quotation\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"sale\".\"sales_quotation\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sale: unable to insert into sales_quotation")
	}

	if !cached {
		salesQuotationInsertCacheMut.Lock()
		salesQuotationInsertCache[key] = cache
		salesQuotationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SalesQuotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SalesQuotation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	salesQuotationUpdateCacheMut.RLock()
	cache, cached := salesQuotationUpdateCache[key]
	salesQuotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			salesQuotationAllColumns,
			salesQuotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sale: unable to update sales_quotation, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"sale\".\"sales_quotation\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, salesQuotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(salesQuotationType, salesQuotationMapping, append(wl, salesQuotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "sale: unable to update sales_quotation row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by update for sales_quotation")
	}

	if !cached {
		salesQuotationUpdateCacheMut.Lock()
		salesQuotationUpdateCache[key] = cache
		salesQuotationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q salesQuotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to update all for sales_quotation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to retrieve rows affected for sales_quotation")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SalesQuotationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesQuotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"sale\".\"sales_quotation\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, salesQuotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to update all in salesQuotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to retrieve rows affected all in update all salesQuotation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SalesQuotation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("sale: no sales_quotation provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(salesQuotationColumnsWithDefault, o)

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

	salesQuotationUpsertCacheMut.RLock()
	cache, cached := salesQuotationUpsertCache[key]
	salesQuotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			salesQuotationAllColumns,
			salesQuotationColumnsWithDefault,
			salesQuotationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			salesQuotationAllColumns,
			salesQuotationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sale: unable to upsert sales_quotation, could not build update column list")
		}

		ret := strmangle.SetComplement(salesQuotationAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(salesQuotationPrimaryKeyColumns) == 0 {
				return errors.New("sale: unable to upsert sales_quotation, could not build conflict column list")
			}

			conflict = make([]string, len(salesQuotationPrimaryKeyColumns))
			copy(conflict, salesQuotationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"sale\".\"sales_quotation\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(salesQuotationType, salesQuotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(salesQuotationType, salesQuotationMapping, ret)
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
		return errors.Wrap(err, "sale: unable to upsert sales_quotation")
	}

	if !cached {
		salesQuotationUpsertCacheMut.Lock()
		salesQuotationUpsertCache[key] = cache
		salesQuotationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SalesQuotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SalesQuotation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sale: no SalesQuotation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), salesQuotationPrimaryKeyMapping)
	sql := "DELETE FROM \"sale\".\"sales_quotation\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete from sales_quotation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by delete for sales_quotation")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q salesQuotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sale: no salesQuotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete all from sales_quotation")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by deleteall for sales_quotation")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SalesQuotationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(salesQuotationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesQuotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"sale\".\"sales_quotation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesQuotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sale: unable to delete all from salesQuotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sale: failed to get rows affected by deleteall for sales_quotation")
	}

	if len(salesQuotationAfterDeleteHooks) != 0 {
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
func (o *SalesQuotation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSalesQuotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SalesQuotationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SalesQuotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesQuotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"sale\".\"sales_quotation\".* FROM \"sale\".\"sales_quotation\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesQuotationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sale: unable to reload all in SalesQuotationSlice")
	}

	*o = slice

	return nil
}

// SalesQuotationExists checks if the SalesQuotation row exists.
func SalesQuotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"sale\".\"sales_quotation\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sale: unable to check if sales_quotation exists")
	}

	return exists, nil
}

// Exists checks if the SalesQuotation row exists.
func (o *SalesQuotation) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SalesQuotationExists(ctx, exec, o.ID)
}
