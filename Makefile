GOBINDATA := $(shell command -v go-bindata 2> /dev/null)

currentDir = $(shell pwd)
releasesOutputDir = ${currentDir}/releases/$(date +'%F')
workOutputDir = ${WINDIR}/Projects/_sshClient/releases/$(date +'%F')/config

## installation
install:
ifndef GOBINDATA
	@echo "==> installing go-bindata"
	@go get -u github.com/go-bindata/go-bindata/...
endif
	@echo "==> installing go dependencies"
	@go mod download
.PHONY: install

embed:
	@echo "==> making embedded config"
	@go-bindata -o config.go config
.PHONY: embed

run:
	@echo "==> running travelDistance"
	@go run *.go
.PHONY: run

## @echo "OS not defined, Usage: make build windows"
build:
	@echo "==> building travelDistance"
	@${currentDir}/scripts/build.sh
.PHONY: build

git:
	@git add -u
	@git commit
	@git push origin
.PHONY: git

clean:
	@go clean --cache
	@go mod tidy
	@git clean -f
.PHONY: clean
