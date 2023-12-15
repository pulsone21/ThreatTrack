package main

import (
	"fmt"
	"frontend/gui"
	"os"
)

func main() {
	fmt.Println("Setting up new Webserver")
	port := os.Getenv("FRONTENDPORT")
	backendPort := os.Getenv("BACKENDPORT")
	server := gui.CreateServer("localhost:"+port, "http://localhost:"+backendPort)
	server.RunServer()
}
