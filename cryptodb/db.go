package cryptodb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	GormDB *gorm.DB
}

func NewDB(dsn string) (db *DB, err error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	db = &DB{
		GormDB: gormDB,
	}
	return
}
