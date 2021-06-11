set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
go mod tidy
go run . -r

pause
