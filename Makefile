SHELL := /bin/bash

run: build game_on

build:
	CGO_ENABLED=0 go build

game_on:
	./wordle

game_debug:
	./wordle -debug


# testing
test: unit_test

unit_test:
	go test ./... -v -cover

unit_test_main:
	go test -v -cover


# run without building
go_run_dir:
	go run .

go_run_module:
	go run github.com/andrew-j-price/wordle


# cleanup items
cleanup: delete_binaries delete_logs

delete_binaries:
	rm -f ./worlde

delete_logs:
	rm -f ./log.txt
