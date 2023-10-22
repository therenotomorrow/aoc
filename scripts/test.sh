#!/usr/bin/env bash

set -e

case "$1" in
  -fast)
    go test ./...
    ;;
  -race)
    go test -race ./...
    ;;
  -cover)
    go test -coverprofile cover.out -cover ./...
    go tool cover -func cover.out | grep total:
    ;;
esac
