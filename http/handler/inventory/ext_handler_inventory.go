package inventory

import (
	"mvrp/domain/service/inventory"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetInventoryExistsByItemID(w http.ResponseWriter, r *http.Request) {
	var resp *inventory.GetInventoryExistsByItemIDResponse
	if idStr := chi.URLParam(r, "id"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			htresp.RespondWithError(w, http.StatusBadRequest,
				errors.WrapError(errors.ErrTypeConversion, "ID must be an integer"),
				"Failed to convert ID to integer")
			return
		}
		svc := inventory.NewInventoryService()
		req := svc.NewGetInventoryExistsByItemIDRequest(r.Context(), id)
		resp, err = svc.GetInventoryExistsByItemID(req)
		if err != nil {
			htresp.RespondWithError(w, http.StatusInternalServerError,
				errors.WrapError(errors.ErrTypeService, err.Error()),
				"Failed to get Inventory")
			return
		}
	} else {
		htresp.RespondWithError(w, http.StatusBadRequest,
			errors.WrapError(errors.ErrTypeBadRequest, "ID is required"),
			"Failed to get Inventory")
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Inventory exists retrieved successfully")
}
