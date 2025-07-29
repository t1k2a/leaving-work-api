package model

type WorkRecord struct {
	ID int `json:"id"`
	UserID string `json:"user_id"`
	ClockOutTime string `json:"clock_out_time"`
}

