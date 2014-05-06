#!/bin/bash

pushd $(dirname $0) > /dev/null

pkill tactique_engine
GOPATH=$(pwd) go build -o tactique_engine engine.go
./tactique_engine > engine.log &

popd > /dev/null
