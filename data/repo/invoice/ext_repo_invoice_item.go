package invoice

import (
	"context"
	"mvrp/data/model/invoice"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) GetInvoiceItemsByInvoiceID(ctx context.Context, exec boil.ContextExecutor, id int) (invoice.InvoiceItemSlice, error) {
	return invoice.InvoiceItems(qm.Where("invoice_id = ?", id)).All(ctx, exec)
}

func (r *InvoiceRepository) GetInvoiceItemByBaseDocumentItemID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.InvoiceItem, error) {
	return invoice.InvoiceItems(qm.Where("base_document_item_id = ?", id)).One(ctx, exec)
}
