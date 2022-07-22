include .env

postgres:
    docker run --name=innotaxi-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

createdb:
    docker exec -it innotaxi-db createdb postgres

dropdb:
    docker exec -it innotaxi-db dropdb postgres

migrateup:
    migrate -path internal/db -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
    migrate -path internal/db -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown