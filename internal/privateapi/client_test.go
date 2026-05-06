package privateapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
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
