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

func (r *InvoiceRepository) ListAllDebitNotes(ctx context.Context, exec boil.ContextExecutor) (invoice.DebitNoteSlice, error) {
	return invoice.DebitNotes().All(ctx, exec)
}

/*
func (r *InvoiceRepository) SearchDebitNotes(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchDebitNoteDTO) (invoice.DebitNoteSlice, error) {
	return invoice.DebitNotes(
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	).All(ctx, exec)
}
*/
func (r *InvoiceRepository) SearchDebitNotes(ctx context.Context, exec boil.ContextExecutor, dto dto.SearchDebitNoteDTO) (invoice.DebitNoteSlice, error) {
	var queryMods []qm.QueryMod

	queryMods = append(queryMods,
		qm.Limit(dto.ItemsPerPage),
		qm.Offset((dto.ItemsPerPage*dto.Page)-dto.ItemsPerPage),
		// qm.GroupBy("id"),
		qm.OrderBy(dto.OrderBy+" "+"ASC"),
	)

	return invoice.DebitNotes(queryMods...).All(ctx, exec)
}

func (r *InvoiceRepository) GetDebitNoteByID(ctx context.Context, exec boil.ContextExecutor, id int) (*invoice.DebitNote, error) {
	return invoice.FindDebitNote(ctx, exec, id)
}

func (r *InvoiceRepository) CreateDebitNote(ctx context.Context, exec boil.ContextExecutor, m *invoice.DebitNote) error {
	/*
		id, err := r.GetNextEntryDebitNoteID(ctx, exec)
		if err != nil {
			return err
		}
		m.ID = id
	*/
	return m.Insert(ctx, exec, boil.Infer())
}

func (r *InvoiceRepository) UpdateDebitNote(ctx context.Context, exec boil.ContextExecutor, m *invoice.DebitNote) error {
	_, err := m.Update(ctx, exec, boil.Infer())
	return err
}

func (r *InvoiceRepository) UpsertDebitNote(ctx context.Context, exec boil.ContextExecutor, m *invoice.DebitNote) error {
	return m.Upsert(ctx, exec, true, nil, boil.Infer(), boil.Infer())
}

func (r *InvoiceRepository) DeleteDebitNote(ctx context.Context, exec boil.ContextExecutor, m *invoice.DebitNote) error {
	_, err := m.Delete(ctx, exec)
	return err
}

func (r *InvoiceRepository) DebitNoteExists(ctx context.Context, exec boil.ContextExecutor, id int) (bool, error) {
	return invoice.DebitNoteExists(ctx, exec, id)
}

func (r *InvoiceRepository) GetDebitNoteRowsCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.DebitNotes().Count(ctx, exec)
	return int(count), err
}

func (r *InvoiceRepository) GetMostRecentDebitNote(ctx context.Context, exec boil.ContextExecutor) (*invoice.DebitNote, error) {
	return invoice.DebitNotes(qm.OrderBy("created_at DESC")).One(ctx, exec)
}

func (r *InvoiceRepository) GetNextEntryDebitNoteID(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	var maxID sql.NullInt64
	err := invoice.DebitNotes(qm.Select("MAX(id)")).QueryRow(exec).Scan(&maxID)
	if err != nil {
		return 0, err
	}

	// Check if maxID is valid (non-NULL), otherwise return 1
	if !maxID.Valid {
		return 1, nil
	}
	return int(maxID.Int64) + 1, nil

	/*
		currID, err := r.GetMostRecentDebitNote(ctx, exec)
		if err != nil {
			if err == sql.ErrNoRows {
				return 1, nil
			}
			return 0, err
		}
		return currID.ID + 1, nil
	*/
}

func (r *InvoiceRepository) GetDebitNoteTotalCount(ctx context.Context, exec boil.ContextExecutor) (int, error) {
	count, err := invoice.DebitNotes().Count(ctx, exec)
	return int(count), err
}