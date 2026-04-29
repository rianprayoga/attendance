package main

import (
	"flag"
	"lentera/internal/config"
	"lentera/internal/repository"
	"lentera/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	var pgString string

	flag.StringVar(&pgString, "pgString", "host=localhost port=5432 user=postgres password=postgres dbname=lentera timezone=UTC", "Postgres connection string")

	conn, err := config.ConnectDb(pgString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.PgRepo{
		DB: conn,
	}

	ge := gin.Default()
	routes.SetupRoutes(ge, repo)
	err = ge.Run(":8080")
	if err != nil {
		panic(err)
	}
}
