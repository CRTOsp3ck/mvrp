package sale

import (
	"context"
	"mvrp/data/model/base"
	"mvrp/data/model/sale"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"
)

// LIST SALES QUOTATION
type ListSalesQuotationRequest struct {
	Ctx context.Context
}

func (s *SaleService) NewListSalesQuotationRequest(ctx context.Context) *ListSalesQuotationRequest {
	return &ListSalesQuotationRequest{Ctx: ctx}
}

type ListSalesQuotationResponse struct {
	Payload sale.SalesQuotationSlice `json:"payload"`
}

func (s *SaleService) NewListSalesQuotationResponse(payload sale.SalesQuotationSlice) *ListSalesQuotationResponse {
	return &ListSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) ListSalesQuotation(req *ListSalesQuotationRequest) (*ListSalesQuotationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.ListAllSalesQuotations(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListSalesQuotationResponse{
		Payload: res,
	}
	return &resp, nil
}

// PREVIEW SALES QUOTATION
type PreviewSalesQuotationRequest struct {
	Ctx     context.Context
	Payload dto.CreateSalesQuotationDTO
}

func (s *SaleService) NewPreviewSalesQuotationRequest(ctx context.Context, payload dto.CreateSalesQuotationDTO) *PreviewSalesQuotationRequest {
	return &PreviewSalesQuotationRequest{Ctx: ctx, Payload: payload}
}

type PreviewSalesQuotationResponse struct {
	Payload dto.CreateSalesQuotationDTO `json:"payload"`
}

func (s *SaleService) NewPreviewSalesQuotationResponse(payload dto.CreateSalesQuotationDTO) *PreviewSalesQuotationResponse {
	return &PreviewSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) PreviewSalesQuotation(req *PreviewSalesQuotationRequest) (*PreviewSalesQuotationResponse, error) {
	// preprocess amounts
	bdis := make([]*base.BaseDocumentItem, 0)
	for _, item := range req.Payload.Items {
		bdis = append(bdis, &item.BaseDocumentItem)
	}
	err := proc.ProcessBaseDocumentAmounts(&req.Payload.BaseDocument, bdis)
	if err != nil {
		return nil, err
	}

	resp := PreviewSalesQuotationResponse{
		Payload: req.Payload,
	}
	return &resp, nil
}

// SEARCH SALES QUOTATION
type SearchSalesQuotationRequest struct {
	Ctx     context.Context
	Payload dto.SearchSalesQuotationDTO
}

func (s *SaleService) NewSearchSalesQuotationRequest(ctx context.Context, payload dto.SearchSalesQuotationDTO) *SearchSalesQuotationRequest {
	return &SearchSalesQuotationRequest{Ctx: ctx, Payload: payload}
}

type SearchSalesQuotationResponse struct {
	Payload sale.SalesQuotationSlice `json:"payload"`
}

func (s *SaleService) NewSearchSalesQuotationResponse(payload sale.SalesQuotationSlice) *SearchSalesQuotationResponse {
	return &SearchSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) SearchSalesQuotation(req *SearchSalesQuotationRequest) (*SearchSalesQuotationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.SearchSalesQuotations(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := SearchSalesQuotationResponse{
		Payload: res,
	}
	return &resp, nil
}

// GET SALES QUOTATION
type GetSalesQuotationRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewGetSalesQuotationRequest(ctx context.Context, id int) *GetSalesQuotationRequest {
	return &GetSalesQuotationRequest{Ctx: ctx, ID: id}
}

type GetSalesQuotationResponse struct {
	Payload sale.SalesQuotation `json:"payload"`
}

func (s *SaleService) NewGetSalesQuotationResponse(payload sale.SalesQuotation) *GetSalesQuotationResponse {
	return &GetSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) GetSalesQuotation(req *GetSalesQuotationRequest) (*GetSalesQuotationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Sale.GetSalesQuotationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetSalesQuotationResponse{
		Payload: *res,
	}
	return &resp, nil
}

// CREATE SALES QUOTATION
type CreateSalesQuotationRequest struct {
	Ctx     context.Context
	Payload dto.CreateSalesQuotationDTO
}

func (s *SaleService) NewCreateSalesQuotationRequest(ctx context.Context, payload dto.CreateSalesQuotationDTO) *CreateSalesQuotationRequest {
	return &CreateSalesQuotationRequest{Ctx: ctx, Payload: payload}
}

type CreateSalesQuotationResponse struct {
	Payload sale.SalesQuotation `json:"payload"`
}

func (s *SaleService) NewCreateSalesQuotationResponse(payload sale.SalesQuotation) *CreateSalesQuotationResponse {
	return &CreateSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) CreateSalesQuotation(req *CreateSalesQuotationRequest) (*CreateSalesQuotationResponse, error) {
	/*
		1. Preprocess Amounts
		2. Create Base Document
		3. Create Sales Quotation
		4. Create Base Document Items
		5. Create Sales Quotation Items
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
	err = s.Repo.Base.CreateBaseDocument(req.Ctx, tx, &req.Payload.BaseDocument)
	if err != nil {
		return nil, err
	}

	// create sales quotation
	req.Payload.SalesQuotation.BaseDocumentID = req.Payload.BaseDocument.ID
	if req.Payload.SalesQuotation.SalesQuotationNumber == "" {
		nextID, err := s.Repo.Sale.GetNextEntrySalesQuotationID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.SalesQuotation.SalesQuotationNumber = util.Util.Str.ToString(nextID)
	}
	err = s.Repo.Sale.CreateSalesQuotation(req.Ctx, tx, &req.Payload.SalesQuotation)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Payload.Items {
		// create base document items
		item.BaseDocumentItem.BaseDocumentID = req.Payload.BaseDocument.ID
		err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
		if err != nil {
			return nil, err
		}

		// create sales quotation items
		item.SalesQuotationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
		item.SalesQuotationItem.SalesQuotationID = req.Payload.SalesQuotation.ID
		err = s.Repo.Sale.CreateSalesQuotationItem(req.Ctx, tx, &item.SalesQuotationItem)
		if err != nil {
			return nil, err
		}
	}

	// get created sales quotation
	SalesQuotation, err := s.Repo.Sale.GetSalesQuotationByID(req.Ctx, tx, req.Payload.SalesQuotation.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateSalesQuotationResponse{
		Payload: *SalesQuotation,
	}

	return &resp, nil
}

// UPDATE SALES QUOTATION
type UpdateSalesQuotationRequest struct {
	Ctx     context.Context
	Payload dto.UpdateSalesQuotationDTO
}

func (s *SaleService) NewUpdateSalesQuotationRequest(ctx context.Context, payload dto.UpdateSalesQuotationDTO) *UpdateSalesQuotationRequest {
	return &UpdateSalesQuotationRequest{Ctx: ctx, Payload: payload}
}

type UpdateSalesQuotationResponse struct {
	Payload sale.SalesQuotation `json:"payload"`
}

func (s *SaleService) NewUpdateSalesQuotationResponse(payload sale.SalesQuotation) *UpdateSalesQuotationResponse {
	return &UpdateSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) UpdateSalesQuotation(req *UpdateSalesQuotationRequest) (*UpdateSalesQuotationResponse, error) {
	/*
		1. Preprocess Amounts
		2. Update Base Document
		3. Update Sales Quotation
		4. Update Base Document Items
		5. Update Sales Quotation Items
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currSq, err := s.Repo.Sale.GetSalesQuotationByID(req.Ctx, tx, req.Payload.SalesQuotation.ID)
	if err != nil {
		return nil, err
	}

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

	// update sales quotation
	err = s.Repo.Sale.UpdateSalesQuotation(req.Ctx, tx, &req.Payload.SalesQuotation)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list but not in the new list
	for _, currItem := range currSq.R.SalesQuotationItems {
		found := false
		for _, item := range req.Payload.Items {
			if currItem.ID == item.SalesQuotationItem.ID {
				found = true
				break
			}
		}
		if !found {
			// get base document item
			baseDocumentItem, err := s.Repo.Base.GetBaseDocumentItemByID(req.Ctx, tx, currItem.BaseDocumentItemID)
			if err != nil {
				return nil, err
			}

			// delete base document item
			err = s.Repo.Base.DeleteBaseDocumentItem(req.Ctx, tx, baseDocumentItem)
			if err != nil {
				return nil, err
			}

			// delete sales quotation item
			err = s.Repo.Sale.DeleteSalesQuotationItem(req.Ctx, tx, currItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update sales quotation items
	for _, item := range req.Payload.Items {
		// check if item is new
		itemExists, err := s.Repo.Sale.SalesQuotationItemExists(req.Ctx, tx, item.SalesQuotationItem.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			// update base document items
			err = s.Repo.Base.UpdateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// update sales quotation items
			err = s.Repo.Sale.UpdateSalesQuotationItem(req.Ctx, tx, &item.SalesQuotationItem)
			if err != nil {
				return nil, err
			}
		} else {
			// create base document items
			err = s.Repo.Base.CreateBaseDocumentItem(req.Ctx, tx, &item.BaseDocumentItem)
			if err != nil {
				return nil, err
			}

			// create sales quotation items
			item.SalesQuotationItem.BaseDocumentItemID = item.BaseDocumentItem.ID
			item.SalesQuotationItem.SalesQuotationID = req.Payload.SalesQuotation.ID
			err = s.Repo.Sale.CreateSalesQuotationItem(req.Ctx, tx, &item.SalesQuotationItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated sales quotation
	SalesQuotation, err := s.Repo.Sale.GetSalesQuotationByID(req.Ctx, tx, req.Payload.SalesQuotation.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateSalesQuotationResponse{
		Payload: *SalesQuotation,
	}

	return &resp, nil
}

// DELETE SALES QUOTATION
type DeleteSalesQuotationRequest struct {
	Ctx context.Context
	ID  int
}

func (s *SaleService) NewDeleteSalesQuotationRequest(ctx context.Context, id int) *DeleteSalesQuotationRequest {
	return &DeleteSalesQuotationRequest{Ctx: ctx, ID: id}
}

type DeleteSalesQuotationResponse struct {
	Payload bool `json:"payload"`
}

func (s *SaleService) NewDeleteSalesQuotationResponse(payload bool) *DeleteSalesQuotationResponse {
	return &DeleteSalesQuotationResponse{Payload: payload}
}

func (s *SaleService) DeleteSalesQuotation(req *DeleteSalesQuotationRequest) (*DeleteSalesQuotationResponse, error) {
	/*
		1. Get Sales Quotation
		2. Delete Sales Quotation
		3. Get Base Document
		4. Delete Base Document
		5. Get Base Document Items
		6. Delete Base Document Items
		7. Get Sales Quotation Items
		8. Delete Sales Quotation Items
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get sales quotation
	SalesQuotation, err := s.Repo.Sale.GetSalesQuotationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete sales quotation
	err = s.Repo.Sale.DeleteSalesQuotation(req.Ctx, tx, SalesQuotation)
	if err != nil {
		return nil, err
	}

	// get base document
	baseDocument, err := s.Repo.Base.GetBaseDocumentByID(req.Ctx, tx, SalesQuotation.BaseDocumentID)
	if err != nil {
		return nil, err
	}

	// delete base document
	err = s.Repo.Base.DeleteBaseDocument(req.Ctx, tx, baseDocument)
	if err != nil {
		return nil, err
	}

	for _, item := range SalesQuotation.R.SalesQuotationItems {
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

		// delete sales quotation item
		err = s.Repo.Sale.DeleteSalesQuotationItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteSalesQuotationResponse{
		Payload: true,
	}

	return &resp, nil
}
