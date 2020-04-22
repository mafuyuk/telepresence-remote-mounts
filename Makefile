all: build

.PHONY: build
build:
	@GOARCH=amd64 GOOS=linux go build -o ./bin/telepresence-remote-mounts main.go
