package common

import "time"

type TimeLogRequestModel struct {
	StartTime   time.Time
	EndTime     time.Time
	Description string
}
