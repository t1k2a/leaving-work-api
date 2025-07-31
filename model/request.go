package model

type WorkRecordQuery struct {
	UserID string `json:"user_id" validate:"required,alphanum"`
}

type CreateWorkRecordRequest struct {
	UserID string `json:"user_id" validate:"required,alphanum"`
	ClockOutTime string `json:"clock_out_time calidate:"required,datetime=2025-07-31T15:04:05Z07:00"`
}