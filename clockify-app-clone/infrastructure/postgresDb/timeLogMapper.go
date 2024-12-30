package postgresDb

import "github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"

func fromEntitytoDbModel(timeLogEntity *entities.TimeLog) *TimeLog {
	t := &TimeLog{
		TimeLogId:     timeLogEntity.TimeLogId,
		StartTime:     timeLogEntity.StartTime,
		EndTime:       timeLogEntity.EndTime,
		LoggedAt:      timeLogEntity.LoggedAt,
		TotalDuration: timeLogEntity.TotalDuration,
		Description:   timeLogEntity.Description,
	}

	return t
}

func fromDbToEntityModel(dbTimeLog *TimeLog) *entities.TimeLog {
	t := &entities.TimeLog{
		TimeLogId:     dbTimeLog.TimeLogId,
		StartTime:     dbTimeLog.StartTime,
		EndTime:       dbTimeLog.EndTime,
		LoggedAt:      dbTimeLog.LoggedAt,
		TotalDuration: dbTimeLog.TotalDuration,
		Description:   dbTimeLog.Description,
	}

	return t
}
