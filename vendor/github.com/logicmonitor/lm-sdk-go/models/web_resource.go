// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebResource web resource
// swagger:model WebResource
type WebResource struct {

	// If type = html this should be a url, if type = iframe this should be an iframe
	// Required: true
	URL *string `json:"URL"`

	// html | iframe
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this web resource
func (m *WebResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebResource) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("URL", "body", m.URL); err != nil {
		return err
	}

	return nil
}

func (m *WebResource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebResource) UnmarshalBinary(b []byte) error {
	var res WebResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
