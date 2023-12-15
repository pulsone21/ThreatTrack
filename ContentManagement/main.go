package main

import (
	"ContentManagement/api"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Setting up new Webserver")
	port := os.Getenv("BACKENDPORT")
	server := api.NewServer(port)
	server.Run()
	fmt.Println("Webserver started")
}
