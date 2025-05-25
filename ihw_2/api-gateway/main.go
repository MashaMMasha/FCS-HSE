package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func newProxy(target string, pathRewrite string) *httputil.ReverseProxy {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Path = strings.Replace(req.URL.Path, pathRewrite, "", 1)
		log.Printf("Proxying: %s -> %s%s", req.URL.Path, targetURL, req.URL.Path)
	}
	return proxy
}

func main() {
	storageProxy := newProxy("http://file-storage:8001", "/api")
	analyzerProxy := newProxy("http://file-analyzer:8002", "/api")

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		storageProxy.ServeHTTP(w, r)
	}).Methods("POST")

	apiRouter.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		storageProxy.ServeHTTP(w, r)
	}).Methods("GET")

	apiRouter.HandleFunc("/analyze", func(w http.ResponseWriter, r *http.Request) {
		analyzerProxy.ServeHTTP(w, r)
	}).Methods("POST")

	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		storageProxy.ServeHTTP(w, r)
	}).Methods("GET")

	router.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		storageProxy.ServeHTTP(w, r)
	}).Methods("GET")

	router.PathPrefix("/static/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		storageProxy.ServeHTTP(w, r)
	})

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	loggedRouter := handlers.LoggingHandler(log.Writer(), router)

	log.Println("API Gateway запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", cors(loggedRouter)))
}
