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

func (s *InvoiceService) NewListCreditNoteRequest(ctx context.Context) *ListCreditNoteRequest {
	return &ListCreditNoteRequest{Ctx: ctx}
}

type ListCreditNoteResponse struct {
	Payload invoice.CreditNoteSlice `json:"payload"`
}

func (s *InvoiceService) NewListCreditNoteResponse(payload invoice.CreditNoteSlice) *ListCreditNoteResponse {
	return &ListCreditNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewSearchCreditNoteRequest(ctx context.Context, payload dto.SearchCreditNoteDTO) *SearchCreditNoteRequest {
	return &SearchCreditNoteRequest{Ctx: ctx, Payload: payload}
}

type SearchCreditNoteResponse struct {
	Payload    invoice.CreditNoteSlice `json:"payload"`
	Pagination dto.PaginationDTO       `json:"pagination"`
}

func (s *InvoiceService) NewSearchCreditNoteResponse(payload invoice.CreditNoteSlice) *SearchCreditNoteResponse {
	return &SearchCreditNoteResponse{Payload: payload}
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

	// Pagination
	totalCount, err := s.Repo.Invoice.GetCreditNoteTotalCount(req.Ctx, tx)
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
	resp := SearchCreditNoteResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET DELIVERY NOTE
type GetCreditNoteRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewGetCreditNoteRequest(ctx context.Context, id int) *GetCreditNoteRequest {
	return &GetCreditNoteRequest{Ctx: ctx, ID: id}
}

type GetCreditNoteResponse struct {
	Payload invoice.CreditNote `json:"payload"`
}

func (s *InvoiceService) NewGetCreditNoteResponse(payload invoice.CreditNote) *GetCreditNoteResponse {
	return &GetCreditNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewCreateCreditNoteRequest(ctx context.Context, payload dto.CreateCreditNoteDTO) *CreateCreditNoteRequest {
	return &CreateCreditNoteRequest{Ctx: ctx, Payload: payload}
}

type CreateCreditNoteResponse struct {
	Payload invoice.CreditNote `json:"payload"`
}

func (s *InvoiceService) NewCreateCreditNoteResponse(payload invoice.CreditNote) *CreateCreditNoteResponse {
	return &CreateCreditNoteResponse{Payload: payload}
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
	nextID, err := s.Repo.Base.GetNextEntryBaseDocumentID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.BaseDocument.ID = nextID
	err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// create delivery note
	nextID, err = s.Repo.Invoice.GetNextEntryCreditNoteID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.CreditNote.ID = nextID
	if req.Payload.CreditNote.CreditNoteNumber == "" {
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

func (s *InvoiceService) NewUpdateCreditNoteRequest(ctx context.Context, payload dto.UpdateCreditNoteDTO) *UpdateCreditNoteRequest {
	return &UpdateCreditNoteRequest{Ctx: ctx, Payload: payload}
}

type UpdateCreditNoteResponse struct {
	Payload invoice.CreditNote `json:"payload"`
}

func (s *InvoiceService) NewUpdateCreditNoteResponse(payload invoice.CreditNote) *UpdateCreditNoteResponse {
	return &UpdateCreditNoteResponse{Payload: payload}
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

func (s *InvoiceService) NewDeleteCreditNoteRequest(ctx context.Context, id int) *DeleteCreditNoteRequest {
	return &DeleteCreditNoteRequest{Ctx: ctx, ID: id}
}

type DeleteCreditNoteResponse struct {
	Payload bool `json:"payload"`
}

func (s *InvoiceService) NewDeleteCreditNoteResponse(payload bool) *DeleteCreditNoteResponse {
	return &DeleteCreditNoteResponse{Payload: payload}
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
