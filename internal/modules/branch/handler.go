package branch

import (
	"CustomerService/internal/common"
	"CustomerService/internal/middleware"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func getTenantID(r *http.Request) (uint, error) {
	claims, ok := r.Context().Value(middleware.UserCtxKey).(*middleware.UserClaims)
	if !ok || claims == nil {
		return 0, http.ErrNoCookie // cualquier error “genérico”
	}
	return claims.TenantID, nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID, err := getTenantID(r)
	if err != nil || tenantID == 0 {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	var b Branch
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	// Tenant scoping: no confiar en lo que mande el cliente
	b.TenantID = tenantID

	created, err := h.service.Create(&b)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, created, common.HTTP_CREATED)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	tenantID, err := getTenantID(r)
	if err != nil || tenantID == 0 {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	b, err := h.service.GetByID(uint(id64), tenantID)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, b, common.HTTP_OK)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	tenantID, err := getTenantID(r)
	if err != nil || tenantID == 0 {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	list, err := h.service.GetAll(tenantID)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, list, common.HTTP_OK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	tenantID, err := getTenantID(r)
	if err != nil || tenantID == 0 {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var b Branch
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	// Tenant scoping
	b.TenantID = tenantID

	updated, err := h.service.Update(uint(id64), tenantID, &b)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updated, common.HTTP_OK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	tenantID, err := getTenantID(r)
	if err != nil || tenantID == 0 {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.Delete(uint(id64), tenantID); err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
