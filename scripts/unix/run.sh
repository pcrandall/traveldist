#! /bin/bash
go mod tidy
go-bindata -o config.go config
GOOS=windows GOARCH=386 go run .
