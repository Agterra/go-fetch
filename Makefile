SRC := $(shell find . -type f -name '*.go' -print)

.PHONY: build
build:
	go build -o ./bin/fetch ./cmd/fetch

.PHONY: run
run:
	make build
	./bin/fetch

# build:
# 	GOARCH=wasm GOOS=js go build -o web/app.wasm internal/*
# 	go build -o cmd/fetch internal/*

# run: build
# 	./cmd/fetch

