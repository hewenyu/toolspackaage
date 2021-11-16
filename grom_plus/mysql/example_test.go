package mysql

import (
	uuid "github.com/satori/go.uuid"
	"github.com/hewenyu/toolspackage/grom_plus/sample_models"
	"github.com/hewenyu/toolspackage/systemctl_plus/random"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"testing"
	"time"
)

var (
	db               *gorm.DB
	modelWithHistory = []interface{}{&sample_models.UserInfo{}}
)

func init() {
	db, _ = NewConn()
}

func NewConn() (*gorm.DB, error) {
	mysqlConfig := MYSQL{
		Name:        "test",
		Password:    "test",
		User:        "test",
		Port:        "3306",
		Host:        "localhost",
		TablePrefix: "test",  // defult dev_
		ParseTime:   "True",  // defult True
		Loc:         "Local", // defult Local
	}
	return mysqlConfig.NewConnection()
}

// GetDB 获取db
func GetDB() *gorm.DB {
	
	sqlDriver, err := db.DB()
	
	if err != nil {
		log.Println(err.Error())
	}
	if err := sqlDriver.Ping(); err != nil {
		err := sqlDriver.Close()
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		db, _ = NewConn()
	}
	return db
}

// SetupDatabase 自动初始化
func SetupDatabase(db *gorm.DB) error {
	err := db.AutoMigrate(modelWithHistory...)
	
	return err
}

func Test_demo(t *testing.T) {
	_db := GetDB()
	
	for {
		_uid := uuid.NewV4()
		
		_name := random.NonceStr()
		// UserInfo :=
		newUser := sample_models.UserInfo{
			ID:        _uid,
			Name:      _name,
			CreatedAt: time.Now().Local(),
		}
		
		// 更新用户字段
		_db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},              // key colume
			DoUpdates: clause.AssignmentColumns([]string{"name"}), // column needed to be updated
		}).Create(&newUser)
		
		log.Println("db commit")
		
		time.Sleep(time.Second)
		
	}
	
}
