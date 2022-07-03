package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/messiasnycolas/user/config"
)

var dbVars = config.LoadEnv("DB_USERNAME", "DB_PASSWORD", "DB_NAME")
var dbUsername, dbPassword, dbName = dbVars[0], dbVars[1], dbVars[2]
var dbUrl string = fmt.Sprintf("%v:%v@/%v", dbUsername, dbPassword, dbName)

func OpenDBConnection() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dbUrl)
	return
}
