package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	version  = "v0.1.0"
	interval = 15 * time.Minute
)

type application struct {
	count int
}

func main() {

	app := application{}
	go app.timer()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, app.version())
	})
	http.ListenAndServe(":8080", nil)
}

func (app *application) timer() {
	statusTicker := time.NewTicker(interval)

	for range statusTicker.C {
		app.count++
		app.apiCheck()
	}
}

func (app *application) version() string {
	return version + "." + strconv.Itoa(app.count)
}

func (app *application) apiCheck() {
	fmt.Println("Place Holder Function")
}
