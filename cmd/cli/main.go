package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tarm/serial"
)

const (
	version  = "v0.1.0"
	interval = 1 * time.Minute
)

type application struct {
	count int
}

func main() {

	app := application{}
	go app.timer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.writeSerial(app.version())
		fmt.Fprintln(w, app.version())
	})
	http.ListenAndServe(":8080", nil)
}

func (app *application) timer() {
	statusTicker := time.NewTicker(interval)

	for range statusTicker.C {
		app.count++
		app.writeSerial(app.version())
		app.apiCheck()
	}
}

func (app *application) version() string {
	return version + "." + strconv.Itoa(app.count)
}

func (app *application) apiCheck() {
	fmt.Println("Place Holder Function")
}

func (app *application) writeSerial(message string) error {
	config := &serial.Config{
		Name: "/dev/ttyS0",
		Baud: 9600,
	}
	port, err := serial.OpenPort(config)
	if err != nil {
		return err
	}
	defer port.Close()

	_, err = port.Write([]byte(message + "\n"))
	return err
}
