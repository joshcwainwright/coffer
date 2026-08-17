package config

import (
	"errors"
	"fmt"
	"os"
)

type loader struct{ errs []error }

func (l *loader) err(e error) {
	l.errs = append(l.errs, e)
}

func (l *loader) fail() error { return errors.Join(l.errs...) }

func parseString(s string) (string, error) { return s, nil }

func parsed[T any](l *loader, key string, parse func(string) (T, error)) T {
	var zero T
	v := os.Getenv(key)

	if v == "" {
		l.err(fmt.Errorf("%s is required", key))
		return zero
	}

	out, err := parse(v)
	if err != nil {
		l.err(fmt.Errorf("%s %w", key, err))
		return zero
	}

	return out
}

func parsedOr[T any](l *loader, key string, parse func(string) (T, error), fallback T) T {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}

	out, err := parse(v)
	if err != nil {
		l.err(fmt.Errorf("%s %w", key, err))
		return fallback
	}

	return out
}
