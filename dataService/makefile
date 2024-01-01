.PHONY: help

include .env
export

build:
	@go get
	@go mod tidy
	@go build -o bin/data_service

setup: build start_db reset_db
	@./bin/backEnd

run: build start_db
	@./bin/data_service

restart: start_db reset_db setup_db generate_data run

test: start_db reset_db setup_db generate_data create_test_img
	@docker start TestBackend
	@go test -run ^TestSuite$  -v ./tests

prep_db: start_db reset_db setup_db generate_data

create_test_img:
	@docker rm --force TestBackend
	@docker build -t backend:test -f Dockerfile.testing .
	@docker create -p 5666:5666 --name TestBackend --network DB backend:test 

start_db:
	@echo "starting db"
	@docker start ContentDB
	@sleep 3
	@echo "db up and running"

generate_data: 
	@echo "generating data"
	@docker exec -i ContentDB mysql -u root -proot < ./devStuff/generate_data.sql
	@echo "data generated"

reset_db: 
	@echo "reseting db"
	@docker exec -i ContentDB mysql -u root -proot < ./devStuff/reset_db.sql
	@echo "db reseted"

setup_db:
	@echo "setting db up"
	@docker exec -i ContentDB mysql -u root -proot < ./devStuff/setup_db.sql
	@echo "db setuped"