
package api

// ImportConfig represents a ImportConfig model.
type ImportConfig struct {
  Service *GenericImport `json:"service,omitempty"`
  Stream *GenericImport `json:"stream,omitempty"`
  Prefix string `json:"prefix,omitempty"`
  To string `json:"to,omitempty"`
}