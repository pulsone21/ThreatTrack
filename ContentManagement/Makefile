.PHONY: help

build:
	@go get
	@go mod tidy
	@go build -o bin/backEnd

setup: build
	@docker start ContentDB
	@sleep 3
	@./bin/backEnd

run: build
	@./bin/backEnd


generate_data: 
	@docker exec -i ContentDB mysql -u root -proot contentdb < ./devStuff/generate_data.sql
