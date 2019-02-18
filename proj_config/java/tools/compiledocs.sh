#!/bin/bash

if [[ "$(ls)" != "" ]]; then
    rm -r docs/*
fi

pushd src/
find -name "*.java" | xargs javadoc -d ../docs
popd