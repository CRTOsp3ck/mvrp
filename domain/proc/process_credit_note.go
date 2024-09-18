package proc

import (
	"mvrp/data/model/invoice"
	"mvrp/errors"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func ProcessCreditNoteAmounts(crn *invoice.CreditNote, crnItems []*invoice.CreditNoteItem) error {
	totalValue := 0.0

	for _, crnItem := range crnItems {
		quantity, ok := crnItem.Quantity.Float64()
		if !ok {
			return errors.New("invalid quantity found while processing goods issue note amounts")
		}
		unitValue, ok := crnItem.UnitValue.Float64()
		if !ok {
			return errors.New("invalid unit value found while processing goods issue note amounts")
		}

		// update goods issue note total value
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	crn.TotalValueGen = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessCreditNoteItemAmounts(crnItem *invoice.CreditNoteItem) error {
	quantity, ok := crnItem.Quantity.Float64()
	if !ok {
		return errors.New("invalid quantity found while processing goods issue note item amounts")
	}
	unitValue, ok := crnItem.UnitValue.Float64()
	if !ok {
		return errors.New("invalid unit value found while processing goods issue note item amounts")
	}

	// update goods issue note item total value
	crnItem.TotalValueGen = types.NewDecimal(decimal.New(int64(quantity*unitValue*100), 2))

	return nil
}
