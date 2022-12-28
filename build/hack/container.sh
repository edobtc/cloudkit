#!/usr/bin/env bash

docker build --platform linux/amd64 -t cloudkit:latest -f ./build/docker/Dockerfile .
