package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := app.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
}

func (app *App) Run(addr string) {}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
id SERIAL,
name TEXT NOT NULL,
company TEXT NOT NULL,
username TEXT NOT NULL,
password TEXT NOT NULL,
CONSTRAINT users_pkey PRIMARY KEY (id)
)`
