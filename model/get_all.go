package model

import "github.com/messiasnycolas/user/db"

func GetAll() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("select * from users")
	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			continue // log somewhere in the future
		}
		users = append(users, user)
	}

	return
}
