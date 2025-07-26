package main

import (
	"bufio"
	"fmt"
	"log"
	"path/filepath"

	"github.com/tarm/serial"
)

func main() {
	device := findTTYUSB()
	config := &serial.Config{
		Name: device, // Adjust if your device changes
		Baud: 9600,   // Must match the sender's baud rate
	}

	port, err := serial.OpenPort(config)
	if err != nil {
		log.Fatalf("Failed to open serial port: %v", err)
	}
	defer port.Close()

	reader := bufio.NewReader(port)
	fmt.Println("Listening on: ", device)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading: %v", err)
			continue
		}
		fmt.Printf("Received: %s", line)
	}
}

func findTTYUSB() string {
	paths, err := filepath.Glob("/dev/ttyUSB*")
	if err != nil || len(paths) == 0 {
		log.Fatal("No /dev/ttyUSB* devices found")
	}
	return paths[0] // Assumes first one is correct
}
