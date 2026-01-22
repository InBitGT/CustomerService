package address

import (
	"CustomerService/internal/common"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateAddressRequest struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type CreateAddressResponse struct {
	ID uint `json:"id"`
}

func (h *Handler) CreateInternal(w http.ResponseWriter, r *http.Request) {
	var req CreateAddressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	addr := &Address{
		Line1:      req.Line1,
		Line2:      req.Line2,
		City:       req.City,
		State:      req.State,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	}

	out, err := h.service.Create(addr)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.CreatedResponse(w, common.SUCCESS_CREATED, CreateAddressResponse{ID: out.ID}, common.HTTP_CREATED)
}

func (h *Handler) DeleteInternal(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.HardDelete(uint(id64)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, "ok", common.HTTP_OK)
}
