package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"leaving-work-api/model"
	"leaving-work-api/service"
)

type WorkRecordHandler struct {
	service service.WorkRecordService
	validator *validator.Validate
}

func NewWorkRecordHandler(s service.WorkRecordService) *WorkRecordHandler {
	return &WorkRecordHandler{
		service: s,
		validator: validator.New(),
	}
}

func (h *WorkRecordHandler) GetWorkRecords(w http.ResponseWriter, r *http.Request) {
	query := model.WorkRecordQuery {
		UserID: r.URL.Query().Get("user_id"),
	}
	
	// バリデーション実行
	if err := h.validator.Struct(query); err != nil {
		http.Error(w, "Invalid or  missing user_id (required, alphanum)", http.StatusBadRequest)
		return
	}

	records := h.service.GetRecordsByUserID(query.UserID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}