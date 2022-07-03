package model

import "github.com/messiasnycolas/user/db"

func Insert(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	res, err := conn.Exec("insert into users(name) values(?)", user.Name)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()
	return
}
