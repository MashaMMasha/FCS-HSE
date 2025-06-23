package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Gateway struct {
	ordersServiceURL   string
	paymentsServiceURL string
}

func NewGateway() *Gateway {
	return &Gateway{
		ordersServiceURL:   "http://orders-service:8001",
		paymentsServiceURL: "http://payments-service:8002",
	}
}

func (g *Gateway) proxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) {
	log.Printf("[INFO] Proxying request %s %s -> %s", r.Method, r.URL.Path, targetURL)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Reading body: %v", err)
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	r.Body.Close()

	req, err := http.NewRequest(r.Method, targetURL, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("[ERROR] Creating request: %v", err)
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Executing request to %s: %v", targetURL, err)
		http.Error(w, "Error making request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("[INFO] Response from %s: %d %s", targetURL, resp.StatusCode, http.StatusText(resp.StatusCode))

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] Reading response body: %v", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	w.Write(respBody)
}

func (g *Gateway) routeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] Incoming request: %s %s", r.Method, r.URL.Path)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		log.Printf("[INFO] Handled preflight OPTIONS for %s", r.URL.Path)
		return
	}

	path := r.URL.Path

	if strings.HasPrefix(path, "/order") {
		targetURL := g.ordersServiceURL + path
		g.proxyRequest(w, r, targetURL)
	} else if strings.HasPrefix(path, "/payment") {
		targetURL := g.paymentsServiceURL + path
		g.proxyRequest(w, r, targetURL)
	} else {
		log.Printf("[WARN] Unknown path: %s", path)
		http.NotFound(w, r)
	}
}

func main() {
	gateway := NewGateway()

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(gateway.routeHandler)

	log.Println("[INFO] API Gateway starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
