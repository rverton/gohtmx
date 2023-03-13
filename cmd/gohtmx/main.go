package main

import (
	"gohtmx"
	"log"
)

func main() {

	server := gohtmx.NewServer(":8080")
	if err := server.Start(); err != nil {
		log.Fatal("cant start server", err)
	}

}
