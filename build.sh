#!/bin/bash

# 编译为 Linux 平台
env GOOS=linux GOARCH=amd64 go build -o mobile_storage_test_service_linux

# 编译为 Windows 平台
env GOOS=windows GOARCH=amd64 go build -o mobile_storage_test_service_win.exe

# 编译为 macOS 平台
env GOOS=darwin GOARCH=arm64 go build -o mobile_storage_test_service_mac
