package zsql

import (
	"database/sql"
	"fmt"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zutils"
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

func (db *DB) AddForeignKey(ftable, fkey, rtable, rkey string) (*DB, error) {
	key := "fk_" + fkey + "_" + rkey
	// var cont int64
	// rows, err := db.Exec("SELECT * FROM information_schema.KEY_COLUMN_USAGE where constraint_name = ?", key).Rows()
	// zlogger.Error(err)
	// for rows.Next() {
	// 	cont = cont + 1
	// }
	// zlogger.Info(cont)
	rows, err := db.StdDB().Query("SELECT * FROM information_schema.KEY_COLUMN_USAGE where constraint_name = ?", key)
	if err == nil {
		n := 0
		for rows.Next() {
			n = n + 1
		}
		// zlogger.Info(n)
		if n > 0 {
			return db, zutils.ErrHasExist
		}
	}
	// n, _ := r.RowsAffected()
	// zlogger.Info(n, err)
	sqlStr := fmt.Sprintf("alter table %s add constraint %s foreign key(%s) REFERENCES %s(%s) ON UPDATE CASCADE ON DELETE SET NULL",
		ftable,
		key,
		fkey,
		rtable,
		rkey)
	// zlogger.Info(sqlStr)
	result := db.Exec(sqlStr)
	return db, result.Error
}
