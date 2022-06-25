package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// https://github.com/go-sql-driver/mysql

var db *sql.DB
var dbx *sqlx.DB

func etude1() {
	rows, err := db.Query("SELECT * FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var birthday time.Time
		rows.Scan(&id, &name, &birthday)
		fmt.Println(id, name, birthday)
	}
}

type User struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Birthday time.Time `db:"birthday"`
}

type Comment struct {
	Id     int       `db:"id"`
	UserId string    `db:"user_id"`
	Body   string    `db:"body"`
	PostAt time.Time `db:"post_at"`
}

func etude2() {
	var users []*User
	err := dbx.Select(&users, "SELECT * FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func etude3() {
	var comments []*Comment
	err := dbx.Select(&comments, "SELECT * FROM comments ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func main() {
	var err error
	dsn := "sqlxetude:sqlxetude@(127.0.0.1:23306)/sqlxetude?parseTime=true"
	dbx, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbx.Close()
	db = dbx.DB

	etude1()
	etude2()
	etude3()
}
