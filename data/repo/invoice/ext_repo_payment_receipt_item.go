package invoice

import (
	"context"
	"mvrp/data/model/invoice"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) GetPaymentReceiptItemsByPaymentReceiptID(ctx context.Context, exec boil.ContextExecutor, id int) (invoice.PaymentReceiptItemSlice, error) {
	return invoice.PaymentReceiptItems(qm.Where("invoice_id = ?", id)).All(ctx, exec)
}
