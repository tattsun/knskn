.PHONY: run build

run:
	go run *.go

build:
	CGO_ENABLED=0 go build -v -trimpath -o knskn ./...
