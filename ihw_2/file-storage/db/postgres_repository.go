package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func SaveFileMeta(db *sql.DB, name string, size int64, path string) (int64, error) {
	var id int64
	err := db.QueryRow(
		"INSERT INTO files (name, size, path) VALUES ($1, $2, $3) RETURNING id",
		name, size, path,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save file meta: %w", err)
	}
	return id, nil
}

func GetAllFiles(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT id, name, size, path FROM files")
	if err != nil {
		return nil, fmt.Errorf("failed to query files: %w", err)
	}
	defer rows.Close()

	var files []map[string]interface{}
	for rows.Next() {
		var id int64
		var name string
		var size int64
		var path string

		if err := rows.Scan(&id, &name, &size, &path); err != nil {
			return nil, fmt.Errorf("failed to scan file row: %w", err)
		}

		files = append(files, map[string]interface{}{
			"id":   id,
			"name": name,
			"size": size,
			"path": path,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return files, nil
}
