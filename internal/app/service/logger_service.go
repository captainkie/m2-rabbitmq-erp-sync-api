package service

import "github.com/captainkie/websync-api/types/request"

type LoggerService interface {
	CreateConnectionLog(log request.CreateLogRequest)
	CreateAddLog(log request.CreateLogRequest)
	CreateUpdateLog(log request.CreateLogRequest)
	CreateStockLog(log request.CreateLogRequest)
	CreateStoreLog(log request.CreateLogRequest)
	CreatePostflagLog(log request.CreateLogRequest)
	CreateImageLog(log request.CreateLogImageRequest)
	CreateDailysaleLog(log request.CreateLogDailyRequest)
}
