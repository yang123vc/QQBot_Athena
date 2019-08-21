package database

import (
	"database/sql"
	"errors"
)
import _ "github.com/go-sql-driver/mysql"

const (
	// 数据库信息
	dbUser       string = "root"
	dbPasswd     string = "syx569927585"
	dbLoc        string = "tcp(localhost:3306)"
	MaxIdleConns int    = 10
	MaxOpenConns int    = 10
)

func ConnectDB(dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dbUser+":"+dbPasswd+"@"+dbLoc+"/"+dbName+"?charset=utf8")
	if err != nil {
		return
	}
	//defer db.Close()

	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)

	if err = db.Ping(); err != nil {
		return
	}
	return
}

// 修改中
func DBInsert(db *sql.DB, table string) (err error) {
	if db == nil {
		return errors.New("db Not Found")
	}

	_, err = db.Prepare("INSERT " + "table" + " SET ")

	return nil
}
