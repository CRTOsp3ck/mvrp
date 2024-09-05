// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package item

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
var ErrSyncFail = errors.New("item: failed to synchronize data after insert")

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

type ItemType string

// Enum values for ItemType
const (
	ItemTypeProduct ItemType = "product"
	ItemTypeService ItemType = "service"
	ItemTypeOther   ItemType = "other"
)

func AllItemType() []ItemType {
	return []ItemType{
		ItemTypeProduct,
		ItemTypeService,
		ItemTypeOther,
	}
}

func (e ItemType) IsValid() error {
	switch e {
	case ItemTypeProduct, ItemTypeService, ItemTypeOther:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ItemType) String() string {
	return string(e)
}

func (e ItemType) Ordinal() int {
	switch e {
	case ItemTypeProduct:
		return 0
	case ItemTypeService:
		return 1
	case ItemTypeOther:
		return 2

	default:
		panic(errors.New("enum is not valid"))
	}
}

type ItemStatus string

// Enum values for ItemStatus
const (
	ItemStatusActive       ItemStatus = "active"
	ItemStatusInactive     ItemStatus = "inactive"
	ItemStatusDiscontinued ItemStatus = "discontinued"
	ItemStatusTesting      ItemStatus = "testing"
)

func AllItemStatus() []ItemStatus {
	return []ItemStatus{
		ItemStatusActive,
		ItemStatusInactive,
		ItemStatusDiscontinued,
		ItemStatusTesting,
	}
}

func (e ItemStatus) IsValid() error {
	switch e {
	case ItemStatusActive, ItemStatusInactive, ItemStatusDiscontinued, ItemStatusTesting:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e ItemStatus) String() string {
	return string(e)
}

func (e ItemStatus) Ordinal() int {
	switch e {
	case ItemStatusActive:
		return 0
	case ItemStatusInactive:
		return 1
	case ItemStatusDiscontinued:
		return 2
	case ItemStatusTesting:
		return 3

	default:
		panic(errors.New("enum is not valid"))
	}
}
