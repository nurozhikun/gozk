package zsql

import (
	"database/sql"

	"gitee.com/sienectagv/gozk/zlogger"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func (db *DB) StdDB() *sql.DB {
	d, err := db.DB.DB()
	if nil != err {
		zlogger.Error(err)
		return nil
	}
	return d
}
