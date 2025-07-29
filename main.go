package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5" // ルーティングライブラリ
	"leaving-work-api/handler"
	"leaving-work-api/repository"
	"leaving-work-api/service"
)

type WorkRecord struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	ClockOutTime string `json:"clock_out_time"`
}

func main() {
	r := chi.NewRouter()

	// DI(依存注入)
	repo := repository.NewWorkRecordRepository()
	svc := service.NewWorkRecordService(repo)
	h := handler.NewWorkRecordHandler(svc)
	r.Get("/work_records", h.GetWorkRecords)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}