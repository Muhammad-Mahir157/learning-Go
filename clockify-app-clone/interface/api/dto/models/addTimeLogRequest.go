package models

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
)

type AddTimeLogRequest struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

func (req AddTimeLogRequest) ToAddLogTimeRequest() *usecase.AddTimeLogRequestModel {
	return &usecase.AddTimeLogRequestModel{
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Description: req.Description,
	}
}
