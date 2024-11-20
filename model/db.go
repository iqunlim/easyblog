package model

import (
	"fmt"
	"sync"

	"github.com/iqunlim/easyblog/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var (
	db *gorm.DB
	o sync.Once
)
func GetDB() *gorm.DB {
	o.Do(func() {
		db = func() *gorm.DB {
			db, err := gorm.Open(sqlite.Open(config.DatabaseFileLocation), config.GORMConfig)
			if err != nil {
				panic("Failed to connect to database: " + err.Error())
			}
			db.AutoMigrate(&User{})
			db.AutoMigrate(&BlogPost{})
			fmt.Println("Connected to database")
			return db
		}()
	})
	return db
}