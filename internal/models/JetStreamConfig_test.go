package api

import (
	"reflect"
	"testing"
)

func TestHealthz(t *testing.T) {
	acc := &JetStreamAccount{MaxMem: 123}
	accMap := make(map[string]*JetStreamAccount)
	accMap["test"] = acc
	message := &JetStreamConfig{Enabled: true, AdditionalProperties: accMap}
	jsonByte, marshalErr := message.MarshalJSON()

	if marshalErr != nil {
		t.Fatalf("Unexpected error when marshalling JSON: %s", marshalErr)
	}

	var config JetStreamConfig
	unMarshalErr := UnmarshalJetStreamConfig(jsonByte, &config)

	if unMarshalErr != nil {
		t.Fatalf("Unexpected error when unmarshalling JSON: %s", unMarshalErr)
	}
	wasEqual := reflect.DeepEqual(*message, config)
	if !wasEqual {
		t.Fatalf("Was not equal to each other")
	}
}
