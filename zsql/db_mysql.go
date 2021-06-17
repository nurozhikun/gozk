package zsql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMysqlDns(dsn string) (db *DB, err error) {
	db = &DB{}
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func OpenMysql(user, pass, addr, dbname string) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, addr, dbname)
	return OpenMysqlDns(dsn)
}
