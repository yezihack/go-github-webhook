#!/usr/bin/env bash

dev:run

fmt:
	gofmt -l -w ./

build:fmt
	go build -o bin/go_hooks .

clean:
	rm -rf bin/*

run:clean build
	bin/go_hooks
