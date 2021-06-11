package zsql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSqlite3(fileName) (db *zsql.DB, err error) {
	db = &zsql.DB{}
	db.Db, err = gorm.Open("sqlte3", fileName)
	return
}
