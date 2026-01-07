package repository

import model "github.com/captainkie/websync-api/internal/app/models"

type LoggerRepository interface {
	CreateConnectionLog(log model.ConnectionLogs)
	CreateAddLog(log model.AddLogs)
	CreateUpdateLog(log model.UpdateLogs)
	CreateStockLog(log model.StockLogs)
	CreateStoreLog(log model.StoreLogs)
	CreatePostflagLog(log model.PostflagLogs)
	CreateImageLog(log model.ImageLogs)
	CreateDailysaleLog(log model.DailysaleLogs)
}
