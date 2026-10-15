package privateapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBatchCompatibilityChunks(t *testing.T) {
	for _, n := range []int{0, 1, 50, 51, 98, 101} {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			calls := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var b BatchRequest
				json.NewDecoder(r.Body).Decode(&b)
				calls++
				if len(b.DocumentIDs) > 50 {
					t.Error("oversized")
					w.WriteHeader(400)
					return
				}
				docs := []Document{}
				for _, id := range b.DocumentIDs {
					docs = append(docs, Document{ID: id})
				}
				json.NewEncoder(w).Encode(BatchResponse{Docs: docs})
			}))
			defer srv.Close()
			ids := []string{}
			for i := 0; i < n; i++ {
				ids = append(ids, fmt.Sprint(i))
			}
			c := Client{BaseURL: srv.URL}
			got, err := c.GetDocumentsBatch(context.Background(), ids)
			if err != nil || len(got.Docs) != n || calls != (n+49)/50 {
				t.Fatalf("calls=%d len=%d err=%v", calls, len(got.Docs), err)
			}
			for i, d := range got.Docs {
				if d.ID != ids[i] {
					t.Fatal("order")
				}
			}
		})
	}
}
