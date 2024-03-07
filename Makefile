postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

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