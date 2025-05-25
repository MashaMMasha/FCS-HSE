package handler

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"file-storage/service"
	"github.com/gorilla/mux"
)

type FileHandler struct {
	service *service.FileService
	tmpl    *template.Template
}

func NewFileHandler(db *sql.DB) (*FileHandler, error) {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	log.Println("Loaded templates:", tmpl.DefinedTemplates())

	return &FileHandler{
		service: service.NewFileService(db),
		tmpl:    tmpl,
	}, nil
}

func (h *FileHandler) RegisterRoutes(r *mux.Router) {
	apiRouter := r.PathPrefix("/api/storage").Subrouter()
	apiRouter.HandleFunc("/upload", h.service.UploadFile).Methods("POST")
	apiRouter.HandleFunc("/files", h.service.ListFiles).Methods("GET")
	apiRouter.HandleFunc("/files/{filename}", h.GetFileContent).Methods("GET")

	r.HandleFunc("/upload", h.UploadPage).Methods("GET")
	r.HandleFunc("/files", h.FilesPage).Methods("GET")
	r.HandleFunc("/download/{filename}", h.DownloadFile).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
}

func (h *FileHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func (h *FileHandler) FilesPage(w http.ResponseWriter, r *http.Request) {
	files, err := h.service.ListFilesData()
	if err != nil {
		log.Printf("Error getting files: %v", err)
		http.Error(w, "Failed to load files", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err = h.tmpl.ExecuteTemplate(w, "files.html", files)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Template rendering failed", http.StatusInternalServerError)
	}
}

func (h *FileHandler) UploadPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := h.tmpl.ExecuteTemplate(w, "upload.html", nil)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Template rendering failed", http.StatusInternalServerError)
	}
}

func (h *FileHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	filePath := "./uploads/" + filename
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	http.ServeFile(w, r, filePath)
}

func (h *FileHandler) GetFileContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	log.Printf("[GET FILE] Запрос на получение файла: %s", filename)

	if filename == "" {
		log.Printf("[ERROR] Не указано имя файла")
		http.Error(w, "Filename is required", http.StatusBadRequest)
		return
	}

	if filepath.IsAbs(filename) || filepath.Clean(filename) != filename {
		log.Printf("[ERROR] Недопустимое имя файла: %s", filename)
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("./uploads", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("[ERROR] Файл не найден: %s", filename)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("[ERROR] Ошибка при открытии файла %s: %v", filename, err)
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("[ERROR] Ошибка при чтении файла %s: %v", filename, err)
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	log.Printf("[SUCCESS] Отдача содержимого файла: %s (%d байт)", filename, len(content))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Filename", filename)
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
