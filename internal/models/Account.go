package models

// Account represents a Account model.
type Account struct {
	Users     []*UserConfig    `json:"users,omitempty"`
	Exports   []*ExportConfig  `json:"exports,omitempty"`
	Imports   []*ImportConfig  `json:"imports,omitempty"`
	Jetstream *JetstreamConfig `json:"jetstream,omitempty"`
}
