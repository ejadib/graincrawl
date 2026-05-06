package cachev6

import "encoding/json"

type File struct {
	Cache Cache `json:"cache"`
}

type Cache struct {
	Version int   `json:"version"`
	State   State `json:"state"`
}

type State struct {
	Documents        map[string]Document          `json:"documents"`
	Transcripts      map[string][]TranscriptChunk `json:"transcripts"`
	MeetingsMetadata map[string]json.RawMessage   `json:"meetingsMetadata"`
	FeatureFlags     map[string]any               `json:"featureFlags"`
}

type Document struct {
	ID                  string          `json:"id"`
	CreatedAt           string          `json:"created_at"`
	Title               *string         `json:"title"`
	UserID              string          `json:"user_id"`
	NotesPlain          *string         `json:"notes_plain"`
	NotesMarkdown       *string         `json:"notes_markdown"`
	UpdatedAt           string          `json:"updated_at"`
	DeletedAt           *string         `json:"deleted_at"`
	Type                string          `json:"type"`
	Status              *string         `json:"status"`
	WorkspaceID         *string         `json:"workspace_id"`
	GoogleCalendarEvent json.RawMessage `json:"google_calendar_event"`
	People              json.RawMessage `json:"people"`
	Raw                 json.RawMessage `json:"-"`
}

type TranscriptChunk struct {
	DocumentID        string          `json:"document_id"`
	StartTimestamp    string          `json:"start_timestamp"`
	Text              string          `json:"text"`
	Source            string          `json:"source"`
	ID                string          `json:"id"`
	IsFinal           bool            `json:"is_final"`
	EndTimestamp      string          `json:"end_timestamp"`
	TranscriberUserID *string         `json:"transcriber_user_id"`
	Raw               json.RawMessage `json:"-"`
}
