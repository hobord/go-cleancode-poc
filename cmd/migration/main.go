package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file:///database/mysql/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	m.Steps(2)
}
