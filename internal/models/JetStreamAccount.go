package api

import (
	"encoding/json"
	"fmt"
)

// JetStreamAccount represents a JetStreamAccount model.
type JetStreamAccount struct {
	MaxMem       int `json:"max_mem,omitempty"`
	MaxFile      int `json:"max_file,omitempty"`
	MaxStreams   int `json:"max_streams,omitempty"`
	MaxConsumers int `json:"max_consumers,omitempty"`
}

func (m *JetStreamAccount) MarshalJSON() ([]byte, error) {
	jsonObj := make(map[string]interface{})
	jsonObj["max_mem"] = m.MaxMem
	jsonObj["max_file"] = m.MaxFile
	jsonObj["max_streams"] = m.MaxStreams
	jsonObj["max_consumers"] = m.MaxConsumers
	return json.Marshal(jsonObj)
}

func UnmarshalJetStreamAccount(body []byte, c *JetStreamAccount) error {
	jsonObj := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonObj)
	if err != nil {
		return err
	}

	for key, element := range jsonObj {
		fmt.Println("Converting Key:", key, " with ", "Element:", element)
		if key == "max_consumers" {
			c.MaxConsumers = int(element.(float64))
		} else if key == "max_streams" {
			c.MaxStreams = int(element.(float64))
		} else if key == "max_file" {
			c.MaxFile = int(element.(float64))
		} else if key == "max_mem" {
			c.MaxMem = int(element.(float64))
		}
	}
	return nil
}
