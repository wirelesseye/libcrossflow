package main

import (
	"libcrossflow/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Initialize()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	http.Handle("/", GetRouter())

	go func() {
		log.Fatal(http.ListenAndServe(":4331", nil))
	}()
	log.Print("Server Started on http://localhost:4331/")

	<-done
	log.Print("Server Stopped")
}
