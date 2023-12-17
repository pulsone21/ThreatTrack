.PHONY: help

include .env
export

build:
	@go get
	@go mod tidy
	@go build -o bin/frontEnd

run: build
	@./bin/frontEnd

	
backend: 
	@./../ContentManagement/bin/contentManagement