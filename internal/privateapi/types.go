package privateapi

import "encoding/json"

type DocumentsRequest struct {
	IncludeSharedWithMe bool    `json:"include_shared_with_me,omitempty"`
	CreatedAfter        *string `json:"created_after,omitempty"`
	CreatedBefore       *string `json:"created_before,omitempty"`
}

type DocumentsResponse struct {
	Docs    []Document `json:"docs"`
	Deleted []string   `json:"deleted"`
	Shared  []Document `json:"shared"`
}

type BatchRequest struct {
	DocumentIDs []string `json:"document_ids"`
}

type BatchResponse struct {
	Docs []Document `json:"docs"`
}

type TranscriptRequest struct {
	DocumentID string `json:"document_id"`
}

type PanelsRequest struct {
	DocumentID string `json:"document_id"`
}

type Document struct {
	ID                  string          `json:"id"`
	CreatedAt           string          `json:"created_at"`
	Title               *string         `json:"title"`
	UserID              string          `json:"user_id"`
	Notes               json.RawMessage `json:"notes"`
	NotesPlain          *string         `json:"notes_plain"`
	NotesMarkdown       *string         `json:"notes_markdown"`
	GoogleCalendarEvent json.RawMessage `json:"google_calendar_event"`
	UpdatedAt           string          `json:"updated_at"`
	DeletedAt           *string         `json:"deleted_at"`
	Type                string          `json:"type"`
	Status              *string         `json:"status"`
	People              json.RawMessage `json:"people"`
	Summary             json.RawMessage `json:"summary"`
	WorkspaceID         *string         `json:"workspace_id"`
}

type TranscriptChunk struct {
	DocumentID        string  `json:"document_id"`
	StartTimestamp    string  `json:"start_timestamp"`
	Text              string  `json:"text"`
	Source            string  `json:"source"`
	ID                string  `json:"id"`
	IsFinal           bool    `json:"is_final"`
	EndTimestamp      string  `json:"end_timestamp"`
	TranscriberUserID *string `json:"transcriber_user_id"`
}

type Panel struct {
	ID               string          `json:"id"`
	DocumentID       string          `json:"document_id"`
	CreatedAt        string          `json:"created_at"`
	Title            *string         `json:"title"`
	Content          json.RawMessage `json:"content"`
	DeletedAt        *string         `json:"deleted_at"`
	TemplateSlug     *string         `json:"template_slug"`
	LastViewedAt     *string         `json:"last_viewed_at"`
	UpdatedAt        *string         `json:"updated_at"`
	ContentUpdatedAt *string         `json:"content_updated_at"`
	YdocVersion      *int64          `json:"ydoc_version"`
}
