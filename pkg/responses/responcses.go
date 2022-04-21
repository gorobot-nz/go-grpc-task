package responses

import "encoding/json"

type TimerResponse struct {
	SecondsRemaining json.Number `json:"seconds-remaining"`
}

type BitlyResponse struct {
	Link string `json:"link"`
}
