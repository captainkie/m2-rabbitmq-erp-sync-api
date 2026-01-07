package service

import (
	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/internal/app/repository"
	"github.com/captainkie/websync-api/types/request"
	"github.com/go-playground/validator/v10"
)

type LoggerServiceImpl struct {
	LoggerRepository repository.LoggerRepository
	Validate         *validator.Validate
}

func NewLoggerServiceImpl(loggerRepository repository.LoggerRepository, validate *validator.Validate) LoggerService {
	return &LoggerServiceImpl{
		LoggerRepository: loggerRepository,
		Validate:         validate,
	}
}

// Create implements LoggerService
func (l *LoggerServiceImpl) CreateConnectionLog(log request.CreateLogRequest) {
	newLog := model.ConnectionLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateConnectionLog(newLog)
}

func (l *LoggerServiceImpl) CreateAddLog(log request.CreateLogRequest) {
	newLog := model.AddLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateAddLog(newLog)
}

func (l *LoggerServiceImpl) CreateUpdateLog(log request.CreateLogRequest) {
	newLog := model.UpdateLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateUpdateLog(newLog)
}

func (l *LoggerServiceImpl) CreateStockLog(log request.CreateLogRequest) {
	newLog := model.StockLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateStockLog(newLog)
}

func (l *LoggerServiceImpl) CreateStoreLog(log request.CreateLogRequest) {
	newLog := model.StoreLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateStoreLog(newLog)
}

func (l *LoggerServiceImpl) CreatePostflagLog(log request.CreateLogRequest) {
	newLog := model.PostflagLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreatePostflagLog(newLog)
}

func (l *LoggerServiceImpl) CreateImageLog(log request.CreateLogImageRequest) {
	newLog := model.ImageLogs{
		TransactionID: log.TransactionID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncImage:     log.SyncImage,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateImageLog(newLog)
}

func (l *LoggerServiceImpl) CreateDailysaleLog(log request.CreateLogDailyRequest) {
	newLog := model.DailysaleLogs{
		TransactionID: log.TransactionID,
		OrderID:       log.OrderID,
		Status:        log.Status,
		StatusCode:    log.StatusCode,
		Message:       log.Message,
		SyncJson:      log.SyncJson,
		SyncData:      log.SyncData,
		SyncDate:      log.SyncDate,
	}

	l.LoggerRepository.CreateDailysaleLog(newLog)
}
