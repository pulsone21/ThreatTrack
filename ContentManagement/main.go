package main

import (
	"ContentManagement/api"
	"fmt"
)

func main() {
	fmt.Println("Setting up new Webserver")
	server := api.NewServer("8080")
	server.Run()
	fmt.Println("Webserver started")
}
