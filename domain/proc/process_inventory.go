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

func ProcessGoodsIssueNoteAmounts(gin *inventory.GoodsIssueNote, ginItems []*inventory.GoodsIssueNoteItem) error {
	totalValue := 0.0

	for _, ginItem := range ginItems {
		quantity, ok := ginItem.Quantity.Float64()
		if !ok {
			return errors.New("invalid quantity found while processing goods issue note amounts")
		}
		unitValue, ok := ginItem.UnitValue.Float64()
		if !ok {
			return errors.New("invalid unit value found while processing goods issue note amounts")
		}

		// update goods issue note item total value
		ginItem.TotalValueGen = types.NewDecimal(decimal.New(int64(quantity*unitValue*100), 2))

		// update goods issue note total value
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	gin.TotalValue = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessReturnMerchandiseAuthorizationAmounts(rma *inventory.ReturnMerchandiseAuthorization, rmaItems []*inventory.ReturnMerchandiseAuthorizationItem) error {
	totalValue := 0.0

	for _, rmaItem := range rmaItems {
		quantity, ok := rmaItem.Quantity.Float64()
		if !ok {
			return errors.New("invalid quantity found while processing goods issue note amounts")
		}
		unitValue, ok := rmaItem.UnitValue.Float64()
		if !ok {
			return errors.New("invalid unit value found while processing goods issue note amounts")
		}
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	rma.TotalValue = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessStockCountSheetAmounts(scs *inventory.StockCountSheet, sysInv *inventory.Inventory) error {
	discrepencyValue := 0.0
	systemInventoryCount, ok := sysInv.QuantityAvailable.Float64()
	if !ok {
		return errors.New("invalid system inventory count found while processing stock count sheet amounts")
	}
	scsQuantity, ok := scs.TotalQuantity.Float64()
	if !ok {
		return errors.New("invalid stock count sheet quantity found while processing stock count sheet amounts")
	}
	discrepencyValue = scsQuantity - systemInventoryCount

	// update stock count sheet amounts
	scs.Discrepancies = types.NewNullDecimal(decimal.New(int64(discrepencyValue*100), 2))

	return nil
}
