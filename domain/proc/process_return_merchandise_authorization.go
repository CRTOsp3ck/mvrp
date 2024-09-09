package proc

import (
	"mvrp/data/model/inventory"
	"mvrp/errors"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

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
