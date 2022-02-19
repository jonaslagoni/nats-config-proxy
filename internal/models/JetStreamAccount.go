package models

// JetstreamAccount represents a JetstreamAccount model.
type JetstreamAccount struct {
	MaxMem       int `json:"max mem,omitempty"`
	MaxFile      int `json:"max_file,omitempty"`
	MaxStreams   int `json:"max_streams,omitempty"`
	MaxConsumers int `json:"max_consumers,omitempty"`
}
