package app

import (
	"database/sql"
	"fmt"
	config2 "github.com/leandrocunha526/goapi-chat/app/config"
	_ "github.com/lib/pq"
	"log"
)

func DbConn() *sql.DB {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()
	config, err := config2.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbPort, config.DbName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
