package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/messiasnycolas/user/config"
)

func OpenConnection() (db *sql.DB, err error) {
	env := config.GetDB()
	dbUrl := fmt.Sprintf("%v:%v@/%v", env.User, env.Pass, env.Name)

	db, err = sql.Open("mysql", dbUrl)
	return
}
