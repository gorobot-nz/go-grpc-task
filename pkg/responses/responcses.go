package responses

type TimerResponse struct {
	SecondsRemaining int `json:"seconds-remaining"`
}

type BitlyResponse struct {
	Link string `json:"link"`
}
