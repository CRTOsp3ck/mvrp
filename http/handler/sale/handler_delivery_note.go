// Code generated by MVRP Codegen Util. DO NOT EDIT.

package sale

import (
	"context"
	"encoding/json"
	"fmt"
	"mvrp/domain/dto"
	"mvrp/domain/service/sale"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const DeliveryNoteKey contextKey = "DeliveryNote"

func DeliveryNoteContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *sale.GetDeliveryNoteResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := sale.NewSaleService()
			req := svc.NewGetDeliveryNoteRequest(r.Context(), id)
			resp, err = svc.GetDeliveryNote(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get DeliveryNote for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get DeliveryNote")
			return
		}
		ctx := context.WithValue(r.Context(), DeliveryNoteKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListDeliveryNote(w http.ResponseWriter, r *http.Request) {
	svc := sale.NewSaleService()
	req := svc.NewListDeliveryNoteRequest(r.Context())
	resp, err := svc.ListDeliveryNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list DeliveryNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DeliveryNote listed successfully")
	
}

func CreateDeliveryNote(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateDeliveryNoteDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewCreateDeliveryNoteRequest(r.Context(), *data)
	resp, err := svc.CreateDeliveryNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create DeliveryNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "DeliveryNote created successfully")

	
}

func GetDeliveryNote(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(DeliveryNoteKey).(*sale.GetDeliveryNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DeliveryNoteKey))),
			"Failed to get DeliveryNote")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "DeliveryNote retrieved successfully")

	
}

func UpdateDeliveryNote(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(DeliveryNoteKey).(*sale.GetDeliveryNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DeliveryNoteKey))),
			"Failed to get DeliveryNote")
		return
	}
	var data *dto.UpdateDeliveryNoteDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewUpdateDeliveryNoteRequest(r.Context(), *data)
	resp, err := svc.UpdateDeliveryNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update DeliveryNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DeliveryNote updated successfully")

	
}

func DeleteDeliveryNote(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(DeliveryNoteKey).(*sale.GetDeliveryNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DeliveryNoteKey))),
			"Failed to get DeliveryNote")
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewDeleteDeliveryNoteRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteDeliveryNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete DeliveryNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DeliveryNote deleted successfully")

	
}

func SearchDeliveryNote(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchDeliveryNoteDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewSearchDeliveryNoteRequest(r.Context(), *dto)
	resp, err := svc.SearchDeliveryNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search DeliveryNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DeliveryNote search successful")

	
}

