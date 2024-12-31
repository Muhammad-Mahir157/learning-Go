package services

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/interfaces"
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/mapper"
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/query"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/entities"
	"github.com/Muhammad-Mahir157/clockify-app-clone/domain/repositories"
	"github.com/google/uuid"
)

type TimeLogService struct {
	Repo repositories.TimeLogRepository
}

func NewTimeLogService(timeLogRepository repositories.TimeLogRepository) interfaces.TimeLogService {
	return &TimeLogService{Repo: timeLogRepository}
}

func (srvc TimeLogService) AddTimeLog(timeLogReq *common.AddTimeLogRequestModel) (*common.TimeLogResponse, error) {
	newTimeLogEntity := entities.NewTimeLogEntity(
		timeLogReq.StartTime,
		timeLogReq.EndTime,
		timeLogReq.Description,
	)

	newlyAddedTimeLog, err := srvc.Repo.Create(newTimeLogEntity)
	if err != nil {
		return nil, err
	}

	return mapper.FromTimeLogEntityToResponse(newlyAddedTimeLog), nil

}

func (srvc TimeLogService) GetAllTimeLogs() (*query.TimeLogQueryResponseList, error) {
	existingTimeLogs, err := srvc.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	var queryResponse query.TimeLogQueryResponseList
	for _, t := range existingTimeLogs {
		queryResponse.List = append(queryResponse.List, mapper.FromTimeLogEntityToResponse(t))
	}

	return &queryResponse, nil
}

func (srvc TimeLogService) UpdateTimeLog(timeLogReq *common.UpdateTimeLogRequestModel) (*common.TimeLogResponse, error) {
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

	return mapper.FromTimeLogEntityToResponse(updatedTimeLog), nil
}

func (srvc TimeLogService) DeleteTimeLog(timeLogId uuid.UUID) (bool, error) {
	isTimeLogDeleted, err := srvc.Repo.Delete(timeLogId)

	if err != nil {
		return false, err
	}

	return isTimeLogDeleted, nil
}

func (srvc TimeLogService) GetTimeLogById(timeLogId uuid.UUID) (*common.TimeLogResponse, error) {
	updatedTimeLog, err := srvc.Repo.GetById(timeLogId)

	if err != nil {
		return nil, err
	}

	return mapper.FromTimeLogEntityToResponse(updatedTimeLog), nil
}
