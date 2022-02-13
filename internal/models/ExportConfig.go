
package api

// ExportConfig represents a ExportConfig model.
type ExportConfig struct {
  Stream string `json:"stream,omitempty"`
  Service string `json:"service,omitempty"`
  Accounts []string `json:"Accounts,omitempty"`
  Response string `json:"response,omitempty"`
}