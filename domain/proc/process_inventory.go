package proc

import (
	"errors"
	"mvrp/data/model/inventory"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

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
