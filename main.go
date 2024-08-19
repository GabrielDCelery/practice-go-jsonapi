package main

import (
	"log"
)

func main() {
	err, store := NewPostgresStore()
	if err != nil {
		log.Fatalln(err)
	}
	server := NewServer(":3000", store)
	server.Run()
}
