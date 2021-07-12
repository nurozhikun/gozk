package zsql

import (
	"database/sql"
	"errors"
	"fmt"

	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Model = gorm.Model

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

func (db *DB) SelectRows(dest interface{}, qury string, args ...interface{}) error {
	xdb := sqlx.NewDb(db.StdDB(), db.Dialector.Name())
	if nil == xdb {
		return errors.New("failed to create the sqlx.DB")
	}
	return xdb.Select(dest, qury, args...)
}
