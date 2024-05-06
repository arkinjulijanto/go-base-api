package base

import (
	"github.com/arkinjulijanto/go-base-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(c config.Config) (err error) {
	dsn := config.InitDBDsn(c)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetDBConn() *gorm.DB {
	return db
}
