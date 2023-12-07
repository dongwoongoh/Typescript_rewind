#!/bin/zsh

echo "Running Go Web Application by MAD"
go test ./...
go run cmd/main.go
