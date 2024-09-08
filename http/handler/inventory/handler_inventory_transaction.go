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

const InventoryTransactionKey contextKey = "InventoryTransaction"

func InventoryTransactionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *inventory.GetInventoryTransactionResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := inventory.NewInventoryService()
			req := svc.NewGetInventoryTransactionRequest(r.Context(), id)
			resp, err = svc.GetInventoryTransaction(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get InventoryTransaction")
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get InventoryTransaction")
			return
		}
		ctx := context.WithValue(r.Context(), InventoryTransactionKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	svc := inventory.NewInventoryService()
	req := svc.NewListInventoryTransactionRequest(r.Context())
	resp, err := svc.ListInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "InventoryTransaction listed successfully")
	
}

func CreateInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateInventoryTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body")
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewCreateInventoryTransactionRequest(r.Context(), *data)
	resp, err := svc.CreateInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "InventoryTransaction created successfully")

	
}

func GetInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(InventoryTransactionKey).(*inventory.GetInventoryTransactionResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryTransactionKey))),
			"Failed to get InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "InventoryTransaction retrieved successfully")

	
}

func UpdateInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(InventoryTransactionKey).(*inventory.GetInventoryTransactionResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryTransactionKey))),
			"Failed to get InventoryTransaction")
		return
	}
	var data *dto.UpdateInventoryTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body")
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewUpdateInventoryTransactionRequest(r.Context(), *data)
	resp, err := svc.UpdateInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "InventoryTransaction updated successfully")

	
}

func DeleteInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(InventoryTransactionKey).(*inventory.GetInventoryTransactionResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InventoryTransactionKey))),
			"Failed to get InventoryTransaction")
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewDeleteInventoryTransactionRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "InventoryTransaction deleted successfully")

	
}

func SearchInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchInventoryTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body")
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewSearchInventoryTransactionRequest(r.Context(), *dto)
	resp, err := svc.SearchInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search InventoryTransaction")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "InventoryTransaction search successful")

	
}

