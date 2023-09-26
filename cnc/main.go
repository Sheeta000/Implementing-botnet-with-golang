package main

import (
	"Xone/cnc"
)

func main() {
	server := cnc.NewServer()
	server.Start()
}

//CGO_ENABLED=0;GOOS=linux;GOARCH=amd64
