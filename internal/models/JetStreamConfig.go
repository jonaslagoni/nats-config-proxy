package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JetStreamConfig represents a JetStreamConfig model.
type JetStreamConfig struct {
	Enabled              bool                         `json:"enabled"`
	Password             string                       `json:"password"`
	Nkey                 string                       `json:"nkey"`
	AdditionalProperties map[string]*JetStreamAccount `json:"-"`
}

func (m *JetStreamConfig) MarshalJSON() ([]byte, error) {
	jsonObj := make(map[string]interface{})
	jsonObj["enabled"] = m.Enabled
	jsonObj["password"] = m.Password
	jsonObj["nkey"] = m.Nkey

	for key, element := range m.AdditionalProperties {
		if jsonObj[key] == nil {
			jsonObj[key] = element
		} else {
			fmt.Println("Could not convert additionalProperty Key:", key, " with ", "Element:", element, " to JSON as it already exists.")
		}
	}
	return json.Marshal(jsonObj)
}

func UnmarshalJetStreamConfig(body []byte, c *JetStreamConfig) error {
	jsonObj := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonObj)
	if err != nil {
		return err
	}

	for key, element := range jsonObj {
		fmt.Println("Converting Key:", key, " with ", "Element:", element)
		if key == "enabled" {
			c.Enabled = element.(bool)
		} else if key == "password" {
			c.Password = element.(string)
		} else if key == "nkey" {
			c.Nkey = element.(string)
		} else {
			var t JetStreamAccount
			if c.AdditionalProperties == nil {
				c.AdditionalProperties = make(map[string]*JetStreamAccount)
			}
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			err := enc.Encode(element)
			if err != nil {
				return err
			}
			UnmarshalJetStreamAccount(buf.Bytes(), &t)
			c.AdditionalProperties[key] = &t
		}
	}
	return nil
}
