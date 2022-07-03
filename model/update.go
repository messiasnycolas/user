package model

import (
	"github.com/messiasnycolas/user/db"
)

func Update(id int64, user User) (rowsAffected int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	res, err := conn.Exec("update users set name=? where id=?", user.Name, id)
	if err != nil {
		return
	}

	rowsAffected, err = res.RowsAffected()
	return
}
