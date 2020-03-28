package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (db *sql.DB, err error) {
	conf := config.Database
	conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err = sql.Open("mysql", conStr)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	return
}

func addSubscribtion(chat int64, tags []string) error {
	db, err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()
	query := "INSERT INTO subscriptions(chat, tag) values (?, ?)"
	for _, t := range tags {
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Println(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(chat, t)

		if err != nil {
			log.Println(err)
		}
	}
	return err
}

func getTagsList() (tags []string, err error) {
	db, err := initDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT DISTINCT tag FROM subscriptions"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			log.Println(err)
		}
		tags = append(tags, tag)
	}
	return
}

func getChatByTag(tag string) (chats []int64, err error) {
	db, err := initDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT chat FROM subscriptions WHERE tag=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tag)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var chat int64
		err := rows.Scan(&chat)
		if err != nil {
			log.Println(err)
		}
		defer rows.Close()
		chats = append(chats, chat)
	}
	return
}
