package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ExampleV1RfidDto Represents the key
// swagger:model exampleV1RfidDto
type ExampleV1RfidDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`
}

// Validate validates this example v1 rfid dto
func (m *ExampleV1RfidDto) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
