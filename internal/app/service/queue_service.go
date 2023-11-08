package service

import "github.com/captainkie/websync-api/types/response"

type QueueService interface {
	Connection(connection []byte) response.CreateConnectionResponse
	CreateConnectionQueue(id string)
	UpdateConnectionQueue(id int, status string)
	CreateProductsQueue(qtype string, products []byte)
	UpdateProductsQueue(id int, qtype, status string, products []byte)
	CreateStockQueue(stocks []byte)
	UpdateStockQueue(id int, status string, stocks []byte)
	CreateStoreQueue(stores []byte)
	UpdateStoreQueue(id int, status string, stores []byte)
	CreatePostflagQueue(id, status, message string)
	UpdatePostflagQueue(id int, status string)
	CreateImageQueue(id, status, message string)
	UpdateImageQueue(id int, status string)
}
