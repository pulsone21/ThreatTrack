package main

import (
	"fmt"
	"frontend/gui"
	"os"
)

func main() {
	fmt.Println("Setting up FrontEnd")
	port := os.Getenv("FRONTEND_PORT")
	backendAdress := fmt.Sprintf("%s:%s", os.Getenv("BACKEND_ADRESS"), os.Getenv("BACKEND_PORT"))
	server := gui.CreateServer(port, backendAdress)
	fmt.Println("trying to start frontEnd Service")
	fmt.Println(server.Port)
	server.Run()
}
