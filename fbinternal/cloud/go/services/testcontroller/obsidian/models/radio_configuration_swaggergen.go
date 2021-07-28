// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RadioConfiguration Radio Configuration for eNodeb devices.
// swagger:model radioConfiguration
type RadioConfiguration struct {

	// cell configuration
	CellConfiguration *CellConfiguration `json:"CellConfiguration,omitempty"`

	// power control parameters
	PowerControlParameters *PowerControlParameters `json:"PowerControlParameters,omitempty"`
}

// Validate validates this radio configuration
func (m *RadioConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCellConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerControlParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RadioConfiguration) validateCellConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.CellConfiguration) { // not required
		return nil
	}

	if m.CellConfiguration != nil {
		if err := m.CellConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CellConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *RadioConfiguration) validatePowerControlParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerControlParameters) { // not required
		return nil
	}

	if m.PowerControlParameters != nil {
		if err := m.PowerControlParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("PowerControlParameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RadioConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RadioConfiguration) UnmarshalBinary(b []byte) error {
	var res RadioConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
