package common

import (
	"time"

	"github.com/google/uuid"
)

type TimeLogResponse struct {
	TimeLogId     uuid.UUID
	StartTime     time.Time
	EndTime       time.Time
	LoggedAt      time.Time
	TotalDuration string
	Description   string
}
