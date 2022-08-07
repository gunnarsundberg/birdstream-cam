# birdstream-cam
Livestream camera from Raspberry Pi to browser using Go and Websockets

## Introduction

Stream H264 video directly from a Raspberry Pi to a browser using Go, Websockets, and libcamera-vid. 
This is the camera portion of Birdstream, an open-source ML-assisted birding application.
This repository borrows heavily from [bezineb5's go-h264-streamer](https://github.com/bezineb5/go-h264-streamer)

## Run

The easiest way to get started is by running:

```sh
# Raspberry Pi 2 and more recent (ARM7)
env GOOS=linux GOARCH=arm GOARM=7 go build
# Raspberry Pi 1 and Zero (ARM6)
env GOOS=linux GOARCH=arm GOARM=6 go build
# Run executable
./birdstream-cam
```
### Run as a Service

If you would like for the camera stream to always be on, use the included service file, 
which pulls updates, builds, and runs on every startup.

```sh
sudo ln-s /home/pi/birdstream-cam/birdstream.service /etc/systemd/system

sudo systemctl enable batterylevel.service
sudo systemctl start batterylevel.service
```