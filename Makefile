GOBINDATA := $(shell command -v go-bindata 2> /dev/null)
currentDir := $(shell pwd)

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


writeRun:
	@echo "==> running travelDistance"
	@${currentDir}/scripts/unix/writeRun.sh
.PHONY: writeRun

run:
	@echo "==> running travelDistance"
	@${currentDir}/scripts/unix/run.sh
.PHONY: run

build:
	@echo "==> building travelDistance"
	@${currentDir}/scripts/unix/build.sh
.PHONY: build

runWindows:
	@echo "==> running WINDOWS travelDistance"
	@${currentDir}/scripts/windows/runWindows.bat
.PHONY: runWindows

runWindowsRest:
	@echo "==> running WINDOWS REST travelDistance"
	@${currentDir}/scripts/windows/runWindowsRest.bat
.PHONY: runWindowsRest

buildWindows:
	@echo "==> building WINDOWS travelDistance"
	@${currentDir}/scripts/windows/buildWindows.bat
.PHONY: buildWindows

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
