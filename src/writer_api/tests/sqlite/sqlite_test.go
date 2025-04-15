package sqlite_test

import (
	"testing"
	"writer-api/internal/sqlite"
	"writer-api/tests/model"
)

var event = mock_event.BuildMockStripeEvent()

func TestSendFailedEvents(t *testing.T) {
	err := sqlite.NewSQLiteStore("./db.sqlite")
	if err != nil {
		t.Fatalf("error with the sqlite store definition %s", err)
	}

	err = sqlite.SaveFailedEvent(event, true, true)
	if err != nil {
		t.Fatalf("error saving the event %s", err)
	} else {
		t.Logf("success saving the event")
	}
}

func TestGetPendingEvents(t *testing.T) {
	err := sqlite.NewSQLiteStore("./db.sqlite")
	if err != nil {
		t.Fatalf("error with the sqlite store definition %s", err)
	}

	_, err = sqlite.GetPendingEvents()
	if err != nil {
		t.Fatalf("error retrieving the event %s", err)
	} else {
		t.Logf("success retrieving the event")
	}
}

func TestDeleteEvents(t *testing.T) {
	err := sqlite.NewSQLiteStore("./db.sqlite")
	if err != nil {
		t.Fatalf("error with the sqlite store definition %s", err)
	}

	err = sqlite.DeleteEvent(event.ID)
	if err != nil {
		t.Fatalf("error deleting the event %s", err)
	} else {
		t.Logf("success deleting the event")
	}
}
