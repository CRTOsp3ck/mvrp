// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sale

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/null/v8/convert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("sale: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type ShippingStatus string

// Enum values for ShippingStatus
const (
	ShippingStatusReadyForPickup ShippingStatus = "ready_for_pickup"
	ShippingStatusInTransit      ShippingStatus = "in_transit"
	ShippingStatusShipped        ShippingStatus = "shipped"
)

func AllShippingStatus() []ShippingStatus {
	return []ShippingStatus{
		ShippingStatusReadyForPickup,
		ShippingStatusInTransit,
		ShippingStatusShipped,
	}
}

func (e ShippingStatus) IsValid() error {
	switch e {
	case ShippingStatusReadyForPickup, ShippingStatusInTransit, ShippingStatusShipped:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ShippingStatus) String() string {
	return string(e)
}

func (e ShippingStatus) Ordinal() int {
	switch e {
	case ShippingStatusReadyForPickup:
		return 0
	case ShippingStatusInTransit:
		return 1
	case ShippingStatusShipped:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}

type SalesOrderStatus string

// Enum values for SalesOrderStatus
const (
	SalesOrderStatusPending  SalesOrderStatus = "pending"
	SalesOrderStatusAccepted SalesOrderStatus = "accepted"
	SalesOrderStatusDeclined SalesOrderStatus = "declined"
)

func AllSalesOrderStatus() []SalesOrderStatus {
	return []SalesOrderStatus{
		SalesOrderStatusPending,
		SalesOrderStatusAccepted,
		SalesOrderStatusDeclined,
	}
}

func (e SalesOrderStatus) IsValid() error {
	switch e {
	case SalesOrderStatusPending, SalesOrderStatusAccepted, SalesOrderStatusDeclined:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e SalesOrderStatus) String() string {
	return string(e)
}

func (e SalesOrderStatus) Ordinal() int {
	switch e {
	case SalesOrderStatusPending:
		return 0
	case SalesOrderStatusAccepted:
		return 1
	case SalesOrderStatusDeclined:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}

type SalesQuotationStatus string

// Enum values for SalesQuotationStatus
const (
	SalesQuotationStatusPending  SalesQuotationStatus = "pending"
	SalesQuotationStatusAccepted SalesQuotationStatus = "accepted"
	SalesQuotationStatusDeclined SalesQuotationStatus = "declined"
)

func AllSalesQuotationStatus() []SalesQuotationStatus {
	return []SalesQuotationStatus{
		SalesQuotationStatusPending,
		SalesQuotationStatusAccepted,
		SalesQuotationStatusDeclined,
	}
}

func (e SalesQuotationStatus) IsValid() error {
	switch e {
	case SalesQuotationStatusPending, SalesQuotationStatusAccepted, SalesQuotationStatusDeclined:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e SalesQuotationStatus) String() string {
	return string(e)
}

func (e SalesQuotationStatus) Ordinal() int {
	switch e {
	case SalesQuotationStatusPending:
		return 0
	case SalesQuotationStatusAccepted:
		return 1
	case SalesQuotationStatusDeclined:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}

// NullShippingStatus is a nullable ShippingStatus enum type. It supports SQL and JSON serialization.
type NullShippingStatus struct {
	Val   ShippingStatus
	Valid bool
}

// NullShippingStatusFrom creates a new ShippingStatus that will never be blank.
func NullShippingStatusFrom(v ShippingStatus) NullShippingStatus {
	return NewNullShippingStatus(v, true)
}

// NullShippingStatusFromPtr creates a new NullShippingStatus that be null if s is nil.
func NullShippingStatusFromPtr(v *ShippingStatus) NullShippingStatus {
	if v == nil {
		return NewNullShippingStatus("", false)
	}
	return NewNullShippingStatus(*v, true)
}

// NewNullShippingStatus creates a new NullShippingStatus
func NewNullShippingStatus(v ShippingStatus, valid bool) NullShippingStatus {
	return NullShippingStatus{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullShippingStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullShippingStatus) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullShippingStatus) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullShippingStatus) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = ShippingStatus(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullShippingStatus value and also sets it to be non-null.
func (e *NullShippingStatus) SetValid(v ShippingStatus) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullShippingStatus value, or a nil pointer if this NullShippingStatus is null.
func (e NullShippingStatus) Ptr() *ShippingStatus {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullShippingStatus) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullShippingStatus) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullShippingStatus) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

// NullSalesOrderStatus is a nullable SalesOrderStatus enum type. It supports SQL and JSON serialization.
type NullSalesOrderStatus struct {
	Val   SalesOrderStatus
	Valid bool
}

// NullSalesOrderStatusFrom creates a new SalesOrderStatus that will never be blank.
func NullSalesOrderStatusFrom(v SalesOrderStatus) NullSalesOrderStatus {
	return NewNullSalesOrderStatus(v, true)
}

// NullSalesOrderStatusFromPtr creates a new NullSalesOrderStatus that be null if s is nil.
func NullSalesOrderStatusFromPtr(v *SalesOrderStatus) NullSalesOrderStatus {
	if v == nil {
		return NewNullSalesOrderStatus("", false)
	}
	return NewNullSalesOrderStatus(*v, true)
}

// NewNullSalesOrderStatus creates a new NullSalesOrderStatus
func NewNullSalesOrderStatus(v SalesOrderStatus, valid bool) NullSalesOrderStatus {
	return NullSalesOrderStatus{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullSalesOrderStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullSalesOrderStatus) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullSalesOrderStatus) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullSalesOrderStatus) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = SalesOrderStatus(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullSalesOrderStatus value and also sets it to be non-null.
func (e *NullSalesOrderStatus) SetValid(v SalesOrderStatus) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullSalesOrderStatus value, or a nil pointer if this NullSalesOrderStatus is null.
func (e NullSalesOrderStatus) Ptr() *SalesOrderStatus {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullSalesOrderStatus) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullSalesOrderStatus) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullSalesOrderStatus) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}

// NullSalesQuotationStatus is a nullable SalesQuotationStatus enum type. It supports SQL and JSON serialization.
type NullSalesQuotationStatus struct {
	Val   SalesQuotationStatus
	Valid bool
}

// NullSalesQuotationStatusFrom creates a new SalesQuotationStatus that will never be blank.
func NullSalesQuotationStatusFrom(v SalesQuotationStatus) NullSalesQuotationStatus {
	return NewNullSalesQuotationStatus(v, true)
}

// NullSalesQuotationStatusFromPtr creates a new NullSalesQuotationStatus that be null if s is nil.
func NullSalesQuotationStatusFromPtr(v *SalesQuotationStatus) NullSalesQuotationStatus {
	if v == nil {
		return NewNullSalesQuotationStatus("", false)
	}
	return NewNullSalesQuotationStatus(*v, true)
}

// NewNullSalesQuotationStatus creates a new NullSalesQuotationStatus
func NewNullSalesQuotationStatus(v SalesQuotationStatus, valid bool) NullSalesQuotationStatus {
	return NullSalesQuotationStatus{
		Val:   v,
		Valid: valid,
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *NullSalesQuotationStatus) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, null.NullBytes) {
		e.Val = ""
		e.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &e.Val); err != nil {
		return err
	}

	e.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
func (e NullSalesQuotationStatus) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return null.NullBytes, nil
	}
	return json.Marshal(e.Val)
}

// MarshalText implements encoding.TextMarshaler.
func (e NullSalesQuotationStatus) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.Val), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *NullSalesQuotationStatus) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		e.Valid = false
		return nil
	}

	e.Val = SalesQuotationStatus(text)
	e.Valid = true
	return nil
}

// SetValid changes this NullSalesQuotationStatus value and also sets it to be non-null.
func (e *NullSalesQuotationStatus) SetValid(v SalesQuotationStatus) {
	e.Val = v
	e.Valid = true
}

// Ptr returns a pointer to this NullSalesQuotationStatus value, or a nil pointer if this NullSalesQuotationStatus is null.
func (e NullSalesQuotationStatus) Ptr() *SalesQuotationStatus {
	if !e.Valid {
		return nil
	}
	return &e.Val
}

// IsZero returns true for null types.
func (e NullSalesQuotationStatus) IsZero() bool {
	return !e.Valid
}

// Scan implements the Scanner interface.
func (e *NullSalesQuotationStatus) Scan(value interface{}) error {
	if value == nil {
		e.Val, e.Valid = "", false
		return nil
	}
	e.Valid = true
	return convert.ConvertAssign((*string)(&e.Val), value)
}

// Value implements the driver Valuer interface.
func (e NullSalesQuotationStatus) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return string(e.Val), nil
}
