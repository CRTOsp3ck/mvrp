// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package base

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("base: failed to synchronize data after insert")

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

type PaymentTerms string

// Enum values for PaymentTerms
const (
	PaymentTermsFull                          PaymentTerms = "full"
	PaymentTermsPartialBeforeAndAfterDelivery PaymentTerms = "partial_before_and_after_delivery"
	PaymentTermsNet30                         PaymentTerms = "net_30"
	PaymentTermsNet60                         PaymentTerms = "net_60"
	PaymentTermsNet90                         PaymentTerms = "net_90"
)

func AllPaymentTerms() []PaymentTerms {
	return []PaymentTerms{
		PaymentTermsFull,
		PaymentTermsPartialBeforeAndAfterDelivery,
		PaymentTermsNet30,
		PaymentTermsNet60,
		PaymentTermsNet90,
	}
}

func (e PaymentTerms) IsValid() error {
	switch e {
	case PaymentTermsFull, PaymentTermsPartialBeforeAndAfterDelivery, PaymentTermsNet30, PaymentTermsNet60, PaymentTermsNet90:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e PaymentTerms) String() string {
	return string(e)
}

func (e PaymentTerms) Ordinal() int {
	switch e {
	case PaymentTermsFull:
		return 0
	case PaymentTermsPartialBeforeAndAfterDelivery:
		return 1
	case PaymentTermsNet30:
		return 2
	case PaymentTermsNet60:
		return 3
	case PaymentTermsNet90:
		return 4

	default:
		panic(errors.New("enum is not valid"))
	}
}

type PaymentStatus string

// Enum values for PaymentStatus
const (
	PaymentStatusPending       PaymentStatus = "pending"
	PaymentStatusPaid          PaymentStatus = "paid"
	PaymentStatusPartiallyPaid PaymentStatus = "partially_paid"
)

func AllPaymentStatus() []PaymentStatus {
	return []PaymentStatus{
		PaymentStatusPending,
		PaymentStatusPaid,
		PaymentStatusPartiallyPaid,
	}
}

func (e PaymentStatus) IsValid() error {
	switch e {
	case PaymentStatusPending, PaymentStatusPaid, PaymentStatusPartiallyPaid:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e PaymentStatus) String() string {
	return string(e)
}

func (e PaymentStatus) Ordinal() int {
	switch e {
	case PaymentStatusPending:
		return 0
	case PaymentStatusPaid:
		return 1
	case PaymentStatusPartiallyPaid:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}
