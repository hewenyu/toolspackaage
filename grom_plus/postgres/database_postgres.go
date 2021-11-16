package postgres

import (
	"fmt"
	"github.com/hewenyu/toolspackage/global_variable"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	ModelWithHistory []interface{}
)

/*
open 链接PG数据库
dsn "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
*/
func OpenPG(dsn string) (db *gorm.DB) {
	
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // 禁用缓存
	}), &gorm.Config{})
	
	if err != nil {
		global_variable.CLOG.Fatal(fmt.Sprintf("数据库链接失败:%v", err.Error()))
	}
	
	return
}

func ResetDatabase(db *gorm.DB) {
	//cleanDatabase(db)
	CleanDatabase(db)
	SetupDatabase(db)
}

/*
cleanDatabase 删除表
*/
func CleanDatabase(db *gorm.DB) {
	_ = db.Migrator().DropTable(ModelWithHistory...)
}

/*
setupDatabase 为了使用一些常用组建
*/
func SetupDatabase(db *gorm.DB) {
	db.Exec("create extension IF NOT EXISTS hstore;")
	// 为了使用uuid
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\"")
	_ = db.AutoMigrate(ModelWithHistory...)
}
