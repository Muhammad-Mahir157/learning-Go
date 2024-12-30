package interfaces

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/common"
	"github.com/Muhammad-Mahir157/clockify-app-clone/application/query"
)

type TimeLogService interface {
	//AddTimeLog
	AddTimeLog(timeLogReq *common.TimeLogRequestModel) (*common.TimeLogResponse, error)
	//UpdateTimeLog
	//GetAllTimeLogs
	GetAllTimeLogs() (*query.TimeLogQueryResponseList, error)
	//GetTimeLogById
	//DeleteTimeLog
}
