package redact

import "testing"

func TestKeyDetectsSensitiveNames(t *testing.T) {
	for _, key := range []string{"access_token", "refresh_token", "notes_markdown", "transcript_text", "Authorization"} {
		if !Key(key) {
			t.Fatalf("expected %q to be sensitive", key)
		}
	}
	if Key("created_at") {
		t.Fatal("created_at should not be sensitive")
	}
}
