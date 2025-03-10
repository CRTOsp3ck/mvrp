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

const SalesQuotationKey contextKey = "SalesQuotation"

func SalesQuotationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *sale.GetSalesQuotationResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := sale.NewSaleService()
			req := svc.NewGetSalesQuotationRequest(r.Context(), id)
			resp, err = svc.GetSalesQuotation(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get SalesQuotation for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get SalesQuotation")
			return
		}
		ctx := context.WithValue(r.Context(), SalesQuotationKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListSalesQuotation(w http.ResponseWriter, r *http.Request) {
	svc := sale.NewSaleService()
	req := svc.NewListSalesQuotationRequest(r.Context())
	resp, err := svc.ListSalesQuotation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list SalesQuotation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotation listed successfully")
	
}

func CreateSalesQuotation(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateSalesQuotationDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewCreateSalesQuotationRequest(r.Context(), *data)
	resp, err := svc.CreateSalesQuotation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create SalesQuotation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "SalesQuotation created successfully")

	
}

func GetSalesQuotation(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(SalesQuotationKey).(*sale.GetSalesQuotationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(SalesQuotationKey))),
			"Failed to get SalesQuotation")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "SalesQuotation retrieved successfully")

	
}

func UpdateSalesQuotation(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(SalesQuotationKey).(*sale.GetSalesQuotationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(SalesQuotationKey))),
			"Failed to get SalesQuotation")
		return
	}
	var data *dto.UpdateSalesQuotationDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewUpdateSalesQuotationRequest(r.Context(), *data)
	resp, err := svc.UpdateSalesQuotation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update SalesQuotation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotation updated successfully")

	
}

func DeleteSalesQuotation(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(SalesQuotationKey).(*sale.GetSalesQuotationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(SalesQuotationKey))),
			"Failed to get SalesQuotation")
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewDeleteSalesQuotationRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteSalesQuotation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete SalesQuotation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotation deleted successfully")

	
}

func SearchSalesQuotation(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchSalesQuotationDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewSearchSalesQuotationRequest(r.Context(), *dto)
	resp, err := svc.SearchSalesQuotation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search SalesQuotation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotation search successful")

	
}

