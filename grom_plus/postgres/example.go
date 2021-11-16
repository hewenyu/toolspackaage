package postgres

import (
	"github.com/hewenyu/toolspackage/grom_plus/sample_models"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db               *gorm.DB
	modelWithHistory = []interface{}{&sample_models.UserInfo{}}
)

func init() {
	
	db = NewConn()
	sqlDriver, err := db.DB()
	
	if err != nil {
		log.Println(err.Error())
	}
	
	sqlDriver.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDriver.SetMaxOpenConns(30)                   //最大连接数
	sqlDriver.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	
	SetupDatabase(db)
}

func NewConn() *gorm.DB {
	
	pgConfig := POSTGRES{
		Name:        "postgres",
		User:        "postgres",
		Host:        "localhost",
		Password:    "example",
		Port:        "5432",
		TablePrefix: "test_",
		SSLMODE:     "disable",
	}
	
	return pgConfig.NewConnection()
	
}

/**
 * GetDB
 */
func GetDB() *gorm.DB {
	
	sqlDriver, err := db.DB()
	
	if err != nil {
		log.Println(err.Error())
	}
	
	if err := sqlDriver.Ping(); err != nil {
		err := sqlDriver.Close()
		if err != nil {
			return nil
		}
		db = NewConn()
	}
	return db
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	err := db.AutoMigrate(modelWithHistory...)
	
	if err != nil {
		log.Fatal(err.Error())
	}
}
