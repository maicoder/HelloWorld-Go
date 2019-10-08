package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

type myHandler struct {}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "HelloWorld!\n")
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	/*
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8888", router))

	 */
	mux := http.NewServeMux()

	h := new(myHandler)
	mux.Handle("/foo", h)
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
