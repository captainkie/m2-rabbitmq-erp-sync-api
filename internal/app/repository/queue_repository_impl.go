package repository

import (
	"errors"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"

	"gorm.io/gorm"
)

type QueueRepositoryImpl struct {
	Db *gorm.DB
}

func NewQueueRepositoryImpl(Db *gorm.DB) QueueRepository {
	return &QueueRepositoryImpl{Db: Db}
}

// Connection implements QueueRepository
func (q *QueueRepositoryImpl) Connection(connection model.Connections) (model.Connections, error) {
	result := q.Db.Create(&connection)
	if result != nil {
		return connection, nil
	} else {
		return connection, errors.New("could not create connection")
	}
}

// Create queue type connection implements QueueRepository
func (q *QueueRepositoryImpl) CreateConnection(connections []model.ConnectionQueues) ([]model.ConnectionQueues, error) {
	result := q.Db.Create(connections)
	if result != nil {
		return connections, nil
	} else {
		return connections, errors.New("could not create connections queue")
	}
}

// Update queue type connection implements QueueRepository
func (q *QueueRepositoryImpl) UpdateConnection(connection model.ConnectionQueues) {
	q.Db.Model(&connection).Update("status", connection.Status)
}

// Create queue type add products implements QueueRepository
func (q *QueueRepositoryImpl) CreateTypeAdd(products []model.AddQueues) ([]model.AddQueues, error) {
	result := q.Db.Create(products)
	if result != nil {
		return products, nil
	} else {
		return products, errors.New("could not create add queue")
	}
}

// Update queue type add products implements QueueRepository
func (q *QueueRepositoryImpl) UpdateTypeAdd(product model.AddQueues) {
	q.Db.Model(&product).Update("status", product.Status)
}

// Create queue type update products implements QueueRepository
func (q *QueueRepositoryImpl) CreateTypeUpdate(products []model.UpdateQueues) ([]model.UpdateQueues, error) {
	result := q.Db.Create(products)
	if result != nil {
		return products, nil
	} else {
		return products, errors.New("could not create update queue")
	}
}

// Update queue type update products implements QueueRepository
func (q *QueueRepositoryImpl) UpdateTypeUpdate(product model.UpdateQueues) {
	q.Db.Model(&product).Update("status", product.Status)
}

// CreateStock implements QueueRepository
func (q *QueueRepositoryImpl) CreateStock(stocks []model.StockQueues) ([]model.StockQueues, error) {
	result := q.Db.Create(stocks)
	if result != nil {
		return stocks, nil
	} else {
		return stocks, errors.New("could not create stocks queue")
	}
}

// UpdateStock implements QueueRepository
func (q *QueueRepositoryImpl) UpdateStock(stock model.StockQueues) {
	q.Db.Model(&stock).Update("status", stock.Status)
}

// CreateStore implements QueueRepository
func (q *QueueRepositoryImpl) CreateStore(stores []model.StoreQueues) ([]model.StoreQueues, error) {
	result := q.Db.Create(stores)
	if result != nil {
		return stores, nil
	} else {
		return stores, errors.New("could not create stores queue")
	}
}

// UpdateStore implements QueueRepository
func (q *QueueRepositoryImpl) UpdateStore(store model.StoreQueues) {
	q.Db.Model(&store).Update("status", store.Status)
}

// Create queue type postflag implements QueueRepository
func (q *QueueRepositoryImpl) CreatePostflag(postflags []model.PostflagQueues) ([]model.PostflagQueues, error) {
	result := q.Db.Create(postflags)
	if result != nil {
		return postflags, nil
	} else {
		return postflags, errors.New("could not create postflags queue")
	}
}

// Update queue type postflag implements QueueRepository
func (q *QueueRepositoryImpl) UpdatePostflag(postflag model.PostflagQueues) {
	q.Db.Model(&postflag).Update("status", postflag.Status)
}

// Create queue type image implements QueueRepository
func (q *QueueRepositoryImpl) CreateImage(images []model.ImageQueues) ([]model.ImageQueues, error) {
	var createdImages []model.ImageQueues
	for _, image := range images {
		// Check if a record with the same "image" value exists
		var existingImage model.ImageQueues
		result := q.Db.Where("image = ?", image.Image).First(&existingImage)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// Record not found, create a new one
				result := q.Db.Create(&image)
				if result.Error == nil {
					createdImages = append(createdImages, image)
				}
			} else {
				return createdImages, result.Error
			}
		}
	}

	return createdImages, nil
}

// Update queue type image implements QueueRepository
func (q *QueueRepositoryImpl) UpdateImage(image model.ImageQueues) {
	q.Db.Model(&image).Update("status", image.Status)
}

// Delete queue type image implements QueueRepository
func (q *QueueRepositoryImpl) DeleteImage(id int) {
	var images model.ImageQueues
	result := q.Db.Where("id = ?", id).Delete(&images)
	helpers.ErrorPanic(result.Error)
}

// Create queue type dailysales implements QueueRepository
func (q *QueueRepositoryImpl) CreateDailySales(dailysales []model.DailysaleQueues) ([]model.DailysaleQueues, error) {
	result := q.Db.Create(dailysales)
	if result != nil {
		return dailysales, nil
	} else {
		return dailysales, errors.New("could not create dailysales queue")
	}
}

// UpdateStock implements QueueRepository
func (q *QueueRepositoryImpl) UpdateDailySales(dailysale model.DailysaleQueues) {
	q.Db.Model(&dailysale).Update("status", dailysale.Status)
}
