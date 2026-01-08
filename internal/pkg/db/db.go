package db

import (
	"database/sql"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Options struct {
	DSN string
	Env string
}

func Open(opts Options) (*gorm.DB, error) {
	gormlogger := logger.Default.LogMode((logger.Silent))
	if opts.Env != "prod" {
		gormlogger = logger.Default.LogMode(logger.Warn)
	}
	gdb, err := gorm.Open(mysql.Open(opts.DSN), &gorm.Config{
		Logger: gormlogger,
	})

	if err != nil {
     return nil,err
	}

	sqlDB,err :=gdb.DB()
	if err !=nil {
		  return nil,err
	}

	if err :=sqlDB.Ping();err != nil {
		  return nil,err
	}
	return gdb,nil
}

func tunePool(sqlDB *sql.DB){
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
}
