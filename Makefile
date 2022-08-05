.PHONY: run build

run:
	go run *.go

build:
	go build -v -o knskn ./...