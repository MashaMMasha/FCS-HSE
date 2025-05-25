package main

import (
	"log"
	"net/http"
	"os"

	"file-storage/db"
	"file-storage/handler"

	"github.com/gorilla/mux"
)

func main() {
	os.MkdirAll("./uploads", 0755)
	os.MkdirAll("./static/css", 0755)
	os.MkdirAll("./static/js", 0755)
	os.MkdirAll("./templates", 0755)
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	r := mux.NewRouter()
	fileHandler, err := handler.NewFileHandler(database)
	if err != nil {
		log.Fatalf("Failed to create fileHandler: %v", err)
	}
	fileHandler.RegisterRoutes(r)

	log.Println("Сервер запущен на :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
