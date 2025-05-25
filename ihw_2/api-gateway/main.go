package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func newProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Ошибка парсинга URL %s: %v", target, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalPath := req.URL.Path
		originalDirector(req)
		log.Printf("[PROXY] %s %s -> %s%s", req.Method, originalPath, targetURL.String(), req.URL.Path)
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[ERROR] При проксировании запроса %s %s: %v", r.Method, r.URL.Path, err)
		http.Error(w, "Прокси-сервер вернул ошибку", http.StatusBadGateway)
	}

	return proxy
}

func logRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[REQUEST] %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("[DONE] %s %s за %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	storageProxy := newProxy("http://file-storage:8001")
	analyzerProxy := newProxy("http://file-analyzer:8002")

	router := mux.NewRouter()
	router.Use(logRequestsMiddleware)

	storageApiRouter := router.PathPrefix("/api/storage").Subrouter()
	storageApiRouter.Handle("/upload", storageProxy).Methods("POST")
	storageApiRouter.Handle("/files", storageProxy).Methods("GET")
	storageApiRouter.Handle("/files/{filename}", storageProxy).Methods("GET")
	storageApiRouter.Handle("/download/{filename}", storageProxy).Methods("GET")

	analysisApiRouter := router.PathPrefix("/api/analysis").Subrouter()
	analysisApiRouter.Handle("/results/{filename:.*}", analyzerProxy).Methods("GET")

	router.Handle("/upload", storageProxy).Methods("GET")
	router.Handle("/files", storageProxy).Methods("GET")
	router.Handle("/download/{filename}", storageProxy).Methods("GET")
	router.Handle("/analyze", analyzerProxy).Methods("GET")
	router.Handle("/analysis/results", analyzerProxy).Methods("GET")

	router.PathPrefix("/static/storage/").Handler(storageProxy)
	router.PathPrefix("/static/analysis/").Handler(analyzerProxy)

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	loggedRouter := handlers.LoggingHandler(log.Writer(), cors(router))

	log.Println("[INFO] API Gateway запущен на :8080")
	if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
