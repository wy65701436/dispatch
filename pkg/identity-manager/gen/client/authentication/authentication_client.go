///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Login ifs not authenticated redirect to login endpoint otherwise go to the home page
*/
func (a *Client) Login(params *LoginParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "login",
		Method:             "GET",
		PathPattern:        "/v1/iam",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
LoginPassword users logs in with username and password the credientials are forwarded to the external identity provider to exchange for auth token
*/
func (a *Client) LoginPassword(params *LoginPasswordParams) (*LoginPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginPassword",
		Method:             "GET",
		PathPattern:        "/v1/iam/login/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginPasswordOK), nil

}

/*
LoginVmware thes URL to which the identity provider e g v ID m will redirect the browser after authorization has been granted by the user
*/
func (a *Client) LoginVmware(params *LoginVmwareParams) (*LoginVmwareOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginVmwareParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginVmware",
		Method:             "GET",
		PathPattern:        "/login/vmware",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginVmwareReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoginVmwareOK), nil

}

/*
Logout logouts the current user by clearing application session cookies
*/
func (a *Client) Logout(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter) (*LogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "logout",
		Method:             "GET",
		PathPattern:        "/v1/iam/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LogoutOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
