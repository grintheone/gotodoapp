package main

import (
	"log"
	"net/http"
	"os"

	"github.com/grintheone/gotodoapp/internal/config"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Loads env variables into the system from .env file in the root directory
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	config := config.New()

	db, err := connectDB(config.DBUser, config.DBPass, config.DBUrl, config.DBName)

	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Println("Connected to db successfully")
	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     config.Host + config.Port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server started on port %s\n", config.Port)
	log.Fatal(srv.ListenAndServe())
}
