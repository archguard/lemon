#!/usr/bin/env bash
mkdir -p bin
GOOS=linux GOARCH=amd64 go build  -o bin/lemon_cli github.com/newlee/lemon/cli

