package responseMapper

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/response"
)

func ToTimeLogResponse(timeLog *common.TimeLogResponse) *response.TimeLogResponse {
	return &response.TimeLogResponse{
		TimeLogId:     timeLog.TimeLogId.String(),
		StartTime:     timeLog.StartTime,
		EndTime:       timeLog.EndTime,
		LoggedAt:      timeLog.LoggedAt,
		TotalDuration: timeLog.TotalDuration,
		Description:   timeLog.Description,
	}
}

func ToTimeLogListResponse(timeLogList []*common.TimeLogResponse) []*response.TimeLogResponse {
	var timeLogResponseList []*response.TimeLogResponse
	for _, t := range timeLogList {
		timeLogResponseList = append(timeLogResponseList, ToTimeLogResponse(t))
	}

	return timeLogResponseList
}
