package main

import (
	"file-analyzer/db"
	"log"
	"net/http"

	"file-analyzer/handler"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()
	h, err := handler.NewAnalysisHandler(database)
	if err != nil {
		log.Fatalf("Failed to initialize handler: %v", err)
	}

	r := mux.NewRouter()
	h.RegisterRoutes(r)

	log.Println("File Analyzer service running on :8002")
	http.ListenAndServe(":8002", r)
}
