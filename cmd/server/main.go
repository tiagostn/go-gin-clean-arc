package main

import (
	"fmt"

	"github.com/tiagostn/go-gin-clean-arc/internal"
)

// @title Go api using Gin with clean architecture
// @version 1.0
// @description Example Golang REST API
// @contact.name Tiago Sant'Ana
// @contact.url https://github.com/tiagostn
// @contact.email tiago.santana@gmail.com
// @BasePath /
func main() {
	fmt.Println("Golang Gin Server 👌")
	server := internal.Server()
	server.Run()
}
