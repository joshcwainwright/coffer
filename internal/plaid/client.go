package plaid

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/joshcwainwright/coffer/internal/config"
)

type Client struct {
	baseURL  string
	clientID string
	secret   config.Secret
	http     *http.Client
}

const (
	sandboxURL    = "https://sandbox.plaid.com"
	productionURL = "https://production.plaid.com"
)

func New(env config.PlaidEnv, clientID string, secret config.Secret) *Client {
	return &Client{
		baseURL:  baseURL(env),
		clientID: clientID,
		secret:   secret,
		http:     &http.Client{Timeout: 30 * time.Second},
	}
}

func baseURL(env config.PlaidEnv) string {
	switch env {
	case config.PlaidProduction:
		return productionURL

	default:
		return sandboxURL
	}
}

func (c *Client) do(ctx context.Context, path string, req, out any) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("new request %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("PLAID-CLIENT-ID", c.clientID)
	httpReq.Header.Set("PLAID-SECRET", c.secret.Value())

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return fmt.Errorf("post %s: %w", path, err)
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr Error
		if err := json.Unmarshal(data, &apiErr); err != nil || apiErr.Code == "" {
			return fmt.Errorf("plaid returned %s", resp.Status)
		}
		return &apiErr
	}

	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}
