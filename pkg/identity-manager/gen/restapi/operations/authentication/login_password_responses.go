///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/identity-manager/gen/models"
)

// LoginPasswordOKCode is the HTTP code returned for type LoginPasswordOK
const LoginPasswordOKCode int = 200

/*LoginPasswordOK OK

swagger:response loginPasswordOK
*/
type LoginPasswordOK struct {
	/*
	  Required: true
	*/
	SetCookie string `json:"Set-Cookie"`

	/*
	  In: Body
	*/
	Payload *models.Auth `json:"body,omitempty"`
}

// NewLoginPasswordOK creates LoginPasswordOK with default headers values
func NewLoginPasswordOK() *LoginPasswordOK {
	return &LoginPasswordOK{}
}

// WithSetCookie adds the setCookie to the login password o k response
func (o *LoginPasswordOK) WithSetCookie(setCookie string) *LoginPasswordOK {
	o.SetCookie = setCookie
	return o
}

// SetSetCookie sets the setCookie to the login password o k response
func (o *LoginPasswordOK) SetSetCookie(setCookie string) {
	o.SetCookie = setCookie
}

// WithPayload adds the payload to the login password o k response
func (o *LoginPasswordOK) WithPayload(payload *models.Auth) *LoginPasswordOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login password o k response
func (o *LoginPasswordOK) SetPayload(payload *models.Auth) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginPasswordOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Set-Cookie

	setCookie := o.SetCookie
	if setCookie != "" {
		rw.Header().Set("Set-Cookie", setCookie)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LoginPasswordDefault error

swagger:response loginPasswordDefault
*/
type LoginPasswordDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoginPasswordDefault creates LoginPasswordDefault with default headers values
func NewLoginPasswordDefault(code int) *LoginPasswordDefault {
	if code <= 0 {
		code = 500
	}

	return &LoginPasswordDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the login password default response
func (o *LoginPasswordDefault) WithStatusCode(code int) *LoginPasswordDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the login password default response
func (o *LoginPasswordDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the login password default response
func (o *LoginPasswordDefault) WithPayload(payload *models.Error) *LoginPasswordDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the login password default response
func (o *LoginPasswordDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoginPasswordDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
