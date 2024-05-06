package models

type User struct {
	BaseUUID
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255);not null"`
}
