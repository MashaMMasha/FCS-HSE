package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AnalysisService struct {
	db *sql.DB
}

type AnalysisResult struct {
	ID         string         `json:"id"`
	Filename   string         `json:"filename"`
	WordCloud  string         `json:"wordCloud"`
	WordsCount int            `json:"wordsCount"`
	Paragraphs int            `json:"paragraphs"`
	AvgWordLen float64        `json:"avgWordLength"`
	TopWords   map[string]int `json:"topWords"`
	CreatedAt  time.Time      `json:"createdAt"`
}

func NewAnalysisService(db *sql.DB) *AnalysisService {
	return &AnalysisService{db: db}
}

func (s *AnalysisService) fetchFileContent(filename string) (string, error) {
	url := fmt.Sprintf("http://api-gateway:8080/api/storage/files/%s", filename)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request to file-storage failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("file-storage returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read file-storage response: %w", err)
	}

	return string(body), nil
}

func (s *AnalysisService) Analyze(filename string) (*AnalysisResult, error) {
	content, err := s.fetchFileContent(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %w", err)
	}

	text := string(content)

	if strings.TrimSpace(text) == "" {
		return nil, fmt.Errorf("file content is empty")
	}

	wordCloudURL, err := generateWordCloud(text)
	if err != nil {
		return nil, fmt.Errorf("word cloud generation failed: %w", err)
	}

	stats := calculateTextStats(text)

	result := &AnalysisResult{
		ID:         uuid.NewString(),
		Filename:   filename,
		WordCloud:  wordCloudURL,
		WordsCount: stats.WordsCount,
		Paragraphs: stats.Paragraphs,
		AvgWordLen: stats.AvgWordLen,
		TopWords:   stats.TopWords,
		CreatedAt:  time.Now(),
	}

	if err := s.saveResult(result); err != nil {
		return nil, fmt.Errorf("failed to save result: %w", err)
	}

	return result, nil
}

func generateWordCloud(text string) (string, error) {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return "", fmt.Errorf("text is empty")
	}

	if len(text) > 8000 {
		text = text[:8000]
	}

	baseURL := "https://quickchart.io/wordcloud"
	params := url.Values{}

	params.Set("text", text)
	params.Set("format", "png")
	params.Set("width", "800")
	params.Set("height", "600")
	params.Set("backgroundColor", "white")
	params.Set("fontFamily", "Arial")
	params.Set("fontScale", "25")
	params.Set("padding", "2")
	params.Set("rotation", "20")
	params.Set("maxNumWords", "100")
	params.Set("removeStopwords", "true")
	params.Set("cleanWords", "true")
	params.Set("colors", `["#1f77b4", "#ff7f0e", "#2ca02c", "#d62728", "#9467bd", "#8c564b"]`)

	fullURL := baseURL + "?" + params.Encode()

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "TextAnalysis/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to request wordcloud service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("WordCloud API Error - Status: %d, URL: %s, Body: %s", resp.StatusCode, fullURL, string(body))
		return "", fmt.Errorf("wordcloud service returned status %d: %s", resp.StatusCode, string(body))
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "image/") {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Unexpected content type from WordCloud API: %s, Body: %s", contentType, string(body))
		return "", fmt.Errorf("wordcloud service returned unexpected content type: %s", contentType)
	}

	return fullURL, nil
}

type TextStats struct {
	WordsCount int
	Paragraphs int
	AvgWordLen float64
	TopWords   map[string]int
}

func calculateTextStats(text string) TextStats {
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	totalLength := 0

	for _, word := range words {
		cleaned := strings.ToLower(strings.Trim(word, ".,!?\"':;()[]{}<>"))
		if cleaned != "" {
			wordCount[cleaned]++
			totalLength += len(cleaned)
		}
	}

	paragraphs := 0
	lines := strings.Split(text, "\n")
	inParagraph := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			if !inParagraph {
				paragraphs++
				inParagraph = true
			}
		} else {
			inParagraph = false
		}
	}

	var avgWordLen float64
	if len(words) > 0 {
		avgWordLen = float64(totalLength) / float64(len(words))
	}

	return TextStats{
		WordsCount: len(words),
		Paragraphs: paragraphs,
		AvgWordLen: avgWordLen,
		TopWords:   findTopWords(wordCount, 5),
	}
}

func (s *AnalysisService) saveResult(result *AnalysisResult) error {
	topWordsJSON, _ := json.Marshal(result.TopWords)

	_, err := s.db.Exec(
		"INSERT INTO analysis_results (id, filename, word_cloud, words_count, paragraphs, avg_word_length, top_words, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		result.ID, result.Filename, result.WordCloud, result.WordsCount, result.Paragraphs, result.AvgWordLen, string(topWordsJSON), result.CreatedAt,
	)
	return err
}

func (s *AnalysisService) GetResult(id string) (*AnalysisResult, error) {
	var result AnalysisResult
	var topWordsJSON string

	err := s.db.QueryRow(
		"SELECT id, filename, word_cloud, words_count, paragraphs, avg_word_length, top_words, created_at FROM analysis_results WHERE id = $1",
		id,
	).Scan(&result.ID, &result.Filename, &result.WordCloud, &result.WordsCount, &result.Paragraphs, &result.AvgWordLen, &topWordsJSON, &result.CreatedAt)

	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(topWordsJSON), &result.TopWords)
	return &result, nil
}

func (s *AnalysisService) GetAllResults() ([]AnalysisResult, error) {
	rows, err := s.db.Query(
		"SELECT id, filename, word_cloud, words_count, paragraphs, avg_word_length, top_words, created_at FROM analysis_results ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []AnalysisResult
	for rows.Next() {
		var result AnalysisResult
		var topWordsJSON string

		if err := rows.Scan(&result.ID, &result.Filename, &result.WordCloud, &result.WordsCount, &result.Paragraphs, &result.AvgWordLen, &topWordsJSON, &result.CreatedAt); err != nil {
			return nil, err
		}

		json.Unmarshal([]byte(topWordsJSON), &result.TopWords)
		results = append(results, result)
	}

	return results, nil
}

func findTopWords(wordCount map[string]int, n int) map[string]int {
	stopWords := map[string]bool{
		"и": true, "в": true, "не": true, "на": true, "я": true, "с": true, "что": true, "а": true, "по": true, "как": true,
		"к": true, "у": true, "за": true, "но": true, "из": true, "это": true, "все": true, "от": true, "так": true, "о": true,
		"же": true, "бы": true, "для": true, "то": true, "мы": true,
		"the": true, "and": true, "a": true, "an": true, "in": true, "on": true, "at": true, "for": true, "to": true,
		"of": true, "with": true, "as": true, "by": true, "this": true, "that": true,
	}

	filtered := make(map[string]int)
	for word, count := range wordCount {
		if !stopWords[word] {
			filtered[word] = count
		}
	}

	type wordFreq struct {
		word  string
		count int
	}
	var wordList []wordFreq
	for word, count := range filtered {
		wordList = append(wordList, wordFreq{word, count})
	}

	sort.Slice(wordList, func(i, j int) bool {
		if wordList[i].count == wordList[j].count {
			return wordList[i].word < wordList[j].word
		}
		return wordList[i].count > wordList[j].count
	})

	topWords := make(map[string]int)
	for i := 0; i < n && i < len(wordList); i++ {
		topWords[wordList[i].word] = wordList[i].count
	}

	return topWords
}
