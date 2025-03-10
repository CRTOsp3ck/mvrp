// Code generated by MVRP Codegen Util. DO NOT EDIT.

package invoice

import (
	"context"
	"encoding/json"
	"fmt"
	"mvrp/domain/dto"
	"mvrp/domain/service/invoice"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const DebitNoteKey contextKey = "DebitNote"

func DebitNoteContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *invoice.GetDebitNoteResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := invoice.NewInvoiceService()
			req := svc.NewGetDebitNoteRequest(r.Context(), id)
			resp, err = svc.GetDebitNote(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get DebitNote for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get DebitNote")
			return
		}
		ctx := context.WithValue(r.Context(), DebitNoteKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListDebitNote(w http.ResponseWriter, r *http.Request) {
	svc := invoice.NewInvoiceService()
	req := svc.NewListDebitNoteRequest(r.Context())
	resp, err := svc.ListDebitNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list DebitNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DebitNote listed successfully")
	
}

func CreateDebitNote(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateDebitNoteDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewCreateDebitNoteRequest(r.Context(), *data)
	resp, err := svc.CreateDebitNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create DebitNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "DebitNote created successfully")

	
}

func GetDebitNote(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(DebitNoteKey).(*invoice.GetDebitNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DebitNoteKey))),
			"Failed to get DebitNote")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "DebitNote retrieved successfully")

	
}

func UpdateDebitNote(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(DebitNoteKey).(*invoice.GetDebitNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DebitNoteKey))),
			"Failed to get DebitNote")
		return
	}
	var data *dto.UpdateDebitNoteDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewUpdateDebitNoteRequest(r.Context(), *data)
	resp, err := svc.UpdateDebitNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update DebitNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DebitNote updated successfully")

	
}

func DeleteDebitNote(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(DebitNoteKey).(*invoice.GetDebitNoteResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(DebitNoteKey))),
			"Failed to get DebitNote")
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewDeleteDebitNoteRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteDebitNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete DebitNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DebitNote deleted successfully")

	
}

func SearchDebitNote(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchDebitNoteDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewSearchDebitNoteRequest(r.Context(), *dto)
	resp, err := svc.SearchDebitNote(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search DebitNote: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "DebitNote search successful")

	
}

