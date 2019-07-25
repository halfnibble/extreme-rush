package main

import (
	"github.com/santhosh-tekuri/jsonschema"
)

func parseSchema() (jsonschema.Schema, error) {
	rushSchema, err := jsonschema.Compile("schemas/rush.schema.json")
	if err != nil {
		return *rushSchema, err
	}
	return *rushSchema, nil
}
