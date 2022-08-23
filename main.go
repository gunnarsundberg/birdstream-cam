package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gunnarsundberg/birdstream-cam/stream"

	"github.com/gorilla/mux"
)

const (
	videoWebsocketURL = "/stream"
	port              = 8080
	width             = 1920
	height            = 1080
	fps               = 30
)

func main() {
	options := stream.CameraOptions{
		Width:          width,
		Height:         height,
		Fps:            fps,
		HorizontalFlip: false,
		VerticalFlip:   false,
		Rotation:       0,
		UseLibcamera:   true,
	}

	router := mux.NewRouter()

	// Websocket
	connectionNumber := make(chan int, 2)
	wsh := NewWebSocketHandler(connectionNumber)
	router.HandleFunc(videoWebsocketURL, wsh.Handler)
	go stream.Video(options, wsh, connectionNumber)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
