# Dcard-Backend-Intern

# Ads Management API

## Overview

This project is 2024 Dcard backend intern assignment.

## Features

- **Create Ads**: Add new advertisements with detailed targeting options.
- **List Ads**: Retrieve active ads with flexible filtering, optimized for performance.
- **Cache Integration**: Leverages Redis for caching, significantly reducing response times for frequent queries.
- **Database Optimization**: Utilizes indexed queries for efficient data retrieval.
- **Scalability**: Designed for high throughput, supporting auto-scaling and load balancing.

## Getting Started

### Prerequisites

- Go 1.15+
- PostgreSQL 12+
- Redis 6+

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Tomlord1122/Dcard-Backend-Intern.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up the environment variables as per the `.env` file.

### Running the API

1.  Start the PostgreSQL and Redis servers.
2.  Run the following command to start the API server via Makefile:

    ```bash
         postgres:
    docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
    ```

    ```

         stop:
         docker stop postgres16

         start:
         docker start postgres16

         createdb:
         docker exec -it postgres16 createdb --username=root --owner=root dcard

         dropdb:
         docker exec -it postgres16 dropdb dcard

         migrateup:
         migrate -path db/migration -database "postgresql://root:secret@localhost:5432/dcard?sslmode=disable" -verbose up
         migratedown:
         migrate -path db/migration -database "postgresql://root:secret@localhost:5432/dcard?sslmode=disable" -verbose down
         server:
         go run main.go
         test:
         go test -v -cover ./...
         sqlc:
         sqlc generate
         k6:
         cd k6 && k6 run loadtest.js
         redis:
         redis-server ./redis.conf

         .PHONY: postgres createdb dropdb migrateup migratedown sqlc k6 test server stop start
    ```

## API Endpoints

- **POST /ads**: Create a new ad.
- **GET /ads**: List ads based on query parameters.

## Load test

![](/asset/test.png)
第一次測試

![](/asset/test2.png)
第二次測試
