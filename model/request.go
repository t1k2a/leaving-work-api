package model

type WorkRecordQuery struct {
	UserID string `json:"user_id" validate:"required,alphanum"`
}