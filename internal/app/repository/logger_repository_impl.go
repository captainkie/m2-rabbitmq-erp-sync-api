package repository

import (
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"

	"gorm.io/gorm"
)

type LoggerRepositoryImpl struct {
	Db *gorm.DB
}

func NewLoggerRepositoryImpl(Db *gorm.DB) LoggerRepository {
	return &LoggerRepositoryImpl{Db: Db}
}

// Create implements LoggerRepository
func (l *LoggerRepositoryImpl) CreateConnectionLog(log model.ConnectionLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateAddLog(log model.AddLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateUpdateLog(log model.UpdateLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateStockLog(log model.StockLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateStoreLog(log model.StoreLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreatePostflagLog(log model.PostflagLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateImageLog(log model.ImageLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}

func (l *LoggerRepositoryImpl) CreateDailysaleLog(log model.DailysaleLogs) {
	result := l.Db.Create(&log)
	helpers.ErrorPanic(result.Error)
}
