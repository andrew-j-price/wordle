SHELL := /bin/bash

run: build game_on

# run without building
go_run_gui:
	go run ./cmd/gui/

go_run_terminal:
	go run ./cmd/terminal/

go_run_module:
	go run github.com/andrew-j-price/wordle/cmd/terminal


# build and run
build:
	CGO_ENABLED=0 go build -o ./wordle ./cmd/terminal/

build_all:
	go build ./...

game_on:
	./wordle

game_debug:
	./wordle -debug


# testing
test: unit_test

unit_test:
	go test ./... -v -cover

unit_test_game:
	go test ./pkg/game -v -cover


# cleanup items
cleanup: delete_binaries delete_logs

delete_binaries:
	rm -f ./worlde

delete_logs:
	rm -f ./log.txt
