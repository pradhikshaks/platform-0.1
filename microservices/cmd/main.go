package main

import (
	"fmt"
	"microservices/server"
	"os"
)

func main() {
	if err := server.RestServer(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}