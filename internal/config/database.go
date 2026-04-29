package config

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func openDB(pgString string) (*sql.DB, error) {

	db, err := sql.Open("pgx", pgString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectDb(pgString string) (*sql.DB, error) {
	conn, err := openDB(pgString)
	if err != nil {
		return nil, err
	}

	log.Println("connected to db")
	return conn, nil
}
