package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	UserId string
	Name   string
}

func QueryUserById(id string, db *sql.DB) (User, error) {
	var user User
	row := db.QueryRow("select id, name from user where id = ?", id)
	err := row.Scan(&user.UserId, &user.Name)

	if err != nil {
		return user, errors.Wrap(err, "no such user")
	}
	return user, nil
}

func main() {
	db, err := sql.Open("mysql", "root:xiaojiao@tcp(127.0.0.1:3306)/user")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user, err := QueryUserById("100", db)

	if err != nil {
		fmt.Printf("query user err: %v", err)
		return
	}

	fmt.Println("query user: ", user)
}
