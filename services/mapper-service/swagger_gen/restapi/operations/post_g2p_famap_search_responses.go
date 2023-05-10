// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostG2pFamapSearchUnauthorizedCode is the HTTP code returned for type PostG2pFamapSearchUnauthorized
const PostG2pFamapSearchUnauthorizedCode int = 401

/*
PostG2pFamapSearchUnauthorized HTTP layer error details

swagger:response postG2pFamapSearchUnauthorized
*/
type PostG2pFamapSearchUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostG2pFamapSearchUnauthorizedBody `json:"body,omitempty"`
}

// NewPostG2pFamapSearchUnauthorized creates PostG2pFamapSearchUnauthorized with default headers values
func NewPostG2pFamapSearchUnauthorized() *PostG2pFamapSearchUnauthorized {

	return &PostG2pFamapSearchUnauthorized{}
}

// WithPayload adds the payload to the post g2p famap search unauthorized response
func (o *PostG2pFamapSearchUnauthorized) WithPayload(payload *PostG2pFamapSearchUnauthorizedBody) *PostG2pFamapSearchUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p famap search unauthorized response
func (o *PostG2pFamapSearchUnauthorized) SetPayload(payload *PostG2pFamapSearchUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pFamapSearchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pFamapSearchForbiddenCode is the HTTP code returned for type PostG2pFamapSearchForbidden
const PostG2pFamapSearchForbiddenCode int = 403

/*
PostG2pFamapSearchForbidden HTTP layer error details

swagger:response postG2pFamapSearchForbidden
*/
type PostG2pFamapSearchForbidden struct {

	/*
	  In: Body
	*/
	Payload *PostG2pFamapSearchForbiddenBody `json:"body,omitempty"`
}

// NewPostG2pFamapSearchForbidden creates PostG2pFamapSearchForbidden with default headers values
func NewPostG2pFamapSearchForbidden() *PostG2pFamapSearchForbidden {

	return &PostG2pFamapSearchForbidden{}
}

// WithPayload adds the payload to the post g2p famap search forbidden response
func (o *PostG2pFamapSearchForbidden) WithPayload(payload *PostG2pFamapSearchForbiddenBody) *PostG2pFamapSearchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p famap search forbidden response
func (o *PostG2pFamapSearchForbidden) SetPayload(payload *PostG2pFamapSearchForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pFamapSearchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostG2pFamapSearchInternalServerErrorCode is the HTTP code returned for type PostG2pFamapSearchInternalServerError
const PostG2pFamapSearchInternalServerErrorCode int = 500

/*
PostG2pFamapSearchInternalServerError HTTP layer error details

swagger:response postG2pFamapSearchInternalServerError
*/
type PostG2pFamapSearchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostG2pFamapSearchInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostG2pFamapSearchInternalServerError creates PostG2pFamapSearchInternalServerError with default headers values
func NewPostG2pFamapSearchInternalServerError() *PostG2pFamapSearchInternalServerError {

	return &PostG2pFamapSearchInternalServerError{}
}

// WithPayload adds the payload to the post g2p famap search internal server error response
func (o *PostG2pFamapSearchInternalServerError) WithPayload(payload *PostG2pFamapSearchInternalServerErrorBody) *PostG2pFamapSearchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p famap search internal server error response
func (o *PostG2pFamapSearchInternalServerError) SetPayload(payload *PostG2pFamapSearchInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pFamapSearchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostG2pFamapSearchDefault Acknowledgement of message received after successful validation of message and signature

swagger:response postG2pFamapSearchDefault
*/
type PostG2pFamapSearchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *PostG2pFamapSearchDefaultBody `json:"body,omitempty"`
}

// NewPostG2pFamapSearchDefault creates PostG2pFamapSearchDefault with default headers values
func NewPostG2pFamapSearchDefault(code int) *PostG2pFamapSearchDefault {
	if code <= 0 {
		code = 500
	}

	return &PostG2pFamapSearchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post g2p famap search default response
func (o *PostG2pFamapSearchDefault) WithStatusCode(code int) *PostG2pFamapSearchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post g2p famap search default response
func (o *PostG2pFamapSearchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post g2p famap search default response
func (o *PostG2pFamapSearchDefault) WithPayload(payload *PostG2pFamapSearchDefaultBody) *PostG2pFamapSearchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post g2p famap search default response
func (o *PostG2pFamapSearchDefault) SetPayload(payload *PostG2pFamapSearchDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostG2pFamapSearchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}