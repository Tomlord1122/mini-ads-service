package main

import (
	"backend-intern/api"
	db "backend-intern/db/sqlc"
	"backend-intern/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // database/sql package does not have a PostgreSQL driver
)

func main() {
	config, err := util.LoadConfig(".") // Load config from current directory
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	query := db.New(conn)
	server := api.NewServer(query)
	server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
