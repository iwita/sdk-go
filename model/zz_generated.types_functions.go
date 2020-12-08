// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package model

import "fmt"
import "encoding/json"

type Function struct {
	// Metadata corresponds to the JSON schema field "metadata".
	Metadata Metadata_1 `json:"metadata,omitempty"`

	// Unique function name
	Name string `json:"name"`

	// Combination of the function/service OpenAPI definition URI and the operationID
	// of the operation that needs to be invoked, separated by a '#'. For example
	// 'https://petstore.swagger.io/v2/swagger.json#getPetById'
	Operation *string `json:"operation,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Function) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain Function
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Function(plain)
	return nil
}

// Serverless Workflow specification - functions schema
type Functions map[string]interface{}
