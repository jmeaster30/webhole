package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("REQUEST GOT!!! %s from %s to [%s] %s %s\n", request.Proto, request.RemoteAddr, request.Method, request.Host, request.URL.Path)
	http.NotFound(writer, request)
}

func main() {
	var port uint
	flag.UintVar(&port, "port", 8000, "The port to listen for requests on")
	flag.Parse()

	if port >= 65536 {
		log.Fatalf("Invalid port number '%d'", port)
	}

	println("Hello WebHole!!")

	http.HandleFunc("/{path...}", requestHandler)

	fmt.Printf("Listening on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
