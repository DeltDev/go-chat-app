
DB_URL=postgresql://root:password@localhost:5433/go-chat?sslmode=disable
VERSION=1
postgresinit:
	docker run --name postgres18 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:18-alpine

postgres:
	docker exec -it postgres18 psql

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres18 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

buildimage:
	docker build -t deltdev/go-chat-app:v"$(VERSION)" .

pushimage:
	docker push deltdev/go-chat-app:v"$(VERSION)"
pullimage:
	docker network create chat-net
	docker run -d --name chat-db --network chat-net -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=go-chat postgres:18-alpine
	docker run -d --name my-chat-app --network chat-net -p 15015:8080 -e DB_SOURCE="postgresql://root:password@chat-db:5432/go-chat?sslmode=disable" deltdev/go-chat-app:v"$(VERSION)"
	docker update --restart always chat-db
	docker update --restart always my-chat-app

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown buildimage pushimage