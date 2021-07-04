package cryptodb

import (
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	GormDB *gorm.DB
}

const (
	MYSQL_ERR_DUP_ENTRY = 1062
)

func NewDB(dsn string) (db *DB, err error) {
	gormDB, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return
	}
	db = &DB{
		GormDB: gormDB,
	}
	return
}

func IsErrDupEntry(err error) bool {
	switch err.(type) {
	case *mysql.MySQLError:
		if err.(*mysql.MySQLError).Number == MYSQL_ERR_DUP_ENTRY {
			return true
		}
	}
	return false
}
