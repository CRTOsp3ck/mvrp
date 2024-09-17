package invoice

import (
	"context"
	"mvrp/data/model/invoice"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) GetInvoiceByBaseDocumentID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.Invoice, error) {
	return invoice.Invoices(qm.Where("base_document_id = ?", id)).One(ctx, exec)
}
