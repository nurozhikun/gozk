package zsql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite3(fileName string) (db *DB, err error) {
	db = &DB{}
	// db.Db, err = gorm.Open("sqlte3", fileName)
	db.DB, err = gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	return
}
