package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
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

func (app *application) apiCheck() {
	fmt.Println("Check if version matches... we will pretend it does")

	if false {
		app.update()
	}
}

func (app *application) update() {
	updateURL := "fake-url"

	resp, err := http.Get(updateURL)
	if err != nil {
		fmt.Println("Failed to download updater")
		return
	}
	defer resp.Body.Close()

	// Create a temporary file for the updater
	tmpUpdaterPath := filepath.Join(os.TempDir(), "stream-sight-updater")
	outFile, err := os.Create(tmpUpdaterPath)
	if err != nil {
		fmt.Println("Failed to create temporary updater file")
		return
	}

	_, err = io.Copy(outFile, resp.Body)
	outFile.Close()
	if err != nil {
		fmt.Println("Failed to save updater")
		return
	}

	// Make it executable
	if err := os.Chmod(tmpUpdaterPath, 0755); err != nil {
		fmt.Println("Failed to mark updater as executable")
		return
	}

	// Run it
	cmd := exec.Command(tmpUpdaterPath, version)
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start updater:", err)
		return
	}
}
