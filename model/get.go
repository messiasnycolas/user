package model

import "github.com/messiasnycolas/user/db"

func Get(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)
	return
}
