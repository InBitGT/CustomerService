package user_branch

import (
	"CustomerService/internal/common"
	"CustomerService/internal/middleware"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func getTenantID(r *http.Request) (uint, bool) {
	claims, ok := r.Context().Value(middleware.UserCtxKey).(*middleware.UserClaims)
	if !ok || claims == nil || claims.TenantID == 0 {
		return 0, false
	}
	return claims.TenantID, true
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID, ok := getTenantID(r)
	if !ok {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	var ub UserBranch
	if err := json.NewDecoder(r.Body).Decode(&ub); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	created, err := h.service.Create(tenantID, &ub)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := "branch not found for tenant"
			common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, &msg)
			return
		}
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, created, common.HTTP_CREATED)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	tenantID, ok := getTenantID(r)
	if !ok {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	ub, err := h.service.GetByID(tenantID, uint(id64))
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, ub, common.HTTP_OK)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	tenantID, ok := getTenantID(r)
	if !ok {
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
	tenantID, ok := getTenantID(r)
	if !ok {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var ub UserBranch
	if err := json.NewDecoder(r.Body).Decode(&ub); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	updated, err := h.service.Update(tenantID, uint(id64), &ub)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := "record not found for tenant"
			common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, &msg)
			return
		}
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updated, common.HTTP_OK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	tenantID, ok := getTenantID(r)
	if !ok {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.Delete(tenantID, uint(id64)); err != nil {
		if err == gorm.ErrRecordNotFound {
			msg := "record not found for tenant"
			common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, &msg)
			return
		}
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
