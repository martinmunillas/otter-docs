
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/martinmunillas/otter/migrate"
	"github.com/martinmunillas/otter/log"
	"github.com/martinmunillas/otter/env"
	
	_ "github.com/martinmunillas/otter-docs/migrations"
)

func main() {
	dbUser := env.RequiredStringEnvVar("DB_USER")
	dbName := env.RequiredStringEnvVar("DB_NAME")
	dbPassword := env.RequiredStringEnvVar("DB_PASSWORD")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	logger := log.NewLogger(false)
	err = migrate.RunAll(db, logger)
	if err != nil {
		logger.Error(err.Error())
	}
}
