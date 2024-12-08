package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var db *Db

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	err := db.InsertRequest(request)
	if err != nil {
		fmt.Printf("INSERT ERROR: %v\n", err)
	}
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

	println("Building db...")
	var err error
	db, err = NewDb("./webhole.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	http.HandleFunc("/{path...}", requestHandler)

	fmt.Printf("Listening on port %d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
