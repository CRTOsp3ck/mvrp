package proc

import (
	"mvrp/data/model/invoice"
	"mvrp/errors"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func ProcessPaymentReceiptAmounts(pr *invoice.PaymentReceipt, prItems []*invoice.PaymentReceiptItem) error {
	totalValue := 0.0

	for _, prItem := range prItems {
		quantity, ok := prItem.Quantity.Float64()
		if !ok {
			return errors.New("invalid quantity found while processing goods issue note amounts")
		}
		unitValue, ok := prItem.UnitValue.Float64()
		if !ok {
			return errors.New("invalid unit value found while processing goods issue note amounts")
		}

		// update goods issue note total value
		totalValue += quantity * unitValue
	}

	// update goods issue note amounts
	pr.TotalValueGen = types.NewDecimal(decimal.New(int64(totalValue*100), 2))

	return nil
}

func ProcessPaymentReceiptItemAmounts(prItem *invoice.PaymentReceiptItem) error {
	quantity, ok := prItem.Quantity.Float64()
	if !ok {
		return errors.New("invalid quantity found while processing goods issue note item amounts")
	}
	unitValue, ok := prItem.UnitValue.Float64()
	if !ok {
		return errors.New("invalid unit value found while processing goods issue note item amounts")
	}

	// update goods issue note item total value
	prItem.TotalValueGen = types.NewDecimal(decimal.New(int64(quantity*unitValue*100), 2))

	return nil
}
