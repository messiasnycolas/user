package model

import "github.com/messiasnycolas/user/db"

func Delete(id int64) (rowsAffected int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	res, err := conn.Exec("delete from users where id=?", id)
	if err != nil {
		return
	}

	id, err = res.RowsAffected()

	return
}
