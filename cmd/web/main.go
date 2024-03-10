package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	host string = "localhost"
	port string = ":3000"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     host + port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
