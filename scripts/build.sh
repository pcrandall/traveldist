#! /bin/bash

go mod tidy

curDir=$(pwd)
today=$(date +'%F')
releaseDir=$curDir/releases/$today
mkdir -p $releaseDir

if [[ "$1" == "linux" ]]; then
# build for linux
    GOOS=linux GOARCH=386 go build .
else
    #embed config.yml
    go-bindata -o config.go config
    # embed icon in the executable
    rsrc -ico assets/favicon.ico
    # build for windows
    mkdir -p $releaseDir/config
    GOOS=windows GOARCH=386 go build .
    cp -r config ${releaseDir}/
    cp travelDist.exe ${releaseDir}/Travel\ Distances.exe
    rm travelDist.exe
fi
