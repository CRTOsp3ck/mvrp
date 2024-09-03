package entity

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"mvrp/domain/dto"
// 	"mvrp/domain/service/entity"
// 	"mvrp/http/resp"
// 	"net/http"
// 	"strconv"

// 	"github.com/go-chi/chi/v5"
// )

// // Custom type for the context key
// const EntityKey contextKey = "Entity"

// func EntityContext(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if idStr := chi.URLParam(r, "id"); idStr != "" {
// 			id, err := strconv.Atoi(idStr)
// 			if err != nil {

// 				resp.RespondWithError(w, http.StatusBadRequest, err, "Failed to convert ID to integer")
// 				return
// 			}
// 			svc := entity.NewEntityService()
// 			req := entity.GetEntityRequest{}
// 			resp := svc.GetEntity(req)
// 			if resp.Error != nil {
// 				resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to get Entity")
// 				return
// 			}
// 			result = resp.StructPayload
// 		} else {
// 			resp.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("status bad request"), "Invalid ID")
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), EntityKey, result)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func SearchEntity(w http.ResponseWriter, r *http.Request) {
// 	var qp *qto.EntityQTO
// 	err := json.NewDecoder(r.Body).Decode(&qp)
// 	if err != nil {
// 		resp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body")
// 		return
// 	}
// 	svc := entity.NewEntityService()
// 	req := svc.NewEntityRequest(r.Context(), nil, nil, qp)
// 	resp := svc.SearchEntity(req)
// 	if resp.Error != nil {
// 		resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to search Entity")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusOK, resp.StructSlicePayload, "Entity search completed successfully")
// }

// func GetEntity(w http.ResponseWriter, r *http.Request) {
// 	value, ok := r.Context().Value(EntityKey).(*dto.EntityDTO)
// 	if !ok {
// 		resp.RespondWithError(w, http.StatusInternalServerError, fmt.Errorf("payload is not in the right format, should be %T", r.Context().Value(EntityKey)), "Invalid Entity result")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusOK, value, "Entity retrieved successfully")
// }

// func CreateEntity(w http.ResponseWriter, r *http.Request) {
// 	var data *dto.EntityDTO
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		resp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body")
// 		return
// 	}
// 	svc := entity.NewEntityService()
// 	req := svc.NewEntityRequest(r.Context(), data, nil, nil)
// 	resp := svc.CreateEntity(req)
// 	if resp.Error != nil {
// 		resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to create Entity")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusCreated, resp.StructPayload, "Entity created successfully")
// }

// func UpdateEntity(w http.ResponseWriter, r *http.Request) {
// 	_, ok := r.Context().Value(EntityKey).(*dto.EntityDTO)
// 	if !ok {
// 		resp.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("unable to complete request"), "Invalid Entity request")
// 		return
// 	}
// 	var data *dto.EntityDTO
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		resp.RespondWithError(w, http.StatusBadRequest, err, "Failed to decode request body")
// 		return
// 	}
// 	svc := entity.NewEntityService()
// 	req := svc.NewEntityRequest(r.Context(), data, nil, nil)
// 	resp := svc.UpdateEntity(req)
// 	if resp.Error != nil {
// 		resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to update Entity")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusOK, resp.StructPayload, "Entity updated successfully")
// }

// func DeleteEntity(w http.ResponseWriter, r *http.Request) {
// 	v, ok := r.Context().Value(EntityKey).(*dto.EntityDTO)
// 	if !ok {
// 		resp.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("unable to complete request"), "Invalid Entity request")
// 		return
// 	}
// 	svc := entity.NewEntityService()
// 	req := svc.NewEntityRequest(r.Context(), v, nil, nil)
// 	resp := svc.DeleteEntity(req)
// 	if resp.Error != nil {
// 		resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to delete Entity")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusOK, nil, "Entity deleted successfully")
// }

// func ListEntities(w http.ResponseWriter, r *http.Request) {
// 	svc := entity.NewEntityService()
// 	req := svc.NewEntityRequest(r.Context(), nil, nil, nil)
// 	resp := svc.ListAllEntities(req)
// 	if resp.Error != nil {
// 		resp.RespondWithError(w, http.StatusInternalServerError, resp.Error, "Failed to list Entities")
// 		return
// 	}
// 	resp.RespondWithJSON(w, http.StatusOK, resp.StructSlicePayload, "Entities listed successfully")
// }
