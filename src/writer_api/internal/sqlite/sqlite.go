package sqlite

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"writer-api/internal/model"
)

type SQLiteStore struct {
	db *sql.DB
}

var store *SQLiteStore

func NewSQLiteStore(path string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
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
	store = &SQLiteStore{db: db}
	log.Println("✅ SQLite db connected")
	if _, err := db.Exec(query); err != nil {
		log.Println("⚠️ SQLite db was initialized already")
		return nil
	}

	log.Println("✅ SQLite db initialized")
	return nil
}

func CloseConnection() error {
	err := store.db.Close()

	if err != nil {
		return err
	}

	return nil
}

func SaveFailedEvent(event model.StripeEvent, failedKafka, failedMongo bool) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// TODO: check sql injection
	query := `
	INSERT OR REPLACE INTO failed_events (id, event_json, failed_kafka, failed_mongo)
	VALUES (?, ?, ?, ?);
	`
	_, err = store.db.Exec(query, event.ID, string(jsonData), failedKafka, failedMongo)
	return err
}

func GetPendingEvents() ([]model.StripeEvent, error) {
	rows, err := store.db.Query("SELECT event_json FROM failed_events")
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

func DeleteEvent(id string) error {
	_, err := store.db.Exec("DELETE FROM failed_events WHERE id = ?", id)
	return err
}
