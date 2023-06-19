package db

import (
	"github.com/Shota-Kurahashi/echo-next/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(mysql.Open(config.GetDbUrl()), &gorm.Config{})

	config.CheckError(err)

	return db
}

func Close(db *gorm.DB) {

	dbSql, err := db.DB()

	config.CheckError(err)

	dbSql.Close()
}
