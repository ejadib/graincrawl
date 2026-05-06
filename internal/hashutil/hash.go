package hashutil

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func String(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:])
}

func JSON(value any) string {
	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return String(string(b))
}
