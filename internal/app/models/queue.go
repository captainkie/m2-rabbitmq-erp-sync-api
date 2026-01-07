package model

type ConnectionQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"type:varchar(191);not null"`
	Type          string `gorm:"type:varchar(191);default:connection;"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type AddQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID int    `gorm:"not null"`
	Type          string `gorm:"type:varchar(191);default:add;"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type UpdateQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID int    `gorm:"not null"`
	Type          string `gorm:"type:varchar(191);default:update;"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type StockQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID int    `gorm:"not null"`
	Type          string `gorm:"type:varchar(191);default:stock;"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type StoreQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID int    `gorm:"not null"`
	Type          string `gorm:"type:varchar(191);default:store;"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type DailysaleQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"not null"`
	OrderID       string `gorm:"type:varchar(191);not null"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type PostflagQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID int    `gorm:"not null"`
	JsonData      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}

type ImageQueues struct {
	ID            int    `gorm:"primaryKey; autoIncrement"`
	TransactionID string `gorm:"not null"`
	Image         string `gorm:"not null"`
	DirectoryPath string `gorm:"not null"`
	SyncDate      string `gorm:"not null"`
	Status        string `gorm:"type:varchar(191);default:pending;"`
	Created       int64  `gorm:"autoCreateTime"`
	Updated       int64  `gorm:"autoUpdateTime"`
}
