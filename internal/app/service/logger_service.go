package service

import "github.com/captainkie/websync-api/types/request"

type LoggerService interface {
	CreateAddLog(log request.CreateLogRequest)
	CreateUpdateLog(log request.CreateLogRequest)
	CreateStockLog(log request.CreateLogRequest)
	CreateStoreLog(log request.CreateLogRequest)
	CreatePostflagLog(log request.CreateLogRequest)
	CreateImageLog(log request.CreateLogRequest)
	CreateDailysaleLog(log request.CreateLogDailyRequest)
}
