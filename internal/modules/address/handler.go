package address

import (
	"CustomerService/internal/common"
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

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var a Address
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	created, err := h.service.Create(&a)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, msg, &msg)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, created, common.HTTP_CREATED)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	a, err := h.service.GetByID(uint(id))
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, a, common.HTTP_OK)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.GetAll()
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, msg, &msg)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, list, common.HTTP_OK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var a Address
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	updated, err := h.service.Update(uint(id), &a)
	if err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updated, common.HTTP_OK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		msg := err.Error()
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, msg, &msg)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
