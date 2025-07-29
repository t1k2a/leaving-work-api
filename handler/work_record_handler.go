package handler

import (
	"encoding/json"
	"net/http"

	"leaving-work-api/service"
)

type WorkRecordHandler struct {
	service service.WorkRecordService
}

func NewWorkRecordHandler(s service.WorkRecordService) *WorkRecordHandler {
	return &WorkRecordHandler{service: s}
}

func (h *WorkRecordHandler) GetWorkRecords(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	
	if userID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	records := h.service.GetRecordsByUserID(userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}