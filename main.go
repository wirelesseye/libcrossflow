package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	handler := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(":4331", nil))
	}()
	log.Print("Server Started on http://localhost:4331/")

	<-done
	log.Print("Server Stopped")
}
