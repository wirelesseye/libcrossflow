package main

import (
	"fmt"
	"libcrossflow/config"
	"libcrossflow/util"
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

	localIP := getOutboundIP()
	localAddr := fmt.Sprintf("http://%s:4331/", localIP)
	log.Printf("Server Started on http://localhost:4331/ or %s", localAddr)
	util.PrintQRCode(localAddr)

	go func() {
		handleInput()
	}()

	<-done
	log.Print("Server Stopped")
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func handleInput() {
	var input string
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading stdin:", err)
			break
		}

		fmt.Println("Unknown command:", input)
	}
}
