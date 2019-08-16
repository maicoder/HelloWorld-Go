package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/time/", func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		writer.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8888", nil)
}
