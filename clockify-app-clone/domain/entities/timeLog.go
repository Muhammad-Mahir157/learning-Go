package entities

import (
	"time"

	"github.com/google/uuid"
)

type TimeLog struct {
	TimeLogId     uuid.UUID
	StartTime     time.Time
	EndTime       time.Time
	LoggedAt      time.Time
	TotalDuration string
	Description   string
}

func NewTimeLogEntity(startTime time.Time, endTime time.Time, description string) *TimeLog {
	return &TimeLog{
		TimeLogId:     uuid.New(),
		StartTime:     startTime,
		EndTime:       endTime,
		LoggedAt:      time.Now(),
		TotalDuration: endTime.Sub(startTime).String(),
		Description:   description,
	}
}

func ExistingTimeLogEntity(timeLogId uuid.UUID, startTime time.Time, endTime time.Time, description string) *TimeLog {
	return &TimeLog{
		TimeLogId:     timeLogId,
		StartTime:     startTime,
		EndTime:       endTime,
		LoggedAt:      time.Now(),
		TotalDuration: endTime.Sub(startTime).String(),
		Description:   description,
	}
}
