package utils

import (
	"daily/go/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var instance *DBUtils

type DBUtils struct {
	db *sql.DB
}

/**
单例
*/
func GetInstance() *DBUtils {
	if instance != nil {
		return instance
	}

	db, err := sql.Open("mysql", "root:123#hxq@tcp(127.0.0.1:3306)/my_daily")
	if err != nil {
		logger.Error.Println(err)
		return nil
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(1)

	var dbUtils = &DBUtils{}
	dbUtils.db = db

	instance = dbUtils

	return instance
}

func (dbUtils *DBUtils) GetDB() *sql.DB {
	return dbUtils.db
}
