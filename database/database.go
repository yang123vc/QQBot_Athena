package database

import (
	"database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func ConnectDB(dbName string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:syx569927585@/"+dbName+"?charset=utf8")
	if err != nil {
		return
	}
	//defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	if err = db.Ping(); err != nil {
		return
	}
	return
}
