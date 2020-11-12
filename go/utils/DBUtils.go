package utils

import (
	"database/sql"
	"fmt"
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

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/mybook")
	if err != nil {
		fmt.Println(err)
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
