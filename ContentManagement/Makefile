.PHONY: help

build:
	@go build -o bin/contentManagement

setup: build
	@docker start ContentDB
	@sleep 3
	@./bin/contentManagement

run: build
	@./bin/contentManagement


generate_data: 
	@docker exec -i ContentDB mysql -u root -proot contentdb < ./devStuff/generate_data.sql
