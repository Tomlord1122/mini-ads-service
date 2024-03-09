package main

import (
	"backend-intern/api"
	db "backend-intern/db/sqlc"
	"backend-intern/util"
	"database/sql"
	"log"
	"time"

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
	conn.SetMaxOpenConns(1000000) // 設置最大開啟連接數
	conn.SetMaxIdleConns(500)     // 設置最大空閒連接數
	conn.SetConnMaxLifetime(time.Hour)
	query := db.New(conn)
	server := api.NewServer(query, config)
	server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
