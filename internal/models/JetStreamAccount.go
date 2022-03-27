package models

// JetstreamAccount represents a JetstreamAccount model.
type JetstreamAccount struct {
	MaxMem       *int64 `json:"max_mem,omitempty"`
	MaxFile      *int64 `json:"max_file,omitempty"`
	MaxStreams   *int64 `json:"max_streams,omitempty"`
	MaxConsumers *int64 `json:"max_consumers,omitempty"`
}
