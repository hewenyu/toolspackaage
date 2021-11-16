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
	sqlDirver, err := db.DB()
	
	if err != nil {
		log.Println(err.Error())
	}
	
	sqlDirver.SetMaxIdleConns(10)                   //最大空闲连接数
	sqlDirver.SetMaxOpenConns(30)                   //最大连接数
	sqlDirver.SetConnMaxLifetime(time.Second * 300) //设置连接空闲超时
	
	// defer sqlDirver.Close()
	
	SetupDatabase(db)
}

func NewConn() *gorm.DB {
	
	_pg_config := POSTGRES{
		Name:        "postgres",
		User:        "postgres",
		Host:        "localhost",
		Password:    "example",
		Port:        "5432",
		TablePrefix: "test_",
		SSLMODE:     "disable",
	}
	
	return _pg_config.NewConnection()
	
}

/**
 * GetDB
 */
func GetDB() *gorm.DB {
	
	sqlDirver, err := db.DB()
	
	if err != nil {
		log.Println(err.Error())
	}
	
	if err := sqlDirver.Ping(); err != nil {
		err := sqlDirver.Close()
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
