#!/bin/sh

GOOS=linux GOARCH=amd64 go build -o dist/stabl main.go
GOOS=darwin GOARCH=amd64 go build -o dist/stabl_darwin main.go
GOOS=windows GOARCH=amd64 go build -o dist/stabl.exe main.go
