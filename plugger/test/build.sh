#!/bin/bash

go build
cd plugin; go build -buildmode=plugin -o test.so test.go
