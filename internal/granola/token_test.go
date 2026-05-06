package granola

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestReadSupabaseParsesNestedTokens(t *testing.T) {
	tokens := WorkOSTokens{AccessToken: "secret", RefreshToken: "refresh", ObtainedAt: 1000, ExpiresIn: 3600, SignInMethod: "CrossAppAuth"}
	user := UserInfo{ID: "u1", WorkspaceIDs: []string{"w1"}}
	tb, _ := json.Marshal(tokens)
	ub, _ := json.Marshal(user)
	raw, _ := json.Marshal(SupabaseFile{SessionID: "s1", WorkOSTokens: string(tb), UserInfoRaw: string(ub)})
	path := filepath.Join(t.TempDir(), "supabase.json")
	if err := os.WriteFile(path, raw, 0o600); err != nil {
		t.Fatal(err)
	}
	_, gotTokens, gotUser, err := ReadSupabase(path)
	if err != nil {
		t.Fatal(err)
	}
	if gotTokens.AccessToken != "secret" || gotUser.WorkspaceIDs[0] != "w1" {
		t.Fatalf("unexpected parsed values: %#v %#v", gotTokens, gotUser)
	}
	summary := SummarizeToken(gotTokens, time.UnixMilli(1000))
	if !summary.Present || summary.Expired || !summary.RefreshPresent {
		t.Fatalf("unexpected summary: %#v", summary)
	}
}
