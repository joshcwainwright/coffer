package config

type Config struct {
	DBPath        string
	PlaidClientID string
	PlaidSecret   Secret
	PlaidEnv      PlaidEnv
	MasterKey     MasterKey
}

func Load() (Config, error) {
	var l loader

	config := Config{
		DBPath:        parsedOr(&l, "COFFER_DB_PATH", parseString, "./data/coffer.db"),
		PlaidClientID: parsed(&l, "PLAID_CLIENT_ID", parseString),
		PlaidSecret:   parsed(&l, "PLAID_SECRET", parseSecret),
		PlaidEnv:      parsed(&l, "PLAID_ENV", parseEnv),
		MasterKey:     parsed(&l, "COFFER_MASTER_KEY", parseMasterKey),
	}

	if fail := l.fail(); fail != nil {
		return Config{}, fail
	}

	return config, nil
}
