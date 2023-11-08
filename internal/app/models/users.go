package model

type Users struct {
	ID       int    `gorm:"primaryKey; autoIncrement"`
	Username string `gorm:"type:varchar(191); unique; not null"`
	Email    string `gorm:"type:varchar(191); unique; not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:int;default:2;not null"`
	Status   string `gorm:"type:int;default:0;not null"`
	Created  int64  `gorm:"autoCreateTime"`
	Updated  int64  `gorm:"autoUpdateTime"`
}
