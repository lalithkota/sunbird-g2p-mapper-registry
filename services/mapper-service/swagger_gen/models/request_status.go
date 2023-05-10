// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RequestStatus Request (e.g disburse, link, unlink, resolve, issue, search, verify, etc.,) status: <br> 1. rcvd: Received; Request received <br> 2. pdng: Pending; Request initiated <br> 3. succ: Success; Request successful <br> 4. rjct: Rejected; Request rejected
//
// swagger:model RequestStatus
type RequestStatus string

func NewRequestStatus(value RequestStatus) *RequestStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RequestStatus.
func (m RequestStatus) Pointer() *RequestStatus {
	return &m
}

const (

	// RequestStatusRcvd captures enum value "rcvd"
	RequestStatusRcvd RequestStatus = "rcvd"

	// RequestStatusPdng captures enum value "pdng"
	RequestStatusPdng RequestStatus = "pdng"

	// RequestStatusSucc captures enum value "succ"
	RequestStatusSucc RequestStatus = "succ"

	// RequestStatusRjct captures enum value "rjct"
	RequestStatusRjct RequestStatus = "rjct"
)

// for schema
var requestStatusEnum []interface{}

func init() {
	var res []RequestStatus
	if err := json.Unmarshal([]byte(`["rcvd","pdng","succ","rjct"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		requestStatusEnum = append(requestStatusEnum, v)
	}
}

func (m RequestStatus) validateRequestStatusEnum(path, location string, value RequestStatus) error {
	if err := validate.EnumCase(path, location, value, requestStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this request status
func (m RequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRequestStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this request status based on context it is used
func (m RequestStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}