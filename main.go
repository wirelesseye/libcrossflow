package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func listFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dirname, _ := os.UserHomeDir()
	files, _ := os.ReadDir(dirname)

	var filenames []string
	for _, e := range files {
		filenames = append(filenames, e.Name())
	}

	res, _ := json.Marshal(filenames)
	w.Write(res)
}

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
	http.HandleFunc("/api", listFiles)
	http.Handle("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(":4331", nil))
	}()
	log.Print("Server Started on http://localhost:4331/")

	<-done
	log.Print("Server Stopped")
}
