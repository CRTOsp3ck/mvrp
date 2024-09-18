package invoice

import (
	"context"
	"mvrp/data/model/invoice"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) GetCreditNoteItemByBaseDocumentItemID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.CreditNoteItem, error) {
	return invoice.CreditNoteItems(qm.Where("base_document_item_id = ?", id)).One(ctx, exec)
}

func (r *InvoiceRepository) GetCreditNoteItemsByCreditNoteID(ctx context.Context, exec boil.ContextExecutor, id int) (invoice.CreditNoteItemSlice, error) {
	return invoice.CreditNoteItems(qm.Where("credit_note_id = ?", id)).All(ctx, exec)
}
