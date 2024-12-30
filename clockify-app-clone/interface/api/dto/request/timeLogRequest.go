package request

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
)

type TimeLogRequest struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

func (req TimeLogRequest) ToServiceLogTimeRequest() *common.TimeLogRequestModel {
	return &common.TimeLogRequestModel{
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Description: req.Description,
	}
}
