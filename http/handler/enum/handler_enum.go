package enum

import (
	"mvrp/domain/service/enum"
	"mvrp/errors"
	"mvrp/htresp"
	"net/http"
)

func ListEnum(w http.ResponseWriter, r *http.Request) {
	svc := enum.NewEnumService()
	req := svc.NewListEnumRequest(r.Context())
	resp, err := svc.ListEnum(req)
	if err != nil {
		htresp.RespondWithError(w, http.StatusInternalServerError,
			errors.WrapError(errors.ErrTypeService, err.Error()),
			"Failed to list Enum: "+err.Error())
		return
	}
	htresp.RespondWithJSON(w, http.StatusOK, resp, "Enum listed successfully")
}
