package common

import (
	"time"

	"github.com/google/uuid"
)

type AddTimeLogRequestModel struct {
	StartTime   time.Time
	EndTime     time.Time
	Description string
}

type UpdateTimeLogRequestModel struct {
	TimeLogId   uuid.UUID
	StartTime   time.Time
	EndTime     time.Time
	Description string
}
