package interfaces

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/query"
	"github.com/google/uuid"
)

type TimeLogService interface {
	AddTimeLog(timeLogReq *common.AddTimeLogRequestModel) (*common.TimeLogResponse, error)
	UpdateTimeLog(timeLogReq *common.UpdateTimeLogRequestModel) (*common.TimeLogResponse, error)
	GetAllTimeLogs() (*query.TimeLogQueryResponseList, error)
	GetTimeLogById(timeLogId uuid.UUID) (*common.TimeLogResponse, error)
	DeleteTimeLog(timeLogId uuid.UUID) (bool, error)
}
