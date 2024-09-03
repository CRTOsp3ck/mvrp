// Code generated by MVRP Codegen Util. DO NOT EDIT.

package invoice

import (
	"context"
	"encoding/json"
	"fmt"
	"mvrp/domain/dto"
	"mvrp/domain/service/invoice"
	"mvrp/errors"
	"mvrp/http/htresp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const InvoiceKey contextKey = "Invoice"

func InvoiceContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *invoice.GetInvoiceResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := invoice.NewInvoiceService()
			req := svc.NewGetInvoiceRequest(r.Context(), id)
			resp, err = svc.GetInvoice(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get Invoice")
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get Invoice")
			return
		}
		ctx := context.WithValue(r.Context(), InvoiceKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


func ListInvoice(w http.ResponseWriter, r *http.Request) {
	svc := invoice.NewInvoiceService()
	req := svc.NewListInvoiceRequest(r.Context())
	resp, err := svc.ListInvoice(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Invoice listed successfully")
	
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateInvoiceDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body")
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewCreateInvoiceRequest(r.Context(), *data)
	resp, err := svc.CreateInvoice(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "Invoice created successfully")

	
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(InvoiceKey).(*invoice.GetInvoiceResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InvoiceKey))),
			"Failed to get Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "Invoice retrieved successfully")

	
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(InvoiceKey).(*invoice.GetInvoiceResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InvoiceKey))),
			"Failed to get Invoice")
		return
	}
	var data *dto.UpdateInvoiceDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body")
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewUpdateInvoiceRequest(r.Context(), *data)
	resp, err := svc.UpdateInvoice(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Invoice updated successfully")

	
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(InvoiceKey).(*invoice.GetInvoiceResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(InvoiceKey))),
			"Failed to get Invoice")
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewDeleteInvoiceRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteInvoice(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Invoice deleted successfully")

	
}

func SearchInvoice(w http.ResponseWriter, r *http.Request) {
	var dto *dto.SearchInvoiceDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body")
		return
	}
	svc := invoice.NewInvoiceService()
	req := svc.NewSearchInvoiceRequest(r.Context(), *dto)
	resp, err := svc.SearchInvoice(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search Invoice")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Invoice search successful")

	
}

