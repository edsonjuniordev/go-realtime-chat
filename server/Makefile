postgresinit:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

postgres:
	docker exec -it postgres psql

create-database:
	docker exec -it postgres createdb --username=root --owner=root go-chat

drop-database:
	docker exec -it postgres dropdb go-chat

migrate-up:
	migrate -path infra/database/migrations -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" -verbose up

migrate-down:
	migrate -path infra/database/migrations -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" -verbose down
