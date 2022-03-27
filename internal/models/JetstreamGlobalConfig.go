package models

// JetstreamConfig represents a JetstreamConfig model.
type JetstreamGlobalConfig struct {
	StoreDir     string `json:"store_dir,omitempty"`
	MaxMem       *int64 `json:"max_mem,omitempty"`
	MaxFile      *int64 `json:"max_file,omitempty"`
	MaxStreams   *int64 `json:"max_streams,omitempty"`
	MaxConsumers *int64 `json:"max_consumers,omitempty"`
}
