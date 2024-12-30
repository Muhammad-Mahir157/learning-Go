package query

import "github.com/Muhammad-Mahir157/clockify-app-clone/application/common"

type TimeLogQueryResponse struct {
	Response *common.TimeLogResponse
}

type TimeLogQueryResponseList struct {
	List []*common.TimeLogResponse
}
