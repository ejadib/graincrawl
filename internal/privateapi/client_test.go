package privateapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientSendsGranolaHeaders(t *testing.T) {
	var gotAuth, gotWorkspace string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotWorkspace = r.Header.Get("X-Granola-Workspace-Id")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"docs":[],"deleted":[]}`))
	}))
	defer srv.Close()

	client := Client{BaseURL: srv.URL, AccessToken: "token", WorkspaceID: "workspace"}
	if _, err := client.GetDocuments(context.Background(), DocumentsRequest{}); err != nil {
		t.Fatal(err)
	}
	if gotAuth != "Bearer token" || gotWorkspace != "workspace" {
		t.Fatalf("headers auth=%q workspace=%q", gotAuth, gotWorkspace)
	}
}

func TestNoteFromDocumentExtractsRawNotesAndSummary(t *testing.T) {
	title := "Planning"
	doc := Document{
		ID:        "doc-1",
		Title:     &title,
		Type:      "meeting",
		CreatedAt: "2026-05-06T10:00:00Z",
		UpdatedAt: "2026-05-06T10:01:00Z",
		Notes:     json.RawMessage(`{"markdown":"## Notes\nship the archive"}`),
		Summary:   json.RawMessage(`{"text":"summary text"}`),
	}
	note, err := NoteFromDocument(doc, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	if note.NotesMarkdown == nil || *note.NotesMarkdown != "## Notes\nship the archive" {
		t.Fatalf("notes markdown not extracted: %#v", note.NotesMarkdown)
	}
	if note.SummaryText == nil || *note.SummaryText != "summary text" {
		t.Fatalf("summary text not extracted: %#v", note.SummaryText)
	}
}
