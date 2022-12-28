#!/usr/bin/env bash

set -e

go test -coverprofile cover.out ./...
go tool cover -func=cover.out
