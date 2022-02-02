package rest

import "encoding/json"

// Response defines data for response
type Response struct {
	Data json.RawMessage `json:"data"`
}