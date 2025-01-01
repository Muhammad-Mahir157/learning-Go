package usecase

import (
	"time"

	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/repositories"
	"github.com/google/uuid"
)

type TimeLogUsecase interface {
	AddTimeLog(timeLogReq *AddTimeLogRequestModel) (*TimeLogResponse, error)
	UpdateTimeLog(timeLogReq *UpdateTimeLogRequestModel) (*TimeLogResponse, error)
	GetAllTimeLogs() (*TimeLogQueryResponseList, error)
	GetTimeLogById(timeLogId uuid.UUID) (*TimeLogResponse, error)
	DeleteTimeLog(timeLogId uuid.UUID) (bool, error)
}

type AddTimeLogRequestModel struct {
	StartTime   time.Time
	EndTime     time.Time
	Description string
}

type UpdateTimeLogRequestModel struct {
	TimeLogId   uuid.UUID
	StartTime   time.Time
	EndTime     time.Time
	Description string
}

type TimeLogResponse struct {
	TimeLogId     uuid.UUID
	StartTime     time.Time
	EndTime       time.Time
	LoggedAt      time.Time
	TotalDuration string
	Description   string
}

// a mapper function to map entity model to response model
func FromEntityToResponse(timeLogEntity *entities.TimeLog) *TimeLogResponse {
	if timeLogEntity == nil {
		return nil
	}

	return &TimeLogResponse{
		TimeLogId:     timeLogEntity.TimeLogId,
		StartTime:     timeLogEntity.StartTime,
		EndTime:       timeLogEntity.EndTime,
		LoggedAt:      timeLogEntity.LoggedAt,
		TotalDuration: timeLogEntity.TotalDuration,
		Description:   timeLogEntity.Description,
	}
}

type TimeLogQueryResponseList struct {
	List []*TimeLogResponse
}

type timeLogUsecase struct {
	Repo repositories.TimeLog
}

func NewTimeLogService(timeLogRepository repositories.TimeLog) TimeLogUsecase {
	return &timeLogUsecase{Repo: timeLogRepository}
}

func (srvc timeLogUsecase) AddTimeLog(timeLogReq *AddTimeLogRequestModel) (*TimeLogResponse, error) {
	newTimeLogEntity := entities.NewTimeLogEntity(
		timeLogReq.StartTime,
		timeLogReq.EndTime,
		timeLogReq.Description,
	)

	newlyAddedTimeLog, err := srvc.Repo.Create(newTimeLogEntity)
	if err != nil {
		return nil, err
	}

	return FromEntityToResponse(newlyAddedTimeLog), nil

}

func (srvc timeLogUsecase) GetAllTimeLogs() (*TimeLogQueryResponseList, error) {
	existingTimeLogs, err := srvc.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	var queryResponse TimeLogQueryResponseList
	for _, t := range existingTimeLogs {
		queryResponse.List = append(queryResponse.List, FromEntityToResponse(t))
	}

	return &queryResponse, nil
}

func (srvc timeLogUsecase) UpdateTimeLog(timeLogReq *UpdateTimeLogRequestModel) (*TimeLogResponse, error) {
	updatedTimeLogEntity := entities.ExistingTimeLogEntity(
		timeLogReq.TimeLogId,
		timeLogReq.StartTime,
		timeLogReq.EndTime,
		timeLogReq.Description,
	)

	updatedTimeLog, err := srvc.Repo.Update(updatedTimeLogEntity)

	if err != nil {
		return nil, err
	}

	return FromEntityToResponse(updatedTimeLog), nil
}

func (srvc timeLogUsecase) DeleteTimeLog(timeLogId uuid.UUID) (bool, error) {
	isTimeLogDeleted, err := srvc.Repo.Delete(timeLogId)

	if err != nil {
		return false, err
	}

	return isTimeLogDeleted, nil
}

func (srvc timeLogUsecase) GetTimeLogById(timeLogId uuid.UUID) (*TimeLogResponse, error) {
	updatedTimeLog, err := srvc.Repo.GetById(timeLogId)

	if err != nil {
		return nil, err
	}

	return FromEntityToResponse(updatedTimeLog), nil
}
