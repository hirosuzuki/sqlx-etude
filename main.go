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
	UserId int       `db:"user_id"`
	Body   string    `db:"body"`
	PostAt time.Time `db:"post_at"`
}

type UserComment struct {
	Comment
	User
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

func etude4() {
	var rows []*UserComment
	err := dbx.Select(&rows, "SELECT c.*, u.* FROM comments c INNER JOIN users u ON u.id = c.user_id ORDER BY c.id")
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows {
		row.User.Id = row.Comment.UserId
	}
	for _, row := range rows {
		fmt.Println(row)
	}
}

func etude5() {
	var comments []*Comment
	err := dbx.Select(&comments, "SELECT * FROM comments ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	ids := make([]int, 0)
	for _, comment := range comments {
		ids = append(ids, comment.UserId)
	}
	query, params, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)
	if err != nil {
		log.Fatal(err)
	}
	var users []*User
	err = dbx.Select(&users, query, params...)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
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
	etude4()
	etude5()
}
