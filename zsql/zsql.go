package zsql

import (
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/nurozhikun/gozk/zlogger"
	"github.com/nurozhikun/gozk/zutils"
	"gorm.io/gorm"
)

const (
	TypeSqlite   = 0
	TypeMysqlite = 1
)

type Cfg struct {
	Type     int    `ini:"type"` // 0:sqlite, 1:mysql
	Addr     string `ini:"addr"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"` // or filename
	MaxOpen  int    `ini:"max_open"`
	MaxIdle  int    `ini:"max_idle"`
}

// type DB struct {
// 	*gorm.DB
// }

type DB struct {
	*gorm.DB
}

func OpenDB(cfg *Cfg) (db *DB) {
	var err error
	switch cfg.Type {
	case TypeMysqlite:
		db, err = OpenMysql(cfg.User, cfg.Password, cfg.Addr, cfg.Database)
	default:
		filename, _ := filepath.Abs(cfg.Database)
		zlogger.Info("sqlite filename:", filename)
		db, err = OpenSqlite3(filename)
	}
	if nil != db {
		db.StdDB().SetMaxOpenConns(cfg.MaxOpen)
		db.StdDB().SetMaxIdleConns(cfg.MaxIdle)
	}
	if nil != err {
		zlogger.Error(err)
	}
	return
}

func (db *DB) StdDB() *sql.DB {
	d, err := db.DB.DB()
	if nil != err {
		zlogger.Error(err)
		return nil
	}
	return d
}

func (db *DB) Close() {
	d := db.StdDB()
	if nil != d {
		d.Close()
	}
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

func (db *DB) SelectRows(dest interface{}, query string, args ...interface{}) error {
	xdb := sqlx.NewDb(db.StdDB(), db.Dialector.Name())
	if nil == xdb {
		return errors.New("failed to create the sqlx.DB")
	}
	return xdb.Select(dest, query, args...)
}
