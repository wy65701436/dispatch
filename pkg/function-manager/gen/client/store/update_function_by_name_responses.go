///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"gitlab.eng.vmware.com/serverless/serverless/pkg/function-manager/gen/models"
)

// UpdateFunctionByNameReader is a Reader for the UpdateFunctionByName structure.
type UpdateFunctionByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFunctionByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateFunctionByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateFunctionByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateFunctionByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateFunctionByNameOK creates a UpdateFunctionByNameOK with default headers values
func NewUpdateFunctionByNameOK() *UpdateFunctionByNameOK {
	return &UpdateFunctionByNameOK{}
}

/*UpdateFunctionByNameOK handles this case with default header values.

Successful update
*/
type UpdateFunctionByNameOK struct {
	Payload *models.Function
}

func (o *UpdateFunctionByNameOK) Error() string {
	return fmt.Sprintf("[PATCH /{functionName}][%d] updateFunctionByNameOK  %+v", 200, o.Payload)
}

func (o *UpdateFunctionByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Function)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionByNameBadRequest creates a UpdateFunctionByNameBadRequest with default headers values
func NewUpdateFunctionByNameBadRequest() *UpdateFunctionByNameBadRequest {
	return &UpdateFunctionByNameBadRequest{}
}

/*UpdateFunctionByNameBadRequest handles this case with default header values.

Invalid input
*/
type UpdateFunctionByNameBadRequest struct {
	Payload *models.Error
}

func (o *UpdateFunctionByNameBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /{functionName}][%d] updateFunctionByNameBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateFunctionByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFunctionByNameNotFound creates a UpdateFunctionByNameNotFound with default headers values
func NewUpdateFunctionByNameNotFound() *UpdateFunctionByNameNotFound {
	return &UpdateFunctionByNameNotFound{}
}

/*UpdateFunctionByNameNotFound handles this case with default header values.

Function not found
*/
type UpdateFunctionByNameNotFound struct {
	Payload *models.Error
}

func (o *UpdateFunctionByNameNotFound) Error() string {
	return fmt.Sprintf("[PATCH /{functionName}][%d] updateFunctionByNameNotFound  %+v", 404, o.Payload)
}

func (o *UpdateFunctionByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
