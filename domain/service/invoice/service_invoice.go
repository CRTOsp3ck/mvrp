package invoice

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/invoice"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST INVOICE
type ListInvoiceRequest struct {
	Ctx context.Context
}

func (s *InvoiceService) NewListInvoiceRequest(ctx context.Context) *ListInvoiceRequest {
	return &ListInvoiceRequest{Ctx: ctx}
}

type ListInvoiceResponse struct {
	Payload invoice.InvoiceSlice `json:"payload"`
}

func (s *InvoiceService) NewListInvoiceResponse(payload invoice.InvoiceSlice) *ListInvoiceResponse {
	return &ListInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) ListInvoice(req *ListInvoiceRequest) (*ListInvoiceResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.ListAllInvoices(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListInvoiceResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW INVOICE
type PreviewInvoiceRequest struct {
	Ctx     context.Context
	Payload dto.CreateInvoiceDTO
}

func (s *InvoiceService) NewPreviewInvoiceRequest(ctx context.Context, payload dto.CreateInvoiceDTO) *PreviewInvoiceRequest {
	return &PreviewInvoiceRequest{Ctx: ctx, Payload: payload}
}

type PreviewInvoiceResponse struct {
	Payload dto.CreateInvoiceDTO `json:"payload"`
}

func (s *InvoiceService) NewPreviewInvoiceResponse(payload dto.CreateInvoiceDTO) *PreviewInvoiceResponse {
	return &PreviewInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) PreviewInvoice(req *PreviewInvoiceRequest) (*PreviewInvoiceResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}
	resp := PreviewInvoiceResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH INVOICE
type SearchInvoiceRequest struct {
	Ctx     context.Context
	Payload dto.SearchInvoiceDTO
}

func (s *InvoiceService) NewSearchInvoiceRequest(ctx context.Context, payload dto.SearchInvoiceDTO) *SearchInvoiceRequest {
	return &SearchInvoiceRequest{Ctx: ctx, Payload: payload}
}

type SearchInvoiceResponse struct {
	Payload    invoice.InvoiceSlice `json:"payload"`
	Pagination dto.PaginationDTO    `json:"pagination"`
}

func (s *InvoiceService) NewSearchInvoiceResponse(payload invoice.InvoiceSlice) *SearchInvoiceResponse {
	return &SearchInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) SearchInvoice(req *SearchInvoiceRequest) (*SearchInvoiceResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.SearchInvoices(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Invoice.GetInvoiceTotalCount(req.Ctx, tx)
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
	resp := SearchInvoiceResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET INVOICE
type GetInvoiceRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewGetInvoiceRequest(ctx context.Context, id int) *GetInvoiceRequest {
	return &GetInvoiceRequest{Ctx: ctx, ID: id}
}

type GetInvoiceResponse struct {
	Payload invoice.Invoice `json:"payload"`
}

func (s *InvoiceService) NewGetInvoiceResponse(payload invoice.Invoice) *GetInvoiceResponse {
	return &GetInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) GetInvoice(req *GetInvoiceRequest) (*GetInvoiceResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Invoice.GetInvoiceByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := GetInvoiceResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE INVOICE
type CreateInvoiceRequest struct {
	Ctx     context.Context
	Payload dto.CreateInvoiceDTO
}

func (s *InvoiceService) NewCreateInvoiceRequest(ctx context.Context, payload dto.CreateInvoiceDTO) *CreateInvoiceRequest {
	return &CreateInvoiceRequest{Ctx: ctx, Payload: payload}
}

type CreateInvoiceResponse struct {
	Payload invoice.Invoice `json:"payload"`
}

func (s *InvoiceService) NewCreateInvoiceResponse(payload invoice.Invoice) *CreateInvoiceResponse {
	return &CreateInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) CreateInvoice(req *CreateInvoiceRequest) (*CreateInvoiceResponse, error) {
	/*
		1. Create Base Document
		2. Create Invoice
		3. Create Base Document Items
		4. Create Invoice Items
	*/
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err = proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

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

	// create invoice
	req.Payload.Invoice.BaseDocumentID = req.Payload.BaseDocument.ID
	nextID, err = s.Repo.Invoice.GetNextEntryInvoiceID(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	req.Payload.Invoice.ID = nextID
	if req.Payload.Invoice.InvoiceNumber == "" {
		req.Payload.Invoice.InvoiceNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Invoice.CreateInvoice(req.Ctx, tx, &req.Payload.Invoice)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Payload.Items {
		// create base document item
		nextID, err = s.Repo.Base.GetNextEntryBaseDocumentItemID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		item.BaseDocumentItem.ID = nextID
		item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
		err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
		if err != nil {
			return nil, err
		}

		// create invoice item
		nextID, err = s.Repo.Invoice.GetNextEntryInvoiceItemID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		item.InvoiceItem.ID = nextID
		item.InvoiceItem.InvoiceID = req.Payload.Invoice.ID
		item.InvoiceItem.BaseDocumentItemID = item.BaseDocumentItem.ID
		err = s.Repo.Invoice.CreateInvoiceItem(req.Ctx, tx, &item.InvoiceItem)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.Repo.Invoice.GetInvoiceByID(req.Ctx, tx, req.Payload.Invoice.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateInvoiceResponse{
		Payload: *res,
	}
	return &resp, nil
}

// UPDATE INVOICE
type UpdateInvoiceRequest struct {
	Ctx     context.Context
	Payload dto.UpdateInvoiceDTO
}

func (s *InvoiceService) NewUpdateInvoiceRequest(ctx context.Context, payload dto.UpdateInvoiceDTO) *UpdateInvoiceRequest {
	return &UpdateInvoiceRequest{Ctx: ctx, Payload: payload}
}

type UpdateInvoiceResponse struct {
	Payload invoice.Invoice `json:"payload"`
}

func (s *InvoiceService) NewUpdateInvoiceResponse(payload invoice.Invoice) *UpdateInvoiceResponse {
	return &UpdateInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) UpdateInvoice(req *UpdateInvoiceRequest) (*UpdateInvoiceResponse, error) {
	/*
		1. Update Base Document
		2. Update Invoice
		3. Update Base Document Items
		4. Update Invoice Items
	*/
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err = proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	// update base document
	err = s.Repo.Base.UpdateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// update invoice
	err = s.Repo.Invoice.UpdateInvoice(req.Ctx, tx, &req.Payload.Invoice)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	currItems, err := s.Repo.Invoice.GetInvoiceItemsByInvoiceID(req.Ctx, tx, req.Payload.Invoice.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range currItems {
		found := false
		for _, newItem := range req.Payload.Items {
			if item.ID == newItem.InvoiceItem.ID {
				found = true
				break
			}
		}
		if !found {
			// get base document item
			baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItemID)
			if err != nil {
				return nil, err
			}

			// delete base document item
			err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
			if err != nil {
				return nil, err
			}

			// delete invoice item
			err = s.Repo.Invoice.DeleteInvoiceItem(req.Ctx, tx, item)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update invoice items
	for _, item := range req.Payload.Items {
		// check if the item is new or existing
		itemExists, err := s.Repo.Invoice.InvoiceItemExists(req.Ctx, tx, item.InvoiceItem.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			currBaseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItem.ID)
			if err != nil {
				return nil, err
			}
			amountOffset := types.NewNullDecimal(decimal.New(0, 2))
			amountOffset.Sub(item.Quantity.Big, currBaseDocumentItem.Quantity.Big)

			// update base document items
			err = s.Repo.Base.UpdateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// update invoice items
			err = s.Repo.Invoice.UpdateInvoiceItem(req.Ctx, tx, &item.InvoiceItem)
			if err != nil {
				return nil, err
			}
		} else {
			// create base document items
			nextID, err := s.Repo.Base.GetNextEntryBaseDocumentItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.BaseDocumentItem.ID = nextID
			item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create invoice items
			nextID, err = s.Repo.Invoice.GetNextEntryInvoiceItemID(req.Ctx, tx)
			if err != nil {
				return nil, err
			}
			item.InvoiceItem.ID = nextID
			item.InvoiceItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.InvoiceItem.InvoiceID = req.Payload.Invoice.ID
			err = s.Repo.Invoice.CreateInvoiceItem(req.Ctx, tx, &item.InvoiceItem)
			if err != nil {
				return nil, err
			}
		}
	}

	res, err := s.Repo.Invoice.GetInvoiceByID(req.Ctx, tx, req.Payload.Invoice.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateInvoiceResponse{
		Payload: *res,
	}
	return &resp, nil
}

// DELETE INVOICE
type DeleteInvoiceRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InvoiceService) NewDeleteInvoiceRequest(ctx context.Context, id int) *DeleteInvoiceRequest {
	return &DeleteInvoiceRequest{Ctx: ctx, ID: id}
}

type DeleteInvoiceResponse struct {
	Payload bool `json:"payload"`
}

func (s *InvoiceService) NewDeleteInvoiceResponse(payload bool) *DeleteInvoiceResponse {
	return &DeleteInvoiceResponse{Payload: payload}
}

func (s *InvoiceService) DeleteInvoice(req *DeleteInvoiceRequest) (*DeleteInvoiceResponse, error) {
	/*
		1. Delete Base Document
		2. Delete Invoice
		3. Delete Base Document Items
		4. Delete Invoice Items
	*/
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get invoice
	inv, err := s.Repo.Invoice.GetInvoiceByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete invoice
	err = s.Repo.Invoice.DeleteInvoice(req.Ctx, tx, inv)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, inv.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	items, err := s.Repo.Invoice.GetInvoiceItemsByInvoiceID(req.Ctx, tx, inv.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		// get base document item
		baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, item.BaseDocumentItemID)
		if err != nil {
			return nil, err
		}

		// delete base document item
		err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
		if err != nil {
			return nil, err
		}

		// delete invoice item
		err = s.Repo.Invoice.DeleteInvoiceItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteInvoiceResponse{
		Payload: true,
	}
	return &resp, nil
}
