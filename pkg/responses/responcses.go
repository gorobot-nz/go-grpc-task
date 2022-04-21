package responses

import "encoding/json"

type TimerResponse struct {
	SecondsRemaining json.Number `json:"seconds_remaining"`
}

type BitlyResponse struct {
	Link string `json:"link"`
}
