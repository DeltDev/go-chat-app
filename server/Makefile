postgresinit:
	docker run --name postgres18 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:18-alpine

postgres:
	docker exec -it postgres18 psql

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres18 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5433/go-chat?sslmode=disable" -verbose down
.PHONY: postgresinit postgres createdb dropdb migrateup migratedown