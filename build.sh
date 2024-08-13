#!/bin/sh

ver=1.0.1

GOOS=linux GOARCH=amd64 go build -o dist/stabl_linux-$ver main.go
GOOS=windows GOARCH=amd64 go build -o dist/stabl_windows-$ver.exe main.go
GOOS=darwin GOARCH=amd64 go build -o dist/stabl_mac-$ver main.go
