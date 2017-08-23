#!/usr/bin/env bash

if [ -d "$1" ]; then
    cd $1
    sh install.sh
    cd ..
fi
