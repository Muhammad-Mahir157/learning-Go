package mapper

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
)

func FromTimeLogEntityToResponse(timeLogEntity *entities.TimeLog) *common.TimeLogResponse {
	if timeLogEntity == nil {
		return nil
	}

	return &common.TimeLogResponse{
		TimeLogId:     timeLogEntity.TimeLogId,
		StartTime:     timeLogEntity.StartTime,
		EndTime:       timeLogEntity.EndTime,
		LoggedAt:      timeLogEntity.LoggedAt,
		TotalDuration: timeLogEntity.TotalDuration,
		Description:   timeLogEntity.Description,
	}
}
