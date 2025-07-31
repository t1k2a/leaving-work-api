package repository

import (
	"leaving-work-api/model"
	"leaving-work-api/db"
)

type WorkRecordRepository interface {
	FindByUserID(userID string) []model.WorkRecord
}

type workRecordRepository  struct{}

func NewWorkRecordRepository() WorkRecordRepository {
	return &workRecordRepository{}
}

func (r *workRecordRepository) FindByUserID(userID string) []model.WorkRecord {
	var records []model.WorkRecord
	db.DB.Where("user_id = ?", userID).Find(&records)
	return records
}