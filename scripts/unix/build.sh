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
    mkdir -p $releaseDir/{config,db}
    GOOS=windows GOARCH=386 go build .
    cp -r config ${releaseDir}/
    cp -r db ${releaseDir}/
    cp travelDist.exe ${releaseDir}/Travel\ Distances.exe
    rm travelDist.exe
    cp -r ${releaseDir} /mnt/c/Users/Phillip.Crandall/Desktop/traveldist
    cd /mnt/c/Users/Phillip.Crandall/Desktop/traveldist
    explorer.exe .
fi
