package storage

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/umputun/tg-spam/lib/spamcheck"
)

func TestDetectedSpam_NewDetectedSpam(t *testing.T) {
	db, err := sqlx.Open("sqlite", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	_, err = NewDetectedSpam(db)
	require.NoError(t, err)

	var exists int
	err = db.Get(&exists, "SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='detected_spam'")
	require.NoError(t, err)
	assert.Equal(t, 1, exists)
}

func TestDetectedSpam_Write(t *testing.T) {
	db, err := sqlx.Open("sqlite", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	ds, err := NewDetectedSpam(db)
	require.NoError(t, err)

	spamEntry := DetectedSpamInfo{
		Text:      "spam message",
		UserID:    1,
		UserName:  "Spammer",
		Timestamp: time.Now(),
	}

	checks := []spamcheck.Response{
		{
			Name:    "Check1",
			Spam:    true,
			Details: "Details 1",
		},
	}

	err = ds.Write(spamEntry, checks)
	require.NoError(t, err)

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM detected_spam")
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestDetectedSpam_Read(t *testing.T) {
	db, err := sqlx.Open("sqlite", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	ds, err := NewDetectedSpam(db)
	require.NoError(t, err)

	spamEntry := DetectedSpamInfo{
		Text:      "spam message",
		UserID:    1,
		UserName:  "Spammer",
		Timestamp: time.Now(),
	}

	checks := []spamcheck.Response{
		{
			Name:    "Check1",
			Spam:    true,
			Details: "Details 1",
		},
	}

	checksJSON, err := json.Marshal(checks)
	require.NoError(t, err)
	_, err = db.Exec("INSERT INTO detected_spam (text, user_id, user_name, timestamp, checks) VALUES (?, ?, ?, ?, ?)", spamEntry.Text, spamEntry.UserID, spamEntry.UserName, spamEntry.Timestamp, checksJSON)
	require.NoError(t, err)

	entries, err := ds.Read()
	require.NoError(t, err)
	require.Len(t, entries, 1)

	assert.Equal(t, spamEntry.Text, entries[0].Text)
	assert.Equal(t, spamEntry.UserID, entries[0].UserID)
	assert.Equal(t, spamEntry.UserName, entries[0].UserName)

	var retrievedChecks []spamcheck.Response
	err = json.Unmarshal([]byte(entries[0].ChecksJSON), &retrievedChecks)
	require.NoError(t, err)
	assert.Equal(t, checks, retrievedChecks)
	t.Logf("retrieved checks: %+v", retrievedChecks)
}