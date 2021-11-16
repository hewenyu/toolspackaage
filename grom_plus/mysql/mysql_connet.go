package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

/**
 * NewConnection
 * 创建链接
 *  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
 */
func (p *MYSQL) NewConnection() (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=%v&loc=%v",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Name,
		p.ParseTime,
		p.Loc,
	)
	
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dbUrl, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: strings.ToLower(p.TablePrefix), // 表名前缀，`User` 的表名应该是 `t_users`
			// SingularTable: true,                           // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	
	return conn, err
}
