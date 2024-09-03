package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
	"mvrp/util"
)

// LIST DEBIT NOTE
type ListDebitNoteRequest struct {
	Ctx context.Context
}
type ListDebitNoteResponse struct {
	Payload invoice.DebitNoteSlice
}

func (s *InvoiceService) ListDebitNote(req *ListDebitNoteRequest) (*ListDebitNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.ListAllDebitNotes(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListDebitNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DEBIT NOTE
type SearchDebitNoteRequest struct {
	Ctx     context.Context
	Payload dto.SearchDebitNoteDTO
}
type SearchDebitNoteResponse struct {
	Payload invoice.DebitNoteSlice
}

func (s *InvoiceService) SearchDebitNote(req *SearchDebitNoteRequest) (*SearchDebitNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.SearchDebitNotes(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchDebitNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET DEBIT NOTE
type GetDebitNoteRequest struct {
	Ctx context.Context
	ID  int
}
type GetDebitNoteResponse struct {
	Payload invoice.DebitNote
}

func (s *InvoiceService) GetDebitNote(req *GetDebitNoteRequest) (*GetDebitNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.GetDebitNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetDebitNoteResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE DEBIT NOTE
type CreateDebitNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateDebitNoteDTO
}
type CreateDebitNoteResponse struct {
	Payload invoice.DebitNote
}

func (s *InvoiceService) CreateDebitNote(req *CreateDebitNoteRequest) (*CreateDebitNoteResponse, error) {
	/*
		1. Create Base Document
		2. Create Debit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create base document
	err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// create debit note
	if req.Payload.DebitNote.DebitNoteNumber == "" {
		nextID, err := s.Repo.Invoice.GetNextEntryDebitNoteID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.DebitNote.DebitNoteNumber = util.Util.Str.ToString(nextID)
	}
	req.Payload.DebitNote.BaseDocumentID = req.Payload.BaseDocument.ID
	err = s.Repo.Invoice.CreateDebitNote(req.Ctx, tx, &req.Payload.DebitNote)
	if err != nil {
		return nil, err
	}

	// get created debit note
	DebitNote, err := s.Repo.Invoice.GetDebitNoteByID(req.Ctx, tx, req.Payload.DebitNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateDebitNoteResponse{
		Payload: *DebitNote,
	}

	return &resp, nil
}

// UPDATE DEBIT NOTE
type UpdateDebitNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateDebitNoteDTO
}
type UpdateDebitNoteResponse struct {
	Payload invoice.DebitNote
}

func (s *InvoiceService) UpdateDebitNote(req *UpdateDebitNoteRequest) (*UpdateDebitNoteResponse, error) {
	/*
		1. Update Base Document
		2. Update Debit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// update base document
	err = s.Repo.Base.UpdateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// update debit note
	err = s.Repo.Invoice.UpdateDebitNote(req.Ctx, tx, &req.Payload.DebitNote)
	if err != nil {
		return nil, err
	}

	// get updated debit note
	DebitNote, err := s.Repo.Invoice.GetDebitNoteByID(req.Ctx, tx, req.Payload.DebitNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateDebitNoteResponse{
		Payload: *DebitNote,
	}

	return &resp, nil
}

// DELETE DEBIT NOTE
type DeleteDebitNoteRequest struct {
	Ctx context.Context
	ID  int
}
type DeleteDebitNoteResponse struct {
	Payload bool
}

func (s *InvoiceService) DeleteDebitNote(req *DeleteDebitNoteRequest) (*DeleteDebitNoteResponse, error) {
	/*
		1. Delete Base Document
		2. Delete Debit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get debit note
	debitNote, err := s.Repo.Invoice.GetDebitNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, debitNote.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete debit note
	err = s.Repo.Invoice.DeleteDebitNote(req.Ctx, tx, debitNote)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteDebitNoteResponse{
		Payload: true,
	}

	return &resp, nil
}
