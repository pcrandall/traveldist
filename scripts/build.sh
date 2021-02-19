#! /bin/bash
TODAY=$(date +'%F')
curDir=$(pwd)
releaseDir=$curDir/releases/$TODAY

go mod tidy

if [[ "$1" == "linux" ]]; then
# build for linux
    GOOS=linux GOARCH=386 go build .
else

#embed config.yml
go-bindata -o config.go config

# embed icon in the executable
rsrc -ico assets/favicon.ico

# build for windows
mkdir -p $releaseDir/config &> /dev/null
GOOS=windows GOARCH=386 go build .
cp ./config/config.yml $curDir/config
