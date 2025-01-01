package adaptor

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/models"
	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
)

func ToTimeLogResponse(timeLog *usecase.TimeLogResponse) *models.TimeLogResponse {
	return &models.TimeLogResponse{
		TimeLogId:     timeLog.TimeLogId.String(),
		StartTime:     timeLog.StartTime,
		EndTime:       timeLog.EndTime,
		LoggedAt:      timeLog.LoggedAt,
		TotalDuration: timeLog.TotalDuration,
		Description:   timeLog.Description,
	}
}

func ToTimeLogListResponse(timeLogList []*usecase.TimeLogResponse) []*models.TimeLogResponse {
	var timeLogResponseList []*models.TimeLogResponse
	for _, t := range timeLogList {
		timeLogResponseList = append(timeLogResponseList, ToTimeLogResponse(t))
	}

	return timeLogResponseList
}
