package db

import (
	"juggle/basic/src/lib"
	"juggle/basic/src/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

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

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	migrate()
}

func migrate()  {
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Printf("database migrate error: %s\n", err.Error())
	}
}

// DB GORM
func DB() *gorm.DB {
	return db
}
