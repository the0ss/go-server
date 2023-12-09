package main

import (
	"backend/internal/routes"
	"fmt"
)

func main() {
	fmt.Println("Yeah Buddy!")
	server := routes.NewAPIServer(":8000")
	server.Run()
}
