package proc

import (
	"mvrp/data/model/base"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

func ProcessBaseDocumentAmounts(bd *base.BaseDocument, bdis []*base.BaseDocumentItem) error {
	totalItemsAmount := decimal.New(0, 2)
	totalItemsDiscAmount := decimal.New(0, 2)
	totalItemsTaxAmount := decimal.New(0, 2)
	totalItemsShippingFees := decimal.New(0, 2)

	for _, bdi := range bdis {
		totalItemsAmount = totalItemsAmount.Add(totalItemsAmount,
			decimal.New(0, 2).Mul(bdi.UnitPrice.Big, bdi.Quantity.Big),
		)
		totalItemsDiscAmount = totalItemsDiscAmount.Add(totalItemsDiscAmount,
			decimal.New(0, 2).Mul(bdi.UnitDiscountAmount.Big, bdi.Quantity.Big),
		)
		totalItemsTaxAmount = totalItemsTaxAmount.Add(totalItemsTaxAmount,
			decimal.New(0, 2).Mul(bdi.UnitTaxAmount.Big, bdi.Quantity.Big),
		)
		totalItemsShippingFees = totalItemsShippingFees.Add(totalItemsShippingFees,
			decimal.New(0, 2).Mul(bdi.UnitShippingFees.Big, bdi.Quantity.Big),
		)
	}

	// update invoice amounts
	bd.GrossAmountGen = types.NewNullDecimal(totalItemsAmount)
	bd.DiscountAmountGen = types.NewNullDecimal(totalItemsDiscAmount)
	bd.TaxAmountGen = types.NewNullDecimal(totalItemsTaxAmount)
	bd.ShippingFeesGen = types.NewNullDecimal(totalItemsShippingFees)

	return nil
}

func ProcessBaseDocumentAmountsForPreview(bd *base.BaseDocument, bdis []*base.BaseDocumentItem) error {
	totalItemsAmount := decimal.New(0, 2)
	totalItemsDiscAmount := decimal.New(0, 2)
	totalItemsTaxAmount := decimal.New(0, 2)
	totalItemsShippingFees := decimal.New(0, 2)

	for _, bdi := range bdis {
		quantity := bdi.Quantity.Big
		unitPrice := bdi.UnitPrice.Big
		unitDiscountAmount := bdi.UnitDiscountAmount.Big
		unitTaxAmount := bdi.UnitTaxAmount.Big
		unitShippingFees := bdi.UnitShippingFees.Big

		totalItemsAmount = totalItemsAmount.Add(totalItemsAmount, decimal.New(0, 2).Mul(unitPrice, quantity))
		totalItemsDiscAmount = totalItemsDiscAmount.Add(totalItemsDiscAmount, decimal.New(0, 2).Mul(unitDiscountAmount, quantity))
		totalItemsTaxAmount = totalItemsTaxAmount.Add(totalItemsTaxAmount, decimal.New(0, 2).Mul(unitTaxAmount, quantity))
		totalItemsShippingFees = totalItemsShippingFees.Add(totalItemsShippingFees, decimal.New(0, 2).Mul(unitShippingFees, quantity))

		// Calculate generated fields for base_document_item
		bdi.TotalDiscountAmountGen = types.NewNullDecimal(decimal.New(0, 2).Mul(unitDiscountAmount, quantity))

		drt := decimal.New(0, 2).Quo(unitDiscountAmount, unitPrice)
		drt = drt.Mul(drt, decimal.New(10000, 2))
		bdi.DiscountRateGen = types.NewNullDecimal(drt)

		bdi.TotalTaxAmountGen = types.NewNullDecimal(decimal.New(0, 2).Mul(unitTaxAmount, quantity))

		trg := decimal.New(0, 2).Quo(unitTaxAmount, unitPrice)
		trg = trg.Mul(trg, decimal.New(10000, 2))
		bdi.TaxRateGen = types.NewNullDecimal(trg)

		bdi.TotalShippingFeesGen = types.NewNullDecimal(decimal.New(0, 2).Mul(unitShippingFees, quantity))

		fupg := decimal.New(0, 2).Sub(unitPrice, unitDiscountAmount)
		fupg = fupg.Add(fupg, unitTaxAmount)
		fupg = fupg.Add(fupg, unitShippingFees)
		bdi.FinalUnitPriceGen = types.NewNullDecimal(fupg)

		bdi.TotalSalePriceGen = types.NewNullDecimal(decimal.New(0, 2).Mul(quantity, bdi.FinalUnitPriceGen.Big))
	}

	// update invoice amounts
	bd.GrossAmountGen = types.NewNullDecimal(totalItemsAmount)
	bd.DiscountAmountGen = types.NewNullDecimal(totalItemsDiscAmount)
	bd.TaxAmountGen = types.NewNullDecimal(totalItemsTaxAmount)
	bd.ShippingFeesGen = types.NewNullDecimal(totalItemsShippingFees)

	// Calculate generated fields for base_document
	drg := decimal.New(0, 2).Quo(totalItemsDiscAmount, totalItemsAmount)
	drg = drg.Mul(drg, decimal.New(10000, 2))
	bd.DiscountRateGen = types.NewNullDecimal(drg)

	adrg := decimal.New(0, 2).Quo(bd.AdditionalDiscountAmount.Big, totalItemsAmount)
	adrg = adrg.Mul(adrg, decimal.New(10000, 2))
	bd.AdditionalDiscountRateGen = types.NewNullDecimal(adrg)

	gaadg := decimal.New(0, 2).Sub(totalItemsAmount, totalItemsDiscAmount)
	gaadg = gaadg.Sub(gaadg, bd.AdditionalDiscountAmount.Big)
	bd.GrossAmountAfterDiscountGen = types.NewNullDecimal(gaadg)

	trg := decimal.New(0, 2).Quo(totalItemsTaxAmount, totalItemsAmount)
	trg = trg.Mul(trg, decimal.New(10000, 2))
	bd.TaxRateGen = types.NewNullDecimal(trg)

	// nag := decimal.New(0, 2).Sub(totalItemsAmount, totalItemsDiscAmount)
	// nag = nag.Sub(nag, bd.AdditionalDiscountAmount.Big)
	// nag = nag.Add(nag, totalItemsTaxAmount)
	// nag = nag.Add(nag, totalItemsShippingFees)
	// nag = nag.Add(nag, bd.OtherFees.Big)
	// nag = nag.Add(nag, bd.CustomAdjustmentAmount.Big)
	// bd.NettAmountGen = types.NewNullDecimal(nag)
	processBaseDocumentNetAmount(bd)

	return nil
}

func processBaseDocumentNetAmount(bd *base.BaseDocument) error {
	nag := decimal.New(0, 2).Sub(bd.GrossAmountGen.Big, bd.DiscountAmountGen.Big)
	nag = nag.Sub(nag, bd.AdditionalDiscountAmount.Big)
	nag = nag.Add(nag, bd.TaxAmountGen.Big)
	nag = nag.Add(nag, bd.ShippingFeesGen.Big)
	nag = nag.Add(nag, bd.OtherFees.Big)
	nag = nag.Add(nag, bd.CustomAdjustmentAmount.Big)
	bd.NetAmountGen = types.NewNullDecimal(nag)

	return nil
}
