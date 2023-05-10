// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// LinkRequestName End-User's full name in displayable form including all name parts, possibly including titles and suffixes, ordered according to the End-User's locale and preferences.
// Example: Mr. John Smith
//
// swagger:model linkRequestName
type LinkRequestName string

// Validate validates this link request name
func (m LinkRequestName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link request name based on context it is used
func (m LinkRequestName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}