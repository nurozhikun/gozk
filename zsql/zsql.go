package zsql

import (
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}
