package config

type Redacted[T any] struct{ v T }

const (
	RedactedJSON   string = `"[REDACTED]"`
	RedactedString string = "[REDACTED]"
)

func NewRedacted[T any](v T) Redacted[T]         { return Redacted[T]{v: v} }
func (r Redacted[T]) Value() T                   { return r.v }
func (Redacted[T]) String() string               { return RedactedString }
func (Redacted[T]) GoString() string             { return RedactedString }
func (Redacted[T]) MarshalJSON() ([]byte, error) { return []byte(RedactedJSON), nil }
