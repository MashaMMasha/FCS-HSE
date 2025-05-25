package service

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"file-storage/db"
)

type FileService struct {
	db *sql.DB
}

func NewFileService(db *sql.DB) *FileService {
	return &FileService{db: db}
}

type FileResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	File    FileMeta `json:"file,omitempty"`
}

type FileMeta struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

type FileListResponse struct {
	Files []FileMeta `json:"files"`
}

func (s *FileService) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		respondWithError(w, http.StatusBadRequest, "Не удалось разобрать форму")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Не удалось получить файл")
		return
	}
	defer file.Close()

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при создании директории")
		return
	}

	dstPath := filepath.Join("./uploads", header.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при сохранении файла")
		return
	}
	defer dst.Close()

	size, err := io.Copy(dst, file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при копировании файла")
		return
	}
	exists, err := db.FileExists(s.db, header.Filename)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при обращении к базе данных")
		return
	}
	if exists {
		respondWithError(w, http.StatusBadRequest, "Вы уже загружали такой файл")
		return
	}
	fileID, err := db.SaveFileMeta(s.db, header.Filename, size, dstPath)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при сохранении метаинформации")
		return
	}

	respondWithJSON(w, http.StatusCreated, FileResponse{
		Success: true,
		Message: "Файл успешно загружен",
		File: FileMeta{
			ID:   fileID,
			Name: header.Filename,
			Size: size,
			Path: dstPath,
		},
	})
}

func (s *FileService) ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := db.GetAllFiles(s.db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Ошибка при получении списка файлов")
		return
	}

	var fileList []FileMeta
	for _, file := range files {
		fileList = append(fileList, FileMeta{
			ID:   file["id"].(int64),
			Name: file["name"].(string),
			Size: file["size"].(int64),
			Path: file["path"].(string),
		})
	}

	respondWithJSON(w, http.StatusOK, FileListResponse{Files: fileList})
}

func (s *FileService) ListFilesData() ([]FileMeta, error) {
	files, err := db.GetAllFiles(s.db)
	if err != nil {
		return nil, err
	}

	var result []FileMeta
	for _, file := range files {
		result = append(result, FileMeta{
			ID:   file["id"].(int64),
			Name: file["name"].(string),
			Size: file["size"].(int64),
			Path: file["path"].(string),
		})
	}
	return result, nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, FileResponse{
		Success: false,
		Message: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
