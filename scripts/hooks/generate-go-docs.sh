#!/usr/bin/env bash

cd ./app/internal || exit
doc2go -config ./doc2go.rc  ./...
