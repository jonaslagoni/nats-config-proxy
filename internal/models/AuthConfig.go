package models

// AuthConfig represents the complete authorization config
// for the NATS Server.
type AuthConfig struct {
	// Users that belong to the global account.
	Users []*UserConfig `json:"users"`

	// Accounts that separate the subject namespaces.
	Accounts map[string]*Account `json:"accounts,omitempty"`
}
