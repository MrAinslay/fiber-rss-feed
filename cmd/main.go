package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/joho/godotenv"
)

func initDatabase() {
	godotenv.Load("Kkey.env")
	db, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}
	config.DBQueris = config.New(db)
}
