package proc

import (
	"mvrp/data/model/inventory"
	"mvrp/errors"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

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
	scs.DiscrepanciesGen = types.NewNullDecimal(decimal.New(int64(discrepencyValue*100), 2))

	return nil
}
