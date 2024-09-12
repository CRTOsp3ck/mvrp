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

func (s *InvoiceService) NewListDebitNoteRequest(ctx context.Context) *ListDebitNoteRequest {
	return &ListDebitNoteRequest{Ctx: ctx}
}

type ListDebitNoteResponse struct {
	Payload invoice.DebitNoteSlice `json:"payload"`
}

func (s *InvoiceService) NewListDebitNoteResponse(payload invoice.DebitNoteSlice) *ListDebitNoteResponse {
	return &ListDebitNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewSearchDebitNoteRequest(ctx context.Context, payload dto.SearchDebitNoteDTO) *SearchDebitNoteRequest {
	return &SearchDebitNoteRequest{Ctx: ctx, Payload: payload}
}

type SearchDebitNoteResponse struct {
	Payload    invoice.DebitNoteSlice `json:"payload"`
	Pagination dto.PaginationDTO      `json:"pagination"`
}

func (s *InvoiceService) NewSearchDebitNoteResponse(payload invoice.DebitNoteSlice) *SearchDebitNoteResponse {
	return &SearchDebitNoteResponse{Payload: payload}
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

	// Pagination
	totalCount, err := s.Repo.Invoice.GetDebitNoteTotalCount(req.Ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	pd := dto.PaginationDTO{
		TotalItems:   totalCount,
		ItemsPerPage: req.Payload.ItemsPerPage,
		Page:         req.Payload.Page,
		SortBy:       req.Payload.SortBy,
		OrderBy:      req.Payload.OrderBy,
	}
	resp := SearchDebitNoteResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DEBIT NOTE
type GetDebitNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewGetDebitNoteRequest(ctx context.Context, id int) *GetDebitNoteRequest {
	return &GetDebitNoteRequest{Ctx: ctx, ID: id}
}

type GetDebitNoteResponse struct {
	Payload invoice.DebitNote `json:"payload"`
}

func (s *InvoiceService) NewGetDebitNoteResponse(payload invoice.DebitNote) *GetDebitNoteResponse {
	return &GetDebitNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewCreateDebitNoteRequest(ctx context.Context, payload dto.CreateDebitNoteDTO) *CreateDebitNoteRequest {
	return &CreateDebitNoteRequest{Ctx: ctx, Payload: payload}
}

type CreateDebitNoteResponse struct {
	Payload invoice.DebitNote `json:"payload"`
}

func (s *InvoiceService) NewCreateDebitNoteResponse(payload invoice.DebitNote) *CreateDebitNoteResponse {
	return &CreateDebitNoteResponse{Payload: payload}
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
	nextID, err := s.Repo.Base.GetNextEntryBaseDocumentID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.BaseDocument.ID = nextID
	err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// create debit note
	nextID, err = s.Repo.Invoice.GetNextEntryDebitNoteID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.DebitNote.ID = nextID
	if req.Payload.DebitNote.DebitNoteNumber == "" {
		req.Payload.DebitNote.DebitNoteNumber = util.Util.Str.ToString(nextID)
	}
	req.Payload.DebitNote.BaseDocumentID = req.Payload.BaseDocument.ID
	err = s.Repo.Invoice.CreateDebitNote(req.Ctx, tx, &req.Payload.DebitNote)
	if err != nil {
		return nil, err
	}

	// get created debit note
	dn, err := s.Repo.Invoice.GetDebitNoteByID(req.Ctx, tx, req.Payload.DebitNote.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateDebitNoteResponse{
		Payload: *dn,
	}

	return &resp, nil
}

// UPDATE DEBIT NOTE
type UpdateDebitNoteRequest struct {
	Ctx     context.Context
	Payload dto.UpdateDebitNoteDTO
}

func (s *InvoiceService) NewUpdateDebitNoteRequest(ctx context.Context, payload dto.UpdateDebitNoteDTO) *UpdateDebitNoteRequest {
	return &UpdateDebitNoteRequest{Ctx: ctx, Payload: payload}
}

type UpdateDebitNoteResponse struct {
	Payload invoice.DebitNote `json:"payload"`
}

func (s *InvoiceService) NewUpdateDebitNoteResponse(payload invoice.DebitNote) *UpdateDebitNoteResponse {
	return &UpdateDebitNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewDeleteDebitNoteRequest(ctx context.Context, id int) *DeleteDebitNoteRequest {
	return &DeleteDebitNoteRequest{Ctx: ctx, ID: id}
}

type DeleteDebitNoteResponse struct {
	Payload bool `json:"payload"`
}

func (s *InvoiceService) NewDeleteDebitNoteResponse(payload bool) *DeleteDebitNoteResponse {
	return &DeleteDebitNoteResponse{Payload: payload}
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
