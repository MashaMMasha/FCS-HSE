package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/upload", h.service.UploadFile).Methods("POST")
	apiRouter.HandleFunc("/files", h.service.ListFiles).Methods("GET")

	r.HandleFunc("/", h.IndexHandler).Methods("GET")
	r.HandleFunc("/upload", h.UploadPage).Methods("GET")
	r.HandleFunc("/files", h.FilesPage).Methods("GET")
	r.HandleFunc("/download/{filename}", h.DownloadFile).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	log.Println("Registered routes:")
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("%-6s %s", strings.Join(methods, ","), path)
		return nil
	})
}

func (h *FileHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func (h *FileHandler) UploadPage(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.ExecuteTemplate(w, "base.html", map[string]interface{}{
		"Content": "upload.html",
	})
	if err != nil {
		log.Printf("Error rendering upload page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *FileHandler) FilesPage(w http.ResponseWriter, r *http.Request) {
	files, err := h.service.ListFilesData()
	if err != nil {
		log.Printf("Error getting files: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Files data: %+v", files)

	data := struct {
		Title string
		Files []service.FileMeta
	}{
		Title: "Files List",
		Files: files,
	}

	err = h.tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Template Error", http.StatusInternalServerError)
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
