package config

type Secret = Redacted[string]

func parseSecret(s string) (Secret, error) {
	return NewRedacted(s), nil
}
