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

const SalesQuotationViewKey contextKey = "SalesQuotationView"

func SalesQuotationViewContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *sale.GetSalesQuotationViewResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := sale.NewSaleService()
			req := svc.NewGetSalesQuotationViewRequest(r.Context(), id)
			resp, err = svc.GetSalesQuotationView(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get SalesQuotationView for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get SalesQuotationView")
			return
		}
		ctx := context.WithValue(r.Context(), SalesQuotationViewKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListSalesQuotationView(w http.ResponseWriter, r *http.Request) {
	svc := sale.NewSaleService()
	req := svc.NewListSalesQuotationViewRequest(r.Context())
	resp, err := svc.ListSalesQuotationView(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list SalesQuotationView: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotationView listed successfully")
	
}

func GetSalesQuotationView(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(SalesQuotationViewKey).(*sale.GetSalesQuotationViewResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(SalesQuotationViewKey))),
			"Failed to get SalesQuotationView")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "SalesQuotationView retrieved successfully")

	
}

func SearchSalesQuotationView(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchSalesQuotationDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewSearchSalesQuotationViewRequest(r.Context(), *dto)
	resp, err := svc.SearchSalesQuotationView(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search SalesQuotationView: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "SalesQuotationView search successful")

	
}

