package database

import (
	"database/sql"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

const (
	// 数据库信息
	dbUser       string = "root"
	dbPasswd     string = "syx569927585"
	dbLoc        string = "tcp(47.100.182.193:3306)"
	MaxIdleConns int    = 500
	MaxOpenConns int    = 500
)

func ConnectDB(dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dbUser+":"+dbPasswd+"@"+dbLoc+"/"+dbName+"?charset=utf8")
	if err != nil {
		return
	}
	//defer db.Close()

	db.SetConnMaxLifetime(time.Second * 4)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	if err = db.Ping(); err != nil {
		return
	}
	return
}
