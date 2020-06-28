package main

import (
	"database/sql"
	"git.rappet.de/rappet/netmap-api/repo"
	"git.rappet.de/rappet/netmap-api/repo_db"
	"git.rappet.de/rappet/netmap-api/repo_mock"
	"git.rappet.de/rappet/netmap-api/routing"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var ptrRepository repo.PtrRepository = repo_mock.NewPtrRepository()

func main() {
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		connStr = "user=netmap dbname=netmap password=netmap sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	ptrRepository = repo_db.NewPtrRepository(db)

	log.Println("starting netmap-api")

	router := routing.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
