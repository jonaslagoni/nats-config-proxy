
package api

// Account represents a Account model.
type Account struct {
  Users []*UserConfig `json:"users,omitempty"`
  Exports []*ExportConfig `json:"exports,omitempty"`
  Imports []*ImportConfig `json:"imports,omitempty"`
  Jetstream []*JetStreamConfig `json:"jetstream,omitempty"`
}