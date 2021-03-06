///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

package v1

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NO TESTS

// ServicePlan service plan
// swagger:model ServicePlan
type ServicePlan struct {

	// bindable
	Bindable bool `json:"bindable,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// free
	Free bool `json:"free,omitempty"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// kind
	// Read Only: true
	// Pattern: ^[\w\d\-]+$
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	// Pattern: ^[\w\d\-]+$
	Name string `json:"name,omitempty"`

	// schema
	Schema *ServicePlanSchema `json:"schema,omitempty"`
}

// Validate validates this service plan
func (m *ServicePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePlan) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServicePlan) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := validate.Pattern("kind", "body", string(m.Kind), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ServicePlan) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *ServicePlan) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {

		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePlan) UnmarshalBinary(b []byte) error {
	var res ServicePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
