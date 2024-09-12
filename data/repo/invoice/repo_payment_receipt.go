// Code generated by MVRP Codegen Util. DO NOT EDIT.

package invoice

import (
	"context"
	"database/sql"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *InvoiceRepository) ListAllPaymentReceipts(ctx context.Context, exec boil.ContextExecutor) (invoice.PaymentReceiptSlice, error) {
	return invoice.PaymentReceipts().All(ctx, exec)
}

/*
func (r *InvoiceRepository) SearchPaymentReceipts(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchPaymentReceiptDTO) (invoice.PaymentReceiptSlice, error) {
	return invoice.PaymentReceipts(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *InvoiceRepository) SearchPaymentReceipts(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchPaymentReceiptDTO) (invoice.PaymentReceiptSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return invoice.PaymentReceipts(queryMods...).All(ctx, exec)
}

func (r *InvoiceRepository) GetPaymentReceiptByID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.PaymentReceipt, error) {
	return invoice.FindPaymentReceipt(ctx, exec, id)
}

func (r *InvoiceRepository) CreatePaymentReceipt(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceipt) error {
	/*
		id, err := r.GetNextEntryPaymentReceiptID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InvoiceRepository) UpdatePaymentReceipt(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceipt) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InvoiceRepository) UpsertPaymentReceipt(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceipt) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InvoiceRepository) DeletePaymentReceipt(ctx context.Context, exec boil.ContextExecutor, m *invoice.PaymentReceipt) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InvoiceRepository) PaymentReceiptExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return invoice.PaymentReceiptExists(ctx, exec, id)
}

func (r *InvoiceRepository) GetPaymentReceiptRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.PaymentReceipts().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) GetMostRecentPaymentReceipt(ctx context.Context, exec boil.ContextExecutor) (*invoice.PaymentReceipt, error) {
	return invoice.PaymentReceipts(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InvoiceRepository) GetNextEntryPaymentReceiptID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := invoice.PaymentReceipts(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentPaymentReceipt(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *InvoiceRepository) GetPaymentReceiptTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.PaymentReceipts().Count(ctx, exec)
	return int(count), err
}