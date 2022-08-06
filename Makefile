.PHONY: run build build-arm

run:
	go run *.go

build:
	CGO_ENABLED=0 go build -v -trimpath -o knskn *.go

build-arm:
	GOOS=linux GOARCH=arm go build -v -trimpath -o knskn *.go