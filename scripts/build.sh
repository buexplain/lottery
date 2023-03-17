#!/bin/bash

# 编译Windows版本
echo build Windows version
mkdir -p ./../build/win
rm -rf ./../build/win/*
mkdir -p ./../build/win/configs
export GOARCH=amd64
export GOOS=windows
go build -trimpath -ldflags "-s -w" -o ./../build/win/lottery-win-amd64.exe ./../cmd/lottery.go
cp ./../configs/lottery.example.toml ./../build/win/configs/lottery.toml

# 编译Linux版本
echo build Linux version
mkdir -p ./../build/linux
rm -rf ./../build/linux/*
mkdir -p ./../build/linux/configs
export GOARCH=amd64
export GOOS=linux
go build -trimpath -ldflags "-s -w" -o ./../build/linux/lottery-linux-amd64.bin ./../cmd/lottery.go
cp ./../configs/lottery.example.toml ./../build/linux/configs/lottery.toml

# 编译Darwin版本
echo build Darwin version
mkdir -p ./../build/darwin
rm -rf ./../build/darwin/*
mkdir -p ./../build/darwin/configs
export GOARCH=amd64
export GOOS=darwin
go build -trimpath -ldflags "-s -w" -o ./../build/darwin/lottery-darwin-amd64.bin ./../cmd/lottery.go
cp ./../configs/lottery.example.toml ./../build/darwin/configs/lottery.toml

echo build successfully