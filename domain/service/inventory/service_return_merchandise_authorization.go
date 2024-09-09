package inventory

import (
	"context"
	"mvrp/data/model/inventory"
	"mvrp/domain/dto"
	"mvrp/domain/proc"
	"mvrp/util"

	"github.com/ericlagergren/decimal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// LIST RETURN MERCHANDISE AUTHORIZATION
type ListReturnMerchandiseAuthorizationRequest struct {
	Ctx context.Context
}

func (s *InventoryService) NewListReturnMerchandiseAuthorizationRequest(ctx context.Context) *ListReturnMerchandiseAuthorizationRequest {
	return &ListReturnMerchandiseAuthorizationRequest{
		Ctx: ctx,
	}
}

type ListReturnMerchandiseAuthorizationResponse struct {
	Payload inventory.ReturnMerchandiseAuthorizationSlice `json:"payload"`
}

func (s *InventoryService) NewListReturnMerchandiseAuthorizationResponse(payload inventory.ReturnMerchandiseAuthorizationSlice) *ListReturnMerchandiseAuthorizationResponse {
	return &ListReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) ListReturnMerchandiseAuthorization(req *ListReturnMerchandiseAuthorizationRequest) (*ListReturnMerchandiseAuthorizationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.ListAllReturnMerchandiseAuthorizations(req.Ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	resp := ListReturnMerchandiseAuthorizationResponse{
		Payload: res,
	}
	return &resp, nil
}

// SEARCH RETURN MERCHANDISE AUTHORIZATION
type SearchReturnMerchandiseAuthorizationRequest struct {
	Ctx     context.Context
	Payload dto.SearchReturnMerchandiseAuthorizationDTO
}

func (s *InventoryService) NewSearchReturnMerchandiseAuthorizationRequest(ctx context.Context, payload dto.SearchReturnMerchandiseAuthorizationDTO) *SearchReturnMerchandiseAuthorizationRequest {
	return &SearchReturnMerchandiseAuthorizationRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type SearchReturnMerchandiseAuthorizationResponse struct {
	Payload    inventory.ReturnMerchandiseAuthorizationSlice `json:"payload"`
	Pagination dto.PaginationDTO                             `json:"pagination"`
}

func (s *InventoryService) NewSearchReturnMerchandiseAuthorizationResponse(payload inventory.ReturnMerchandiseAuthorizationSlice) *SearchReturnMerchandiseAuthorizationResponse {
	return &SearchReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) SearchReturnMerchandiseAuthorization(req *SearchReturnMerchandiseAuthorizationRequest) (*SearchReturnMerchandiseAuthorizationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	res, err := s.Repo.Inventory.SearchReturnMerchandiseAuthorizations(req.Ctx, tx, req.Payload)
	if err != nil {
		return nil, err
	}

	// Pagination
	totalCount, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationTotalCount(req.Ctx, tx)
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
	resp := SearchReturnMerchandiseAuthorizationResponse{
		Payload:    res,
		Pagination: pd,
	}
	return &resp, nil
}

// GET RETURN MERCHANDISE AUTHORIZATION
type GetReturnMerchandiseAuthorizationRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewGetReturnMerchandiseAuthorizationRequest(ctx context.Context, id int) *GetReturnMerchandiseAuthorizationRequest {
	return &GetReturnMerchandiseAuthorizationRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type GetReturnMerchandiseAuthorizationResponse struct {
	Payload dto.GetReturnMerchandiseAuthorizationDTO `json:"payload"`
}

func (s *InventoryService) NewGetReturnMerchandiseAuthorizationResponse(payload dto.GetReturnMerchandiseAuthorizationDTO) *GetReturnMerchandiseAuthorizationResponse {
	return &GetReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) GetReturnMerchandiseAuthorization(req *GetReturnMerchandiseAuthorizationRequest) (*GetReturnMerchandiseAuthorizationResponse, error) {
	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	ginRes, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// get return merchandise authorization items
	ginItems, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	ginItemRes := make([]dto.GetReturnMerchandiseAuthorizationItemDTO, 0)
	for _, item := range ginItems {
		ginItemRes = append(ginItemRes, dto.GetReturnMerchandiseAuthorizationItemDTO{
			ReturnMerchandiseAuthorizationItem: *item,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := GetReturnMerchandiseAuthorizationResponse{
		Payload: dto.GetReturnMerchandiseAuthorizationDTO{
			ReturnMerchandiseAuthorization: *ginRes,
			Items:                          ginItemRes,
		},
	}
	return &resp, nil
}

// CREATE RETURN MERCHANDISE AUTHORIZATION
type CreateReturnMerchandiseAuthorizationRequest struct {
	Ctx     context.Context
	Payload dto.CreateReturnMerchandiseAuthorizationDTO
}

func (s *InventoryService) NewCreateReturnMerchandiseAuthorizationRequest(ctx context.Context, payload dto.CreateReturnMerchandiseAuthorizationDTO) *CreateReturnMerchandiseAuthorizationRequest {
	return &CreateReturnMerchandiseAuthorizationRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type CreateReturnMerchandiseAuthorizationResponse struct {
	Payload inventory.ReturnMerchandiseAuthorization `json:"payload"`
}

func (s *InventoryService) NewCreateReturnMerchandiseAuthorizationResponse(payload inventory.ReturnMerchandiseAuthorization) *CreateReturnMerchandiseAuthorizationResponse {
	return &CreateReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) CreateReturnMerchandiseAuthorization(req *CreateReturnMerchandiseAuthorizationRequest) (*CreateReturnMerchandiseAuthorizationResponse, error) {
	/*
		1. Create ReturnMerchandiseAuthorization
		2. Create ReturnMerchandiseAuthorizationItems
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// create return merchandise authorization
	if req.Payload.ReturnMerchandiseAuthorization.RmaNumber == "" {
		nextID, err := s.Repo.Inventory.GetNextEntryReturnMerchandiseAuthorizationID(req.Ctx, tx)
		if err != nil {
			return nil, err
		}
		req.Payload.ReturnMerchandiseAuthorization.RmaNumber = util.Util.Str.ToString(nextID)
	}
	var ginItems []*inventory.ReturnMerchandiseAuthorizationItem
	for _, item := range req.Payload.Items {
		ginItems = append(ginItems, &item.ReturnMerchandiseAuthorizationItem)
	}
	err = proc.ProcessReturnMerchandiseAuthorizationAmounts(&req.Payload.ReturnMerchandiseAuthorization, ginItems)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.CreateReturnMerchandiseAuthorization(req.Ctx, tx, &req.Payload.ReturnMerchandiseAuthorization)
	if err != nil {
		return nil, err
	}

	// create return merchandise authorization items
	for _, item := range req.Payload.Items {
		item.ReturnMerchandiseAuthorizationItem.RmaID = null.IntFrom(req.Payload.ReturnMerchandiseAuthorization.ID)
		err = proc.ProcessReturnMerchandiseAuthorizationItemAmounts(&item.ReturnMerchandiseAuthorizationItem)
		if err != nil {
			return nil, err
		}
		err = s.Repo.Inventory.CreateReturnMerchandiseAuthorizationItem(req.Ctx, tx, &item.ReturnMerchandiseAuthorizationItem)
		if err != nil {
			return nil, err
		}

		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.Quantity.Big)
		err = proc.ProcessInventoryAmounts(inv)
		if err != nil {
			return nil, err
		}
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeReturn,
			Quantity:        item.Quantity,
			Reason:          null.StringFrom("Return Merchandise Authorization Creation"),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

	}

	// get created return merchandise authorization
	gin, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.Payload.ReturnMerchandiseAuthorization.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := CreateReturnMerchandiseAuthorizationResponse{
		Payload: *gin,
	}

	return &resp, nil
}

// UPDATE RETURN MERCHANDISE AUTHORIZATION
type UpdateReturnMerchandiseAuthorizationRequest struct {
	Ctx     context.Context
	Payload dto.UpdateReturnMerchandiseAuthorizationDTO
}

func (s *InventoryService) NewUpdateReturnMerchandiseAuthorizationRequest(ctx context.Context, payload dto.UpdateReturnMerchandiseAuthorizationDTO) *UpdateReturnMerchandiseAuthorizationRequest {
	return &UpdateReturnMerchandiseAuthorizationRequest{
		Ctx:     ctx,
		Payload: payload,
	}
}

type UpdateReturnMerchandiseAuthorizationResponse struct {
	Payload inventory.ReturnMerchandiseAuthorization `json:"payload"`
}

func (s *InventoryService) NewUpdateReturnMerchandiseAuthorizationResponse(payload inventory.ReturnMerchandiseAuthorization) *UpdateReturnMerchandiseAuthorizationResponse {
	return &UpdateReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) UpdateReturnMerchandiseAuthorization(req *UpdateReturnMerchandiseAuthorizationRequest) (*UpdateReturnMerchandiseAuthorizationResponse, error) {
	/*
		1. Update ReturnMerchandiseAuthorization
		2. Update ReturnMerchandiseAuthorization Items
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	currGin, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.Payload.ReturnMerchandiseAuthorization.ID)
	if err != nil {
		return nil, err
	}

	// update return merchandise authorization
	var ginItems []*inventory.ReturnMerchandiseAuthorizationItem
	for _, item := range req.Payload.Items {
		ginItems = append(ginItems, &item.ReturnMerchandiseAuthorizationItem)
	}
	err = proc.ProcessReturnMerchandiseAuthorizationAmounts(&req.Payload.ReturnMerchandiseAuthorization, ginItems)
	if err != nil {
		return nil, err
	}
	err = s.Repo.Inventory.UpdateReturnMerchandiseAuthorization(req.Ctx, tx, &req.Payload.ReturnMerchandiseAuthorization)
	if err != nil {
		return nil, err
	}

	// delete the ones that are in the current list and not in the new list
	currGinItems, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(req.Ctx, tx, currGin.ID)
	if err != nil {
		return nil, err
	}
	for _, currGinItem := range currGinItems {
		found := false
		for _, item := range req.Payload.Items {
			if currGinItem.ID == item.ID {
				found = true
				break
			}
		}
		if !found {
			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, currGinItem.InventoryID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, currGinItem.Quantity.Big)
			err = proc.ProcessInventoryAmounts(inv)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     currGinItem.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeReturnCancellation,
				Quantity:        currGinItem.Quantity,
				Reason:          null.StringFrom("Return Merchandise Authorization Adjustment"),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}

			// delete return merchandise authorization item
			err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorizationItem(req.Ctx, tx, currGinItem)
			if err != nil {
				return nil, err
			}
		}
	}

	// create or update return merchandise authorization items
	for _, item := range req.Payload.Items {
		// check if the item is new
		itemExists, err := s.Repo.Inventory.ReturnMerchandiseAuthorizationItemExists(req.Ctx, tx, item.ID)
		if err != nil {
			return nil, err
		}

		if itemExists {
			currGinItem, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemByID(req.Ctx, tx, item.ID)
			if err != nil {
				return nil, err
			}
			amountOffset := types.NewNullDecimal(decimal.New(0, 2))
			amountOffset.Sub(item.Quantity.Big, currGinItem.Quantity.Big)

			// update return merchandise authorization item
			err = proc.ProcessReturnMerchandiseAuthorizationItemAmounts(&item.ReturnMerchandiseAuthorizationItem)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateReturnMerchandiseAuthorizationItem(req.Ctx, tx, &item.ReturnMerchandiseAuthorizationItem)
			if err != nil {
				return nil, err
			}

			quantityChanged := amountOffset.Big.Cmp(decimal.New(0, 2)) != 0
			if quantityChanged {
				// update inventory
				inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
				if err != nil {
					return nil, err
				}
				inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, amountOffset.Big)
				err = proc.ProcessInventoryAmounts(inv)
				if err != nil {
					return nil, err
				}
				err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
				if err != nil {
					return nil, err
				}

				// create inventory transaction
				invTx := &inventory.InventoryTransaction{
					InventoryID:     item.InventoryID,
					TransactionType: inventory.InventoryTransactionTypeReturnAdjustment,
					Quantity:        types.Decimal(amountOffset),
					Reason:          null.StringFrom("Return Merchandise Authorization Adjustment"),
				}
				err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
				if err != nil {
					return nil, err
				}
			}
		} else {
			// create return merchandise authorization item
			err = proc.ProcessReturnMerchandiseAuthorizationItemAmounts(&item.ReturnMerchandiseAuthorizationItem)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.CreateReturnMerchandiseAuthorizationItem(req.Ctx, tx, &item.ReturnMerchandiseAuthorizationItem)
			if err != nil {
				return nil, err
			}

			// update inventory
			inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
			if err != nil {
				return nil, err
			}
			inv.QuantityAvailable.Sub(inv.QuantityAvailable.Big, item.Quantity.Big)
			err = proc.ProcessInventoryAmounts(inv)
			if err != nil {
				return nil, err
			}
			err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
			if err != nil {
				return nil, err
			}

			// create inventory transaction
			invTx := &inventory.InventoryTransaction{
				InventoryID:     item.InventoryID,
				TransactionType: inventory.InventoryTransactionTypeReturn,
				Quantity:        item.Quantity,
				Reason:          null.StringFrom("Return Merchandise Authorization Adjustment"),
			}
			err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
			if err != nil {
				return nil, err
			}
		}
	}

	// get updated return merchandise authorization
	gin, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.Payload.ReturnMerchandiseAuthorization.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := UpdateReturnMerchandiseAuthorizationResponse{
		Payload: *gin,
	}

	return &resp, nil
}

// DELETE RETURN MERCHANDISE AUTHORIZATION
type DeleteReturnMerchandiseAuthorizationRequest struct {
	Ctx context.Context
	ID  int
}

func (s *InventoryService) NewDeleteReturnMerchandiseAuthorizationRequest(ctx context.Context, id int) *DeleteReturnMerchandiseAuthorizationRequest {
	return &DeleteReturnMerchandiseAuthorizationRequest{
		Ctx: ctx,
		ID:  id,
	}
}

type DeleteReturnMerchandiseAuthorizationResponse struct {
	Payload bool `json:"payload"`
}

func (s *InventoryService) NewDeleteReturnMerchandiseAuthorizationResponse(payload bool) *DeleteReturnMerchandiseAuthorizationResponse {
	return &DeleteReturnMerchandiseAuthorizationResponse{
		Payload: payload,
	}
}

func (s *InventoryService) DeleteReturnMerchandiseAuthorization(req *DeleteReturnMerchandiseAuthorizationRequest) (*DeleteReturnMerchandiseAuthorizationResponse, error) {
	/*
		1. Delete ReturnMerchandiseAuthorization
		2. Delete ReturnMerchandiseAuthorizationItems
		3. Update Inventory
		4. Create InventoryTransaction
	*/

	tx, err := s.Repo.Begin(req.Ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// get return merchandise authorization
	gin, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationByID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}

	// delete return merchandise authorization
	err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorization(req.Ctx, tx, gin)
	if err != nil {
		return nil, err
	}

	// delete return merchandise authorization items
	ginItems, err := s.Repo.Inventory.GetReturnMerchandiseAuthorizationItemsByReturnMerchandiseAuthorizationID(req.Ctx, tx, req.ID)
	if err != nil {
		return nil, err
	}
	for _, item := range ginItems {
		// update inventory
		inv, err := s.Repo.Inventory.GetInventoryByID(req.Ctx, tx, item.InventoryID.Int)
		if err != nil {
			return nil, err
		}
		inv.QuantityAvailable.Add(inv.QuantityAvailable.Big, item.Quantity.Big)
		err = proc.ProcessInventoryAmounts(inv)
		if err != nil {
			return nil, err
		}
		err = s.Repo.Inventory.UpdateInventory(req.Ctx, tx, inv)
		if err != nil {
			return nil, err
		}

		// create inventory transaction
		invTx := &inventory.InventoryTransaction{
			InventoryID:     item.InventoryID,
			TransactionType: inventory.InventoryTransactionTypeReturnCancellation,
			Quantity:        item.Quantity,
			Reason:          null.StringFrom("Return Merchandise Authorization Cancellation"),
		}
		err = s.Repo.Inventory.CreateInventoryTransaction(req.Ctx, tx, invTx)
		if err != nil {
			return nil, err
		}

		// delete return merchandise authorization item
		err = s.Repo.Inventory.DeleteReturnMerchandiseAuthorizationItem(req.Ctx, tx, item)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	resp := DeleteReturnMerchandiseAuthorizationResponse{
		Payload: true,
	}

	return &resp, nil
}
