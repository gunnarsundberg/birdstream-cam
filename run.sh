#!/bin/bash

git pull
env GOOS=linux GOARCH=arm GOARM=7 go build
./birdstream-cam