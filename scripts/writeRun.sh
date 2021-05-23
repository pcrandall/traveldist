#! /bin/bash
go mod tidy
GOOS=windows GOARCH=386 go run . -w
