package proc

import (
	"mvrp/data/model/inventory"
	"mvrp/errors"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func ProcessInventoryAmounts(inv *inventory.Inventory) error {
	pricePerUnit, ok := inv.PricePerUnit.Float64()
	if !ok {
		return errors.WrapError(errors.ErrTypeInvalidValue, "invalid price per unit found while processing inventory amounts")
	}
	if pricePerUnit <= 0.0 {
		return errors.WrapError(errors.ErrTypeMissingField, "invalid price per unit found while processing inventory amounts")
	}

	// when creating iventory, these values may be undefined
	if inv.QuantityReserved.Big == nil {
		inv.QuantityReserved = types.NewNullDecimal(decimal.New(0, 2))
	}
	if inv.QuantityAvailable.Big == nil {
		inv.QuantityAvailable = types.NewNullDecimal(decimal.New(0, 2))
	}
	if inv.QuantityReturned.Big == nil {
		inv.QuantityReturned = types.NewNullDecimal(decimal.New(0, 2))
	}

	// calculate total value
	qtyReserved, ok := inv.QuantityReserved.Float64()
	if !ok {
		return errors.WrapError(errors.ErrTypeInvalidValue, "invalid quantity reserved found while processing inventory amounts")
	}
	qtyAvailable, ok := inv.QuantityAvailable.Float64()
	if !ok {
		return errors.WrapError(errors.ErrTypeInvalidValue, "invalid quantity available found while processing inventory amounts")
	}
	qtyReturned, ok := inv.QuantityReturned.Float64()
	if !ok {
		return errors.WrapError(errors.ErrTypeInvalidValue, "invalid quantity returned found while processing inventory amounts")
	}
	qtyTotal := qtyReserved + qtyAvailable + qtyReturned
	if qtyTotal < 0.0 {
		return errors.WrapError(errors.ErrTypeInvalidValue, "invalid total quantity found while processing inventory amounts")
	}

	inv.QuantityTotalGen = types.NewNullDecimal(decimal.New(int64(qtyTotal*100), 2))
	inv.TotalValueOnHandGen = types.NewNullDecimal(decimal.New(int64(qtyTotal*pricePerUnit*100), 2))

	return nil
}
