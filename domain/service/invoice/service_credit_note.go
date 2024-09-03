package invoice

import (
	"context"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
	"mvrp/util"
)

// LIST DELIVERY NOTE
type ListCreditNoteRequest struct {
	Ctx context.Context
}
type ListCreditNoteResponse struct {
	Payload invoice.CreditNoteSlice
}

func (s *InvoiceService) ListCreditNote(req *ListCreditNoteRequest) (*ListCreditNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.ListAllCreditNotes(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListCreditNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH DELIVERY NOTE
type SearchCreditNoteRequest struct {
	Ctx     context.Context
	Payload dto.SearchCreditNoteDTO
}
type SearchCreditNoteResponse struct {
	Payload invoice.CreditNoteSlice
}

func (s *InvoiceService) SearchCreditNote(req *SearchCreditNoteRequest) (*SearchCreditNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.SearchCreditNotes(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchCreditNoteResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetCreditNoteRequest struct {
	Ctx context.Context
	ID  int
}
type GetCreditNoteResponse struct {
	Payload invoice.CreditNote
}

func (s *InvoiceService) GetCreditNote(req *GetCreditNoteRequest) (*GetCreditNoteResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.GetCreditNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetCreditNoteResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE DELIVERY NOTE
type CreateCreditNoteRequest struct {
	Ctx     context.Context
	Payload dto.CreateCreditNoteDTO
}
type CreateCreditNoteResponse struct {
	Payload invoice.CreditNote
}

func (s *InvoiceService) CreateCreditNote(req *CreateCreditNoteRequest) (*CreateCreditNoteResponse, error) {
	/*
		1. Create Base Document
		2. Create Credit Note
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

	// create delivery note
	if req.Payload.CreditNote.CreditNoteNumber == "" {
		nextID, err := s.Repo.Invoice.GetNextEntryCreditNoteID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.CreditNote.CreditNoteNumber = util.Util.Str.ToString(nextID)
	}
	req.Payload.CreditNote.BaseDocumentID = req.Payload.BaseDocument.ID
	err = s.Repo.Invoice.CreateCreditNote(req.Ctx, tx, &req.Payload.CreditNote)
	if err != nil {
		return nil, err
	}

	// get created delivery note
	CreditNote, err := s.Repo.Invoice.GetCreditNoteByID(req.Ctx, tx, req.Payload.CreditNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateCreditNoteResponse{
		Payload: *CreditNote,
	}

	return &resp, nil
}

// UPDATE DELIVERY NOTE
type UpdateCreditNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateCreditNoteDTO
}
type UpdateCreditNoteResponse struct {
	Payload invoice.CreditNote
}

func (s *InvoiceService) UpdateCreditNote(req *UpdateCreditNoteRequest) (*UpdateCreditNoteResponse, error) {
	/*
		1. Update Base Document
		2. Update Credit Note
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

	// update delivery note
	err = s.Repo.Invoice.UpdateCreditNote(req.Ctx, tx, &req.Payload.CreditNote)
	if err != nil {
		return nil, err
	}

	// get updated delivery note
	CreditNote, err := s.Repo.Invoice.GetCreditNoteByID(req.Ctx, tx, req.Payload.CreditNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateCreditNoteResponse{
		Payload: *CreditNote,
	}

	return &resp, nil
}

// DELETE DELIVERY NOTE
type DeleteCreditNoteRequest struct {
	Ctx context.Context
	ID  int
}
type DeleteCreditNoteResponse struct {
	Payload bool
}

func (s *InvoiceService) DeleteCreditNote(req *DeleteCreditNoteRequest) (*DeleteCreditNoteResponse, error) {
	/*
		1. Delete Base Document
		2. Delete Credit Note
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get credit note
	creditNote, err := s.Repo.Invoice.GetCreditNoteByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, creditNote.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete credit note
	err = s.Repo.Invoice.DeleteCreditNote(req.Ctx, tx, creditNote)
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

	resp := DeleteCreditNoteResponse{
		Payload: true,
	}

	return &resp, nil
}
