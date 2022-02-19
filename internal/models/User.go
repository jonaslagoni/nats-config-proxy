package models

// Identity represents a Identity model.
type User struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Nkey        string `json:"nkey,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Account     string `json:"account,omitempty"`
}
