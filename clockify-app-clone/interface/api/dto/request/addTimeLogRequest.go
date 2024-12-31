package request

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
)

type AddTimeLogRequest struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Description string    `json:"description"`
}

func (req AddTimeLogRequest) ToAddLogTimeRequest() *common.AddTimeLogRequestModel {
	return &common.AddTimeLogRequestModel{
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Description: req.Description,
	}
}
