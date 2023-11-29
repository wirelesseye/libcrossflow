package main

import (
	"libcrossflow/config"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.LoadConfig()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	http.Handle("/", GetRouter())

	go func() {
		log.Fatal(http.ListenAndServe(":4331", nil))
	}()
	localIP := GetOutboundIP()
	log.Printf("Server Started on http://localhost:4331/ or http://%s:4331/", localIP)

	<-done
	log.Print("Server Stopped")
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
