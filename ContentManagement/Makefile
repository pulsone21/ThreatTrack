.PHONY: help

include .env
export

build:
	@go get
	@go mod tidy
	@go build -o bin/backEnd

setup: build start_db reset_db
	@./bin/backEnd

run: build
	@./bin/backEnd

restart: start_db reset_db setup_db generate_data run

test: start_db reset_db setup_db generate_data
	@go test -run ^TestSuite$  -v ./tests

start_db:
	@echo "starting db"
	@docker start ContentDB
	@sleep 3
	@echo "db up and running"

generate_data: 
	@echo "generating data"
	@docker exec -i ContentDB mysql -u root -proot contentdb < ./devStuff/generate_data.sql
	@echo "data generated"

reset_db: 
	@echo "reseting db"
	@docker exec -i ContentDB mysql -u root -proot contentdb < ./devStuff/reset_db.sql
	@echo "db reseted"

setup_db:
	@echo "setting db up"
	@docker exec -i ContentDB mysql -u root -proot < ./devStuff/setup_db.sql
	@echo "db setuped"