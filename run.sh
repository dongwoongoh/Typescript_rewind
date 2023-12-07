#!/bin/zsh

echo "Running Go Web Application by MAD"
go test ./... -v
go run cmd/main.go
