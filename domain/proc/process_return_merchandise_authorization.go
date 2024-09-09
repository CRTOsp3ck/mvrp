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
			return errors.New("invalid quantity found while processing return merchandise authorization amounts")
		}
		unitValue, ok := rmaItem.UnitValue.Float64()
		if !ok {
			return errors.New("invalid unit value found while processing return merchandise authorization amounts")
		}
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	rma.TotalValueGen = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessReturnMerchandiseAuthorizationItemAmounts(rmaItem *inventory.ReturnMerchandiseAuthorizationItem) error {
	quantity, ok := rmaItem.Quantity.Float64()
	if !ok {
		return errors.New("invalid quantity found while processing return merchandise authorization item amounts")
	}
	unitValue, ok := rmaItem.UnitValue.Float64()
	if !ok {
		return errors.New("invalid unit value found while processing return merchandise authorization item amounts")
	}

	// update goods issue note item total value
	rmaItem.TotalValueGen = types.NewDecimal(decimal.New(int64(quantity*unitValue*100), 2))

	return nil
}
