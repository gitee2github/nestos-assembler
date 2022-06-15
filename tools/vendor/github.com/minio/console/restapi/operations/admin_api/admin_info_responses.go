// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// AdminInfoOKCode is the HTTP code returned for type AdminInfoOK
const AdminInfoOKCode int = 200

/*AdminInfoOK A successful response.

swagger:response adminInfoOK
*/
type AdminInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.AdminInfoResponse `json:"body,omitempty"`
}

// NewAdminInfoOK creates AdminInfoOK with default headers values
func NewAdminInfoOK() *AdminInfoOK {

	return &AdminInfoOK{}
}

// WithPayload adds the payload to the admin info o k response
func (o *AdminInfoOK) WithPayload(payload *models.AdminInfoResponse) *AdminInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin info o k response
func (o *AdminInfoOK) SetPayload(payload *models.AdminInfoResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AdminInfoDefault Generic error response.

swagger:response adminInfoDefault
*/
type AdminInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAdminInfoDefault creates AdminInfoDefault with default headers values
func NewAdminInfoDefault(code int) *AdminInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &AdminInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the admin info default response
func (o *AdminInfoDefault) WithStatusCode(code int) *AdminInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the admin info default response
func (o *AdminInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the admin info default response
func (o *AdminInfoDefault) WithPayload(payload *models.Error) *AdminInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the admin info default response
func (o *AdminInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AdminInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
