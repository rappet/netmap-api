package main

import (
	"database/sql"
	"git.rappet.de/rappet/netmap-api/repo"
	"git.rappet.de/rappet/netmap-api/repo_db"
	"git.rappet.de/rappet/netmap-api/repo_mock"
	"git.rappet.de/rappet/netmap-api/routing"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

var ptrRepository repo.PtrRepository = repo_mock.NewPtrRepository()

func openDb() *sql.DB {
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		connStr = "postgres://netmap:netmap@localhost?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func doMigrations(db *sql.DB)  {
	log.Println("starting database migrations")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Steps(1)
}

func main() {
	db := openDb()
	doMigrations(db)

	ptrRepository = repo_db.NewPtrRepository(db)

	log.Println("starting netmap-api")

	router := routing.NewRouter(routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
