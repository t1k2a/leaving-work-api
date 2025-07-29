package repository

import (
	"leaving-work-api/model"
	"leaving-work-api/mock"
)

type WorkRecordRepository interface {
	FindByUserID(userID string) []model.WorkRecord
}

type workRecordRepository  struct{}

func NewWorkRecordRepository() WorkRecordRepository {
	return &workRecordRepository{}
}

func (r *workRecordRepository) FindByUserID(userID string) []model.WorkRecord {
	return mock.GetMockWorkRecords(userID)
}