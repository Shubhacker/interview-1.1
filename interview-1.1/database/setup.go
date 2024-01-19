package database

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func Setup()*sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	fmt.Println("sql open " + err.Error())
	return db	
}