package repository

import model "github.com/captainkie/websync-api/internal/app/models"

type QueueRepository interface {
	Connection(connection model.Connections) (model.Connections, error)
	CreateConnection(connection []model.ConnectionQueues) ([]model.ConnectionQueues, error)
	UpdateConnection(connection model.ConnectionQueues)
	CreateTypeAdd(products []model.AddQueues) ([]model.AddQueues, error)
	UpdateTypeAdd(product model.AddQueues)
	CreateTypeUpdate(products []model.UpdateQueues) ([]model.UpdateQueues, error)
	UpdateTypeUpdate(products model.UpdateQueues)
	CreateStock(stocks []model.StockQueues) ([]model.StockQueues, error)
	UpdateStock(stock model.StockQueues)
	CreateStore(stores []model.StoreQueues) ([]model.StoreQueues, error)
	UpdateStore(store model.StoreQueues)
	CreatePostflag(stores []model.PostflagQueues) ([]model.PostflagQueues, error)
	UpdatePostflag(store model.PostflagQueues)
	CreateImage(images []model.ImageQueues) ([]model.ImageQueues, error)
	UpdateImage(image model.ImageQueues)
	DeleteImage(id int)
	CreateDailySales(dailysales []model.DailysaleQueues) ([]model.DailysaleQueues, error)
	UpdateDailySales(dailysale model.DailysaleQueues)
}
