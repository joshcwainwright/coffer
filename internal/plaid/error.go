package plaid

import "fmt"

type Error struct {
	Type      string `json:"error_type"`
	Code      string `json:"error_code"`
	Message   string `json:"error_message"`
	RequestID string `json:"request_id"`
}

func (e *Error) Error() string { return fmt.Sprintf("plaid %s: %s", e.Code, e.Message) }
