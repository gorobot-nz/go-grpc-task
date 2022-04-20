package responses

type TimerResponse struct {
	SecondsRemaining int64 `json:"seconds-remaining"`
}

type BitlyResponse struct {
	Link string `json:"link"`
}
