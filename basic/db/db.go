package db

import (
	"juggle/basic/lib"
	"juggle/basic/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		lib.ShutDown(err)
		return
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("database migrate error: %s", err.Error())
	}
}

// DB GORM
func DB() *gorm.DB {
	return db
}
