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

const OrderConfirmationKey contextKey = "OrderConfirmation"

func OrderConfirmationContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var resp *sale.GetOrderConfirmationResponse
		if idStr := chi.URLParam(r, "id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				htresp.RespondWithError(w, http.StatusBadRequest,
					errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
					"Failed to convert ID to integer")
				return
			}
			svc := sale.NewSaleService()
			req := svc.NewGetOrderConfirmationRequest(r.Context(), id)
			resp, err = svc.GetOrderConfirmation(req)
			if err != nil {
				htresp.RespondWithError(w, http.StatusInternalServerError,
					errors.WrapError(errors.ErrTypeService, err.Error()),
					"Failed to get OrderConfirmation for context: " + err.Error())
				return
			}
		} else {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
				"Failed to get OrderConfirmation")
			return
		}
		ctx := context.WithValue(r.Context(), OrderConfirmationKey, resp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}



func ListOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	svc := sale.NewSaleService()
	req := svc.NewListOrderConfirmationRequest(r.Context())
	resp, err := svc.ListOrderConfirmation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list OrderConfirmation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "OrderConfirmation listed successfully")
	
}

func CreateOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateOrderConfirmationDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewCreateOrderConfirmationRequest(r.Context(), *data)
	resp, err := svc.CreateOrderConfirmation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to create OrderConfirmation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusCreated, resp, "OrderConfirmation created successfully")

	
}

func GetOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value(OrderConfirmationKey).(*sale.GetOrderConfirmationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(OrderConfirmationKey))),
			"Failed to get OrderConfirmation")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, value, "OrderConfirmation retrieved successfully")

	
}

func UpdateOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value(OrderConfirmationKey).(*sale.GetOrderConfirmationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(OrderConfirmationKey))),
			"Failed to get OrderConfirmation")
		return
	}
	var data *dto.UpdateOrderConfirmationDTO
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeDecoding, err.Error()),
			"Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewUpdateOrderConfirmationRequest(r.Context(), *data)
	resp, err := svc.UpdateOrderConfirmation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to update OrderConfirmation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "OrderConfirmation updated successfully")

	
}

func DeleteOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	v, ok := r.Context().Value(OrderConfirmationKey).(*sale.GetOrderConfirmationResponse)
	if !ok {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeAssertion,
				fmt.Sprintf("Incorrect response format, should be %T", r.Context().Value(OrderConfirmationKey))),
			"Failed to get OrderConfirmation")
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewDeleteOrderConfirmationRequest(r.Context(), v.Payload.ID)
	resp, err := svc.DeleteOrderConfirmation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to delete OrderConfirmation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "OrderConfirmation deleted successfully")

	
}

func SearchOrderConfirmation(w http.ResponseWriter, r *http.Request) {
	
	var dto *dto.SearchOrderConfirmationDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: " + err.Error())
		return
	}
	svc := sale.NewSaleService()
	req := svc.NewSearchOrderConfirmationRequest(r.Context(), *dto)
	resp, err := svc.SearchOrderConfirmation(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search OrderConfirmation: " + err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "OrderConfirmation search successful")

	
}

