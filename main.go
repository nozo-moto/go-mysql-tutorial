package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "user:pass@tcp(0.0.0.0:13306)/test_db"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"` //ほんとはhash化してね
}

func main() {
	// データベースへの接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, pass FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Pass)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.ID, user.Name, user.Pass)
	}

}
