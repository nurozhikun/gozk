package zsql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenMysqlDns(dsn string) (db *DB, err error) {
	db = &DB{}
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			},
		),
	})
	return
}

// time zone = UTC time
func OpenMysql(user, pass, addr, dbname string) (db *DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", user, pass, addr, dbname)
	return OpenMysqlDns(dsn)
}
