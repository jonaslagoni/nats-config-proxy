package models

// ResponseRules represents the reponse permissions.
type ResponseRules struct {
	Max     int    `json:"max,omitempty"`
	Expires string `json:"expires,omitempty"`
}
