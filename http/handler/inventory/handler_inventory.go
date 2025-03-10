// Code generated by MVRP Codegen Util. DO NOT EDIT.

package inventory

import (
	"context"
	"encoding/json"
	"fmt"
	"mvrp/domain/dto"
	"mvrp/domain/service/inventory"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const InventoryKey contextKey = "Inventory"

func InventoryContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *inventory.GetInventoryResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := inventory.NewInventoryService()
			req := svc.NewGetInventoryRequest(r.Context(), id)
			resp, err = svc.GetInventory(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get Inventory for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get Inventory")
			return
		}
		ctx := context.WithValue(r.Context(), InventoryKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListInventory(w http.ResponseWriter, r *http.Request) {
	svc := inventory.NewInventoryService()
	req := svc.NewListInventoryRequest(r.Context())
	resp, err := svc.ListInventory(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list Inventory: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Inventory listed successfully")
	
}

func CreateInventory(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateInventoryDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewCreateInventoryRequest(r.Context(), *data)
	resp, err := svc.CreateInventory(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create Inventory: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "Inventory created successfully")

	
}

func GetInventory(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(InventoryKey).(*inventory.GetInventoryResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryKey))),
			"Failed to get Inventory")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "Inventory retrieved successfully")

	
}

func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(InventoryKey).(*inventory.GetInventoryResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryKey))),
			"Failed to get Inventory")
		return
	}
	var data *dto.UpdateInventoryDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewUpdateInventoryRequest(r.Context(), *data)
	resp, err := svc.UpdateInventory(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update Inventory: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Inventory updated successfully")

	
}

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(InventoryKey).(*inventory.GetInventoryResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryKey))),
			"Failed to get Inventory")
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewDeleteInventoryRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteInventory(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete Inventory: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Inventory deleted successfully")

	
}

func SearchInventory(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchInventoryDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewSearchInventoryRequest(r.Context(), *dto)
	resp, err := svc.SearchInventory(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search Inventory: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Inventory search successful")

	
}

