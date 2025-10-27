package dto

type ExternalAuthProfile struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Source  string `json:"source"`
}
