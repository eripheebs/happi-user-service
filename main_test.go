package main_test

import (
	"log"
	"os"
	"testing"

	"."
	"github.com/joho/godotenv"
)

var app main.App

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app = main.App{}
	app.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
id SERIAL,
name TEXT NOT NULL,
company TEXT NOT NULL,
username TEXT NOT NULL,
password TEXT NOT NULL,
CONSTRAINT users_pkey PRIMARY KEY (id)
)`

func clearTable() {
	app.DB.Exec("DELETE FROM users")
	app.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}
