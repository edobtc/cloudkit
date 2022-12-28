#!/usr/bin/env bash

## Builds zip file artifacts for the lambda functions

set -e

buildBasePath="build/artifacts/functions"

function build() {
    local buildPath="$1"

    local srcPath="events/aws/lambda/artifacts"

    echo $buildPath
    GOARCH=amd64 GOOS=linux go build -o ${buildBasePath}/${buildPath}/main ${srcPath}/${buildPath}/main.go
    pushd ${buildBasePath}/${buildPath}
    zip function.zip main
    rm main
    popd
}

rm -rf ${buildBasePath}

build "bitcoind/rpc/getblock"
build "bitcoind/rpc/gettransaction"
build "status"
