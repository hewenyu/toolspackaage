package postgres

import (
	"gorm.io/gorm"
	"log"
)

func ResetDatabase(db *gorm.DB) {
	CleanDatabase(db)
	SetupDatabase(db)
}

/*
cleanDatabase 删除表
*/
func CleanDatabase(db *gorm.DB) {
	err := db.Migrator().DropTable(modelWithHistory...)
	if err != nil {
		log.Fatal(err.Error())
	}
}
