package sqlite

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"

	"writer-api/internal/model"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS failed_events (
		id TEXT PRIMARY KEY,
		event_json TEXT NOT NULL,
		failed_kafka BOOLEAN NOT NULL,
		failed_mongo BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

func (s *SQLiteStore) SaveFailedEvent(event model.StripeEvent, failedKafka, failedMongo bool) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	query := `
	INSERT OR REPLACE INTO failed_events (id, event_json, failed_kafka, failed_mongo)
	VALUES (?, ?, ?, ?);
	`
	_, err = s.db.Exec(query, event.ID, string(jsonData), failedKafka, failedMongo)
	return err
}

func (s *SQLiteStore) GetPendingEvents() ([]model.StripeEvent, error) {
	rows, err := s.db.Query("SELECT event_json FROM failed_events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.StripeEvent
	for rows.Next() {
		var jsonStr string
		if err := rows.Scan(&jsonStr); err != nil {
			return nil, err
		}

		var event model.StripeEvent
		if err := json.Unmarshal([]byte(jsonStr), &event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (s *SQLiteStore) DeleteEvent(id string) error {
	_, err := s.db.Exec("DELETE FROM failed_events WHERE id = ?", id)
	return err
}
