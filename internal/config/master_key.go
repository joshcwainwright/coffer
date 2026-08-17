package config

import (
	"encoding/base64"
	"fmt"
)

type MasterKey = Redacted[[]byte]

func parseMasterKey(s string) (MasterKey, error) {
	k, err := base64.StdEncoding.DecodeString(s)
	switch {
	case err != nil:
		return MasterKey{}, fmt.Errorf("must be valid base64: %w", err)
	case len(k) != 32:
		return MasterKey{}, fmt.Errorf("must decode to 32 bytes, got %d", len(k))
	default:
		return NewRedacted(k), nil
	}
}
