package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"project-NDL/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	ConfigurationString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", ConfigurationString)
	if err != nil {
		panic("connection error")
	}

	err := db.Ping()

	if err != nil {
		fmt.Println(err)
		panic("DSN error")
	}
}

func CreateCon() *sql.DB {
	return db
}
