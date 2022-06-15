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

// RemoteBucketDetailsOKCode is the HTTP code returned for type RemoteBucketDetailsOK
const RemoteBucketDetailsOKCode int = 200

/*RemoteBucketDetailsOK A successful response.

swagger:response remoteBucketDetailsOK
*/
type RemoteBucketDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.RemoteBucket `json:"body,omitempty"`
}

// NewRemoteBucketDetailsOK creates RemoteBucketDetailsOK with default headers values
func NewRemoteBucketDetailsOK() *RemoteBucketDetailsOK {

	return &RemoteBucketDetailsOK{}
}

// WithPayload adds the payload to the remote bucket details o k response
func (o *RemoteBucketDetailsOK) WithPayload(payload *models.RemoteBucket) *RemoteBucketDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remote bucket details o k response
func (o *RemoteBucketDetailsOK) SetPayload(payload *models.RemoteBucket) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoteBucketDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*RemoteBucketDetailsDefault Generic error response.

swagger:response remoteBucketDetailsDefault
*/
type RemoteBucketDetailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRemoteBucketDetailsDefault creates RemoteBucketDetailsDefault with default headers values
func NewRemoteBucketDetailsDefault(code int) *RemoteBucketDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &RemoteBucketDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the remote bucket details default response
func (o *RemoteBucketDetailsDefault) WithStatusCode(code int) *RemoteBucketDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the remote bucket details default response
func (o *RemoteBucketDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the remote bucket details default response
func (o *RemoteBucketDetailsDefault) WithPayload(payload *models.Error) *RemoteBucketDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remote bucket details default response
func (o *RemoteBucketDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoteBucketDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
