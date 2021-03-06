package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExampleKeyPageResponse Response for GetPage
// swagger:model exampleKeyPageResponse
type ExampleKeyPageResponse struct {

	// count
	Count int32 `json:"count,omitempty"`

	// is first
	IsFirst bool `json:"is_first,omitempty"`

	// is last
	IsLast bool `json:"is_last,omitempty"`

	// keys
	Keys []*ExampleV1RfidDto `json:"keys"`

	// page number
	PageNumber int32 `json:"page_number,omitempty"`

	// page size
	PageSize int32 `json:"page_size,omitempty"`

	// total count
	TotalCount int32 `json:"total_count,omitempty"`
}

// Validate validates this example key page response
func (m *ExampleKeyPageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeys(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExampleKeyPageResponse) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {

		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {

			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
