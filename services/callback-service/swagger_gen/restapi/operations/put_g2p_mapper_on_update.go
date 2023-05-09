// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/sunbirdrc/callback-service/swagger_gen/models"
)

// PutG2pMapperOnUpdateHandlerFunc turns a function with the right signature into a put g2p mapper on update handler
type PutG2pMapperOnUpdateHandlerFunc func(PutG2pMapperOnUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutG2pMapperOnUpdateHandlerFunc) Handle(params PutG2pMapperOnUpdateParams) middleware.Responder {
	return fn(params)
}

// PutG2pMapperOnUpdateHandler interface for that can handle valid put g2p mapper on update params
type PutG2pMapperOnUpdateHandler interface {
	Handle(PutG2pMapperOnUpdateParams) middleware.Responder
}

// NewPutG2pMapperOnUpdate creates a new http.Handler for the put g2p mapper on update operation
func NewPutG2pMapperOnUpdate(ctx *middleware.Context, handler PutG2pMapperOnUpdateHandler) *PutG2pMapperOnUpdate {
	return &PutG2pMapperOnUpdate{Context: ctx, Handler: handler}
}

/*
	PutG2pMapperOnUpdate swagger:route POST /mapper/on-update putG2pMapperOnUpdate

FAMAP-ON-UPDT	: /mapper/on-update

Update response through callback end point
*/
type PutG2pMapperOnUpdate struct {
	Context *middleware.Context
	Handler PutG2pMapperOnUpdateHandler
}

func (o *PutG2pMapperOnUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutG2pMapperOnUpdateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutG2pMapperOnUpdateBody put g2p mapper on update body
//
// swagger:model PutG2pMapperOnUpdateBody
type PutG2pMapperOnUpdateBody struct {

	// header
	// Required: true
	Header struct {
		models.MsgCallbackHeader
	} `json:"header"`

	// message
	Message *PutG2pMapperOnUpdateParamsBodyMessage `json:"message,omitempty"`

	// signature
	Signature models.MsgSignature `json:"signature,omitempty"`
}

// Validate validates this put g2p mapper on update body
func (o *PutG2pMapperOnUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateBody) validateHeader(formats strfmt.Registry) error {

	return nil
}

func (o *PutG2pMapperOnUpdateBody) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(o.Message) { // not required
		return nil
	}

	if o.Message != nil {
		if err := o.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if err := o.Signature.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signature")
		}
		return err
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update body based on the context it is used
func (o *PutG2pMapperOnUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateBody) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (o *PutG2pMapperOnUpdateBody) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.Message != nil {
		if err := o.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Signature.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signature")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signature")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateBody) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateDefaultBody put g2p mapper on update default body
//
// swagger:model PutG2pMapperOnUpdateDefaultBody
type PutG2pMapperOnUpdateDefaultBody struct {

	// message
	Message *PutG2pMapperOnUpdateDefaultBodyMessage `json:"message,omitempty"`
}

// Validate validates this put g2p mapper on update default body
func (o *PutG2pMapperOnUpdateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBody) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(o.Message) { // not required
		return nil
	}

	if o.Message != nil {
		if err := o.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update default body based on the context it is used
func (o *PutG2pMapperOnUpdateDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBody) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.Message != nil {
		if err := o.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateDefaultBody) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateDefaultBodyMessage put g2p mapper on update default body message
//
// swagger:model PutG2pMapperOnUpdateDefaultBodyMessage
type PutG2pMapperOnUpdateDefaultBodyMessage struct {

	// ack status
	AckStatus models.Ack `json:"ack_status,omitempty"`

	// error
	Error *models.Error `json:"error,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp models.Timestamp `json:"timestamp,omitempty"`
}

// Validate validates this put g2p mapper on update default body message
func (o *PutG2pMapperOnUpdateDefaultBodyMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAckStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) validateAckStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.AckStatus) { // not required
		return nil
	}

	if err := o.AckStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "ack_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "ack_status")
		}
		return err
	}

	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := o.Timestamp.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "timestamp")
		}
		return err
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update default body message based on the context it is used
func (o *PutG2pMapperOnUpdateDefaultBodyMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAckStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) contextValidateAckStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := o.AckStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "ack_status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "ack_status")
		}
		return err
	}

	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateDefaultBodyMessage) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Timestamp.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "timestamp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("put_g2p_mapper_on-update default" + "." + "message" + "." + "timestamp")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateDefaultBodyMessage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateDefaultBodyMessage) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateDefaultBodyMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateForbiddenBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PutG2pMapperOnUpdateForbiddenBody
type PutG2pMapperOnUpdateForbiddenBody struct {

	// error
	Error *PutG2pMapperOnUpdateForbiddenBodyError `json:"error,omitempty"`
}

// Validate validates this put g2p mapper on update forbidden body
func (o *PutG2pMapperOnUpdateForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateForbiddenBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateForbidden" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update forbidden body based on the context it is used
func (o *PutG2pMapperOnUpdateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateForbiddenBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateForbidden" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateForbidden" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateForbiddenBodyError put g2p mapper on update forbidden body error
//
// swagger:model PutG2pMapperOnUpdateForbiddenBodyError
type PutG2pMapperOnUpdateForbiddenBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this put g2p mapper on update forbidden body error
func (o *PutG2pMapperOnUpdateForbiddenBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put g2p mapper on update forbidden body error based on context it is used
func (o *PutG2pMapperOnUpdateForbiddenBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateForbiddenBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateForbiddenBodyError) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateForbiddenBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateInternalServerErrorBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PutG2pMapperOnUpdateInternalServerErrorBody
type PutG2pMapperOnUpdateInternalServerErrorBody struct {

	// error
	Error *PutG2pMapperOnUpdateInternalServerErrorBodyError `json:"error,omitempty"`
}

// Validate validates this put g2p mapper on update internal server error body
func (o *PutG2pMapperOnUpdateInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateInternalServerErrorBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update internal server error body based on the context it is used
func (o *PutG2pMapperOnUpdateInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateInternalServerErrorBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateInternalServerError" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateInternalServerError" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateInternalServerErrorBodyError put g2p mapper on update internal server error body error
//
// swagger:model PutG2pMapperOnUpdateInternalServerErrorBodyError
type PutG2pMapperOnUpdateInternalServerErrorBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this put g2p mapper on update internal server error body error
func (o *PutG2pMapperOnUpdateInternalServerErrorBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put g2p mapper on update internal server error body error based on context it is used
func (o *PutG2pMapperOnUpdateInternalServerErrorBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateInternalServerErrorBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateInternalServerErrorBodyError) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateInternalServerErrorBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateParamsBodyMessage put g2p mapper on update params body message
//
// swagger:model PutG2pMapperOnUpdateParamsBodyMessage
type PutG2pMapperOnUpdateParamsBodyMessage struct {

	// transaction id
	// Required: true
	TransactionID *models.TransactionID `json:"transaction_id"`

	// update response
	// Required: true
	UpdateResponse []*models.UpdateResponse `json:"update_response"`
}

// Validate validates this put g2p mapper on update params body message
func (o *PutG2pMapperOnUpdateParamsBodyMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTransactionID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateParamsBodyMessage) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"message"+"."+"transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"message"+"."+"transaction_id", "body", o.TransactionID); err != nil {
		return err
	}

	if o.TransactionID != nil {
		if err := o.TransactionID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message" + "." + "transaction_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message" + "." + "transaction_id")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateParamsBodyMessage) validateUpdateResponse(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"message"+"."+"update_response", "body", o.UpdateResponse); err != nil {
		return err
	}

	for i := 0; i < len(o.UpdateResponse); i++ {
		if swag.IsZero(o.UpdateResponse[i]) { // not required
			continue
		}

		if o.UpdateResponse[i] != nil {
			if err := o.UpdateResponse[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "message" + "." + "update_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "message" + "." + "update_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this put g2p mapper on update params body message based on the context it is used
func (o *PutG2pMapperOnUpdateParamsBodyMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTransactionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUpdateResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateParamsBodyMessage) contextValidateTransactionID(ctx context.Context, formats strfmt.Registry) error {

	if o.TransactionID != nil {
		if err := o.TransactionID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "message" + "." + "transaction_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "message" + "." + "transaction_id")
			}
			return err
		}
	}

	return nil
}

func (o *PutG2pMapperOnUpdateParamsBodyMessage) contextValidateUpdateResponse(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.UpdateResponse); i++ {

		if o.UpdateResponse[i] != nil {
			if err := o.UpdateResponse[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "message" + "." + "update_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "message" + "." + "update_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateParamsBodyMessage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateParamsBodyMessage) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateParamsBodyMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateUnauthorizedBody HTTP transport layer error codes. Used by components like gateways, LB responding with HTTP status codes 1xx, 2xx, 3xx, 4xx and 5xx
//
// swagger:model PutG2pMapperOnUpdateUnauthorizedBody
type PutG2pMapperOnUpdateUnauthorizedBody struct {

	// error
	Error *PutG2pMapperOnUpdateUnauthorizedBodyError `json:"error,omitempty"`
}

// Validate validates this put g2p mapper on update unauthorized body
func (o *PutG2pMapperOnUpdateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateUnauthorizedBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateUnauthorized" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateUnauthorized" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put g2p mapper on update unauthorized body based on the context it is used
func (o *PutG2pMapperOnUpdateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutG2pMapperOnUpdateUnauthorizedBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putG2pMapperOnUpdateUnauthorized" + "." + "error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putG2pMapperOnUpdateUnauthorized" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutG2pMapperOnUpdateUnauthorizedBodyError put g2p mapper on update unauthorized body error
//
// swagger:model PutG2pMapperOnUpdateUnauthorizedBodyError
type PutG2pMapperOnUpdateUnauthorizedBodyError struct {

	// error code
	Code string `json:"code,omitempty"`

	// error message
	Message string `json:"message,omitempty"`
}

// Validate validates this put g2p mapper on update unauthorized body error
func (o *PutG2pMapperOnUpdateUnauthorizedBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put g2p mapper on update unauthorized body error based on context it is used
func (o *PutG2pMapperOnUpdateUnauthorizedBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateUnauthorizedBodyError) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutG2pMapperOnUpdateUnauthorizedBodyError) UnmarshalBinary(b []byte) error {
	var res PutG2pMapperOnUpdateUnauthorizedBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
