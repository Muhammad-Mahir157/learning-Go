package response

import "time"

type TimeLogResponse struct {
	TimeLogId     string
	StartTime     time.Time
	EndTime       time.Time
	LoggedAt      time.Time
	TotalDuration string
	Description   string
}
