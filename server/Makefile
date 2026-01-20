include .env
export

postgresinit:
	docker run --name postgres18 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:18-alpine

postgres:
	docker exec -it postgres18 psql

createdb:
	docker exec -it postgres18 createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres18 dropdb go-chat

migrateup:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down

buildimage:
	docker build -t deltdev/go-chat-app:v"$(VERSION)" .

pushimage:
	docker push deltdev/go-chat-app:v"$(VERSION)"

pullimage:
	docker network create chat-net || true

	docker run -d --name chat-db \
		--network chat-net \
		-e POSTGRES_USER="$(PG_USER)" \
		-e POSTGRES_PASSWORD="$(PG_PASSWORD)" \
		-e POSTGRES_DB="$(PG_DB)" \
		postgres:18-alpine

	docker run -d --name my-chat-app \
		--network chat-net \
		-p 15015:8080 \
		-e DB_SOURCE="$(DB_SOURCE)" \
		-e JWT_SECRET="$(JWT_SECRET)" \
		-e CLIENT_ORIGINS="$(CLIENT_ORIGINS)" \
		deltdev/go-chat-app:v"$(VERSION)"

	docker update --restart always chat-db
	docker update --restart always my-chat-app

	sleep 5
	docker exec my-chat-app migrate -path db/migrations -database "$(DB_SOURCE)" up
.PHONY: postgresinit postgres createdb dropdb migrateup migratedown buildimage pushimage