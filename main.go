package main

import (
	"lesson11/internal/client"
	"lesson11/internal/server"
	"log"
	"time"
)

func main() {
	srv := server.NewServer()
	go func() {
		cl := client.NewClient()
		cl.RequestServer("http://127.0.0.1:8080/")
		cl.RequestServer("http://127.0.0.1:8080/time")
		<-time.After(2 * time.Second)
	}()

	if err := srv.Start("127.0.0.1", "8080"); err != nil {
		log.Fatal(err)
	}
}
