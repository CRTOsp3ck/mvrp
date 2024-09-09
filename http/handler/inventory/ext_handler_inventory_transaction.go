package inventory

import (
	"encoding/json"
	"mvrp/domain/dto"
	"mvrp/domain/service/inventory"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
)

func SearchAllInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	var dto *dto.SearchInventoryTransactionDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		htresp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body: "+err.Error())
		return
	}
	svc := inventory.NewInventoryService()
	req := svc.NewSearchAllInventoryTransactionRequest(r.Context(), *dto)
	resp, err := svc.SearchAllInventoryTransaction(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to search InventoryTransaction: "+err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "InventoryTransaction search successful")
}
