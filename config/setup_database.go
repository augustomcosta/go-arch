package config

import (
	"database/sql"
	db "github.com/augustomcosta/go-arch/sql/sqlc_gen"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

func ConnectToDB() *sql.DB {
	dbcon, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatal("Error connecting to the db. Error: ", err)
	}

	err = dbcon.Ping()
	if err != nil {
		log.Fatal("Database connection failed. Error: ", err)
	}

	slog.Info("Database connection established")

	return dbcon
}

func NewDBQueries(dbConnection *sql.DB) *db.Queries {
	return db.New(dbConnection)
}
