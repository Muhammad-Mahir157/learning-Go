package request

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/google/uuid"
)

type UpdateTimeLogRequest struct {
	TimeLogId   uuid.UUID `json:"timeLogId"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

func (req UpdateTimeLogRequest) ToUpdateLogTimeRequest() *common.UpdateTimeLogRequestModel {
	return &common.UpdateTimeLogRequestModel{
		TimeLogId:   req.TimeLogId,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Description: req.Description,
	}
}
