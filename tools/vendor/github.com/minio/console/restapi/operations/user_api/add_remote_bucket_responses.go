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

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// AddRemoteBucketCreatedCode is the HTTP code returned for type AddRemoteBucketCreated
const AddRemoteBucketCreatedCode int = 201

/*AddRemoteBucketCreated A successful response.

swagger:response addRemoteBucketCreated
*/
type AddRemoteBucketCreated struct {
}

// NewAddRemoteBucketCreated creates AddRemoteBucketCreated with default headers values
func NewAddRemoteBucketCreated() *AddRemoteBucketCreated {

	return &AddRemoteBucketCreated{}
}

// WriteResponse to the client
func (o *AddRemoteBucketCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*AddRemoteBucketDefault Generic error response.

swagger:response addRemoteBucketDefault
*/
type AddRemoteBucketDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddRemoteBucketDefault creates AddRemoteBucketDefault with default headers values
func NewAddRemoteBucketDefault(code int) *AddRemoteBucketDefault {
	if code <= 0 {
		code = 500
	}

	return &AddRemoteBucketDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add remote bucket default response
func (o *AddRemoteBucketDefault) WithStatusCode(code int) *AddRemoteBucketDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add remote bucket default response
func (o *AddRemoteBucketDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add remote bucket default response
func (o *AddRemoteBucketDefault) WithPayload(payload *models.Error) *AddRemoteBucketDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add remote bucket default response
func (o *AddRemoteBucketDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddRemoteBucketDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
