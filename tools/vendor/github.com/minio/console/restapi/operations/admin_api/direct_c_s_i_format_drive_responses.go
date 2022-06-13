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

// DirectCSIFormatDriveOKCode is the HTTP code returned for type DirectCSIFormatDriveOK
const DirectCSIFormatDriveOKCode int = 200

/*DirectCSIFormatDriveOK A successful response.

swagger:response directCSIFormatDriveOK
*/
type DirectCSIFormatDriveOK struct {

	/*
	  In: Body
	*/
	Payload *models.FormatDirectCSIDrivesResponse `json:"body,omitempty"`
}

// NewDirectCSIFormatDriveOK creates DirectCSIFormatDriveOK with default headers values
func NewDirectCSIFormatDriveOK() *DirectCSIFormatDriveOK {

	return &DirectCSIFormatDriveOK{}
}

// WithPayload adds the payload to the direct c s i format drive o k response
func (o *DirectCSIFormatDriveOK) WithPayload(payload *models.FormatDirectCSIDrivesResponse) *DirectCSIFormatDriveOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the direct c s i format drive o k response
func (o *DirectCSIFormatDriveOK) SetPayload(payload *models.FormatDirectCSIDrivesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DirectCSIFormatDriveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DirectCSIFormatDriveDefault Generic error response.

swagger:response directCSIFormatDriveDefault
*/
type DirectCSIFormatDriveDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDirectCSIFormatDriveDefault creates DirectCSIFormatDriveDefault with default headers values
func NewDirectCSIFormatDriveDefault(code int) *DirectCSIFormatDriveDefault {
	if code <= 0 {
		code = 500
	}

	return &DirectCSIFormatDriveDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the direct c s i format drive default response
func (o *DirectCSIFormatDriveDefault) WithStatusCode(code int) *DirectCSIFormatDriveDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the direct c s i format drive default response
func (o *DirectCSIFormatDriveDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the direct c s i format drive default response
func (o *DirectCSIFormatDriveDefault) WithPayload(payload *models.Error) *DirectCSIFormatDriveDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the direct c s i format drive default response
func (o *DirectCSIFormatDriveDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DirectCSIFormatDriveDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
