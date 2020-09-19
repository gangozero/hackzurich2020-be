// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IngredientActionOptions ingredient action options
//
// swagger:model ingredientActionOptions
type IngredientActionOptions struct {

	// delivery
	Delivery int64 `json:"delivery,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// product ID
	ProductID string `json:"productID,omitempty"`

	// shop
	Shop string `json:"shop,omitempty"`
}

// Validate validates this ingredient action options
func (m *IngredientActionOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IngredientActionOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientActionOptions) UnmarshalBinary(b []byte) error {
	var res IngredientActionOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
