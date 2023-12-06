package main

import (
	"fmt"
	"lesson11/internal/client"
	"lesson11/internal/models"
	"lesson11/internal/server"
	"log"
	"time"
)

func main() {
	srv := server.NewServer()
	go func() {
		cl := client.NewClient()
		<-time.After(1 * time.Second)
		resp, err := cl.Request("GET", "http://127.0.0.1:8080/", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))

		resp, err = cl.Request("GET", "http://127.0.0.1:8080/time", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))

		user := models.NewPerson("Alex", 20)
		resp, err = cl.Request("POST", "http://127.0.0.1:8080/user", user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))
		resp, err = cl.Request("GET", "http://127.0.0.1:8080/user", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))
	}()

	if err := srv.Start("127.0.0.1", "8080"); err != nil {
		log.Fatal(err)
	}
}
