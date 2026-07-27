package config

import "fmt"

type PlaidEnv string

const (
	PlaidSandbox    PlaidEnv = "sandbox"
	PlaidProduction PlaidEnv = "production"
)

func parseEnv(s string) (PlaidEnv, error) {
	switch e := PlaidEnv(s); e {
	case PlaidSandbox, PlaidProduction:
		return e, nil
	default:
		return "", fmt.Errorf("must be %q or %q, got %q", PlaidSandbox, PlaidProduction, e)
	}
}
