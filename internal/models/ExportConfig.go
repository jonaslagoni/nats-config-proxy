package models

// ExportConfig represents a ExportConfig model.
type ExportConfig struct {
	Stream   string   `json:"stream,omitempty"`
	Service  string   `json:"service,omitempty"`
	Accounts []string `json:"accounts,omitempty"`
	Response string   `json:"response,omitempty"`
}
