///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// AddDriverCreatedCode is the HTTP code returned for type AddDriverCreated
const AddDriverCreatedCode int = 201

/*AddDriverCreated Driver created

swagger:response addDriverCreated
*/
type AddDriverCreated struct {

	/*
	  In: Body
	*/
	Payload *v1.EventDriver `json:"body,omitempty"`
}

// NewAddDriverCreated creates AddDriverCreated with default headers values
func NewAddDriverCreated() *AddDriverCreated {

	return &AddDriverCreated{}
}

// WithPayload adds the payload to the add driver created response
func (o *AddDriverCreated) WithPayload(payload *v1.EventDriver) *AddDriverCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver created response
func (o *AddDriverCreated) SetPayload(payload *v1.EventDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverBadRequestCode is the HTTP code returned for type AddDriverBadRequest
const AddDriverBadRequestCode int = 400

/*AddDriverBadRequest Invalid input

swagger:response addDriverBadRequest
*/
type AddDriverBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddDriverBadRequest creates AddDriverBadRequest with default headers values
func NewAddDriverBadRequest() *AddDriverBadRequest {

	return &AddDriverBadRequest{}
}

// WithPayload adds the payload to the add driver bad request response
func (o *AddDriverBadRequest) WithPayload(payload *v1.Error) *AddDriverBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver bad request response
func (o *AddDriverBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverUnauthorizedCode is the HTTP code returned for type AddDriverUnauthorized
const AddDriverUnauthorizedCode int = 401

/*AddDriverUnauthorized Unauthorized Request

swagger:response addDriverUnauthorized
*/
type AddDriverUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddDriverUnauthorized creates AddDriverUnauthorized with default headers values
func NewAddDriverUnauthorized() *AddDriverUnauthorized {

	return &AddDriverUnauthorized{}
}

// WithPayload adds the payload to the add driver unauthorized response
func (o *AddDriverUnauthorized) WithPayload(payload *v1.Error) *AddDriverUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver unauthorized response
func (o *AddDriverUnauthorized) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverForbiddenCode is the HTTP code returned for type AddDriverForbidden
const AddDriverForbiddenCode int = 403

/*AddDriverForbidden access to this resource is forbidden

swagger:response addDriverForbidden
*/
type AddDriverForbidden struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddDriverForbidden creates AddDriverForbidden with default headers values
func NewAddDriverForbidden() *AddDriverForbidden {

	return &AddDriverForbidden{}
}

// WithPayload adds the payload to the add driver forbidden response
func (o *AddDriverForbidden) WithPayload(payload *v1.Error) *AddDriverForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver forbidden response
func (o *AddDriverForbidden) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDriverConflictCode is the HTTP code returned for type AddDriverConflict
const AddDriverConflictCode int = 409

/*AddDriverConflict Already Exists

swagger:response addDriverConflict
*/
type AddDriverConflict struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddDriverConflict creates AddDriverConflict with default headers values
func NewAddDriverConflict() *AddDriverConflict {

	return &AddDriverConflict{}
}

// WithPayload adds the payload to the add driver conflict response
func (o *AddDriverConflict) WithPayload(payload *v1.Error) *AddDriverConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver conflict response
func (o *AddDriverConflict) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddDriverDefault Unknown error

swagger:response addDriverDefault
*/
type AddDriverDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewAddDriverDefault creates AddDriverDefault with default headers values
func NewAddDriverDefault(code int) *AddDriverDefault {
	if code <= 0 {
		code = 500
	}

	return &AddDriverDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add driver default response
func (o *AddDriverDefault) WithStatusCode(code int) *AddDriverDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add driver default response
func (o *AddDriverDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add driver default response
func (o *AddDriverDefault) WithPayload(payload *v1.Error) *AddDriverDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add driver default response
func (o *AddDriverDefault) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDriverDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}