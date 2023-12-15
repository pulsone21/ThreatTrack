package main

import (
	"fmt"
	"threat_track/frontend/gui"
)

func main() {
	fmt.Println("Setting up new Webserver")
	server := gui.CreateServer("localhost:5051", "http://localhost:8080")
	server.RunServer()
}
