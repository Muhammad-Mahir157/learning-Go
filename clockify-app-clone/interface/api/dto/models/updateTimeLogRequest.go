package models

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
	"github.com/google/uuid"
)

type UpdateTimeLogRequest struct {
	TimeLogId   uuid.UUID `json:"timeLogId"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

func (req UpdateTimeLogRequest) ToUpdateLogTimeRequest() *usecase.UpdateTimeLogRequestModel {
	return &usecase.UpdateTimeLogRequestModel{
		TimeLogId:   req.TimeLogId,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Description: req.Description,
	}
}
