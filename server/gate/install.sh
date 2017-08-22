#!/usr/bin/env bash

export GOPATH=$(pwd)/../..

export GOOS=linux
export GOARCH=amd64
go install -ldflags "-s -w"