# Makefile
SPREADSHEET_ID := 'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
CREDENTIALS_PATH := '../../credentials/sample.json'

# default target
all: go-db-init

go-db-init: go-db-reset go-db-migrate go-seed

go-db-reset:
	migrate --path ../migrate/ddl --database 'mysql://proto:password@tcp(mysql:3306)/proto?charset=utf8mb4' -verbose down

go-db-migrate:
	migrate --path ../migrate/ddl --database 'mysql://proto:password@tcp(mysql:3306)/proto?charset=utf8mb4' -verbose up

go-seed: seed-movie-type seed-movie-format seed-movies

seed-movie-type:
	GOOGLE_SPREADSHEET_ID=$(SPREADSHEET_ID) GOOGLE_APPLICATION_CREDENTIALS=$(CREDENTIALS_PATH) go run ./movieType/main.go

seed-movie-format:
	GOOGLE_SPREADSHEET_ID=$(SPREADSHEET_ID) GOOGLE_APPLICATION_CREDENTIALS=$(CREDENTIALS_PATH) go run ./movieFormat/main.go

seed-movies:
	GOOGLE_SPREADSHEET_ID=$(SPREADSHEET_ID) GOOGLE_APPLICATION_CREDENTIALS=$(CREDENTIALS_PATH) go run ./movies/main.go
