package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5" // ルーティングライブラリ
)

type WorkRecord struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	ClockOutTime string `json:"ckick_out_time"`
}

func main() {
	r := chi.NewRouter()

	// GET /work_records?user_id=xxxに対応
	r.Get("/work_records", func(w http.ResponseWriter, r* http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id is required", http.StatusBadRequest)
			return
		}

		// モックデータ（学習目的の仮データ）
		mock := []WorkRecord {
			{ID: 1, UserID: userID, ClockOutTime: "2024-07-01T18:30:00Z"},
			{ID: 2, UserID: userID, ClockOutTime: "2024-07-02T18:45:00Z"},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mock)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}