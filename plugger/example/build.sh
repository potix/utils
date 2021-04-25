#!/bin/bash

go get -u
go build 

cd plugin
go get -u 
go build -buildmode=plugin -o test.so test.go
cd ..
