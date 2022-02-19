package models

// UserConfig represents a UserConfig model.
type UserConfig struct {
	Username    string       `json:"username,omitempty"`
	Password    string       `json:"password,omitempty"`
	Nkey        string       `json:"nkey,omitempty"`
	Permissions *Permissions `json:"permissions,omitempty"`
}
