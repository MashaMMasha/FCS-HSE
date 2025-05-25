CREATE TABLE IF NOT EXISTS analysis_results (
    id VARCHAR(36) PRIMARY KEY,
    filename TEXT NOT NULL,
    word_cloud TEXT,
    words_count INT NOT NULL DEFAULT 0,
    paragraphs INT NOT NULL DEFAULT 0,
    avg_word_length DECIMAL(5,2) NOT NULL DEFAULT 0,
    top_words JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );

CREATE INDEX IF NOT EXISTS idx_analysis_results_filename ON analysis_results(filename);
CREATE INDEX IF NOT EXISTS idx_analysis_results_created_at ON analysis_results(created_at);