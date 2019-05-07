// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CellData cell data
// swagger:model CellData
type CellData struct {

	// alert severity
	AlertSeverity string `json:"alertSeverity,omitempty"`

	// alert status
	// Read Only: true
	AlertStatus string `json:"alertStatus,omitempty"`

	// days until alert list
	DaysUntilAlertList []*DaysUntilAlert `json:"daysUntilAlertList,omitempty"`

	// forecast day
	// Read Only: true
	ForecastDay int32 `json:"forecastDay,omitempty"`

	// instance Id
	// Read Only: true
	InstanceID int32 `json:"instanceId,omitempty"`

	// instance name
	// Read Only: true
	InstanceName string `json:"instanceName,omitempty"`

	// value
	// Read Only: true
	Value float64 `json:"value,omitempty"`
}

// Validate validates this cell data
func (m *CellData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysUntilAlertList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CellData) validateDaysUntilAlertList(formats strfmt.Registry) error {

	if swag.IsZero(m.DaysUntilAlertList) { // not required
		return nil
	}

	for i := 0; i < len(m.DaysUntilAlertList); i++ {
		if swag.IsZero(m.DaysUntilAlertList[i]) { // not required
			continue
		}

		if m.DaysUntilAlertList[i] != nil {
			if err := m.DaysUntilAlertList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("daysUntilAlertList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CellData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CellData) UnmarshalBinary(b []byte) error {
	var res CellData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
