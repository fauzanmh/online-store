# DATABASE
DB_USER=postgres
DB_PASSWORD=*r00t123*
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=online_store
DB_SSL=disable

# RUN PROGRAM
run:
	swag init && go run main.go

# RUN TESTING
test:
	go test -v -cover -coverprofile=cover.out ./unit_test

# DB MIGRATION
migrate-up:
	migrate -source file:./script/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up

migrate-down:
	migrate -source file:./script/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} down 1

migrate-drop:
	migrate -source file:./script/migration/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} drop
