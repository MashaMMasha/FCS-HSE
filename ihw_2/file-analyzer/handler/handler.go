package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"file-analyzer/service"
	"github.com/gorilla/mux"
)

type AnalysisHandler struct {
	service *service.AnalysisService
}

func NewAnalysisHandler(db *sql.DB) (*AnalysisHandler, error) {
	svc := service.NewAnalysisService(db)
	return &AnalysisHandler{service: svc}, nil
}

func (h *AnalysisHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/analysis/results/{filename:.*}", h.GetResult).Methods("GET")
	r.HandleFunc("/api/analysis/results", h.GetAllResults).Methods("GET")
}

func (h *AnalysisHandler) GetResult(w http.ResponseWriter, r *http.Request) {
	filename := mux.Vars(r)["filename"]
	log.Printf(">>> Получен запрос на анализ файла: %s", filename)

	result, err := h.service.GetResult(filename)
	if err != nil {
		log.Printf("Результат анализа не найден для %s, запускаем анализ...", filename)

		result, err = h.service.Analyze(filename)
		if err != nil {
			log.Printf("Ошибка при анализе файла %s: %v", filename, err)
			http.Error(w, "Не удалось проанализировать файл", http.StatusInternalServerError)
			return
		}
	} else {
		log.Printf("Результат анализа для %s найден", filename)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Ошибка при кодировании результата: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
	}
}

func (h *AnalysisHandler) GetAllResults(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> Получен запрос на список всех результатов анализа")
	results, err := h.service.GetAllResults()
	if err != nil {
		log.Printf("Ошибка при получении списка результатов: %v", err)
		http.Error(w, "Failed to fetch results", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
