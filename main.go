package main

import (
	"libcrossflow/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	resDir, success := os.LookupEnv("RES_PATH")
	if !success {
		ex, _ := os.Executable()
		exPath := filepath.Dir(ex)
		resDir = filepath.Join(exPath, "res")
	}

	handler := http.FileServer(http.Dir(resDir))
	http.HandleFunc("/api", api.APIHandler)
	http.Handle("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(":4331", nil))
	}()
	log.Print("Server Started on http://localhost:4331/")

	<-done
	log.Print("Server Stopped")
}
