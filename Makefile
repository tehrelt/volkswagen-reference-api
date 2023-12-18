.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: run
run:
	make
	./app.exe

.DEFAULT_GOAL := build