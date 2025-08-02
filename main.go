package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5" // ルーティングライブラリ
	"leaving-work-api/handler"
	"leaving-work-api/repository"
	"leaving-work-api/service"
	"leaving-work-api/db"
	"github.com/go-chi/cors"
)

type WorkRecord struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	ClockOutTime string `json:"clock_out_time"`
}

func main() {
	db.Init()
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// 本番向けとローカル環境でURLを変える
		AllowedOrigins: []string{"http://localhost:3000"}, // Next.jsの開発サーバー
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	// DI(依存注入)
	repo := repository.NewWorkRecordRepository()
	svc := service.NewWorkRecordService(repo)
	h := handler.NewWorkRecordHandler(svc)

	r.Get("/work_records", h.GetWorkRecords)
	r.Post("/work_records", h.CreateWorkRecord)
	
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}