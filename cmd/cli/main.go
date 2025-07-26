package main

import (
	"fmt"
	"net/http"
	"time"
)

const version = "v0.1.0"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, version)
	})
	http.ListenAndServe(":8080", nil)
}

func timer() {
	statusTicker := time.NewTicker(time.Hour)

	for range statusTicker.C {

	}
}
