package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB = DBInit()

func DBInit() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("/Users/luoren/GolandProjects/src/gin-gorn-oj/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(200)
	sqlDb.SetMaxOpenConns(200)
	return db
}
