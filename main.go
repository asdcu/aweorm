package main

import (
	"aweorm/log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Account struct {
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("sqlite3", "awe.db")
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		_ = db.Close()
	}()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) VALUES (?),(?)", "Tom", "Sam")
	if err != nil {
		log.Error(err)
	} else {
		affected, _ := result.RowsAffected()
		log.Info(affected)
	}
	row := db.QueryRow("SELECT Name from User Limit 1")
	var name string
	if err = row.Scan(&name); err != nil {
		log.Error(err)
		return
	}
	log.Info(name)
}
