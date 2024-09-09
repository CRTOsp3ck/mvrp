package proc

import (
	"mvrp/data/model/inventory"
	"mvrp/errors"

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

		// update goods issue note total value
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	gin.TotalValueGen = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessGoodsIssueNoteItemAmounts(ginItem *inventory.GoodsIssueNoteItem) error {
	quantity, ok := ginItem.Quantity.Float64()
	if !ok {
		return errors.New("invalid quantity found while processing goods issue note item amounts")
	}
	unitValue, ok := ginItem.UnitValue.Float64()
	if !ok {
		return errors.New("invalid unit value found while processing goods issue note item amounts")
	}

	// update goods issue note item total value
	ginItem.TotalValueGen = types.NewDecimal(decimal.New(int64(quantity*unitValue*100), 2))

	return nil
}
