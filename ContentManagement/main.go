package main

import (
	"ContentManagement/api"
	"fmt"
)

func main() {
	fmt.Println("Setting up new Webserver")
	server := api.NewServer()
	server.Run()
	fmt.Println("Webserver started")
}
