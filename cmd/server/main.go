package main

import (
	"fmt"

	"github.com/tiagostn/go-gin-clean-arc/internal"
)

func main() {
	fmt.Println("Golang Gin Server 👌")
	server := internal.Server()
	server.Run()
}
