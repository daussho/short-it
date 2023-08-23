#!/bin/bash

run-http:
	@echo "Running HTTP server"
	@go run cmd/http/main.go
run-http-nodemon:
	@echo "Running HTTP server with nodemon"
	@nodemon -e go,html --signal SIGTERM --exec 'make run-http || exit 1' 
run-migration:
	@echo "Running migration"
	@go run cmd/migrations/main.go