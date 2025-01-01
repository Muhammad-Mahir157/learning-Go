package models

import "time"

type TimeLogResponse struct {
	TimeLogId     string    `json:"timeLogId"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	LoggedAt      time.Time `json:"loggedAt"`
	TotalDuration string    `json:"totalDuration"`
	Description   string    `json:"description"`
}
