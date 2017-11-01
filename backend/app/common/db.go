package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", getConnectionString())

	if err != nil {
		log.Fatal("Db connection error: ", err)
	}
}

func getConnectionString() string {
	var connstr string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		AppConfig.DbUser,
		AppConfig.DbPassword,
		AppConfig.DbHost,
		AppConfig.DbPort,
		AppConfig.DbName,
	)

	return connstr
}
