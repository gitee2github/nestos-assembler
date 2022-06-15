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

// NotificationEndpointListOKCode is the HTTP code returned for type NotificationEndpointListOK
const NotificationEndpointListOKCode int = 200

/*NotificationEndpointListOK A successful response.

swagger:response notificationEndpointListOK
*/
type NotificationEndpointListOK struct {

	/*
	  In: Body
	*/
	Payload *models.NotifEndpointResponse `json:"body,omitempty"`
}

// NewNotificationEndpointListOK creates NotificationEndpointListOK with default headers values
func NewNotificationEndpointListOK() *NotificationEndpointListOK {

	return &NotificationEndpointListOK{}
}

// WithPayload adds the payload to the notification endpoint list o k response
func (o *NotificationEndpointListOK) WithPayload(payload *models.NotifEndpointResponse) *NotificationEndpointListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification endpoint list o k response
func (o *NotificationEndpointListOK) SetPayload(payload *models.NotifEndpointResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationEndpointListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NotificationEndpointListDefault Generic error response.

swagger:response notificationEndpointListDefault
*/
type NotificationEndpointListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNotificationEndpointListDefault creates NotificationEndpointListDefault with default headers values
func NewNotificationEndpointListDefault(code int) *NotificationEndpointListDefault {
	if code <= 0 {
		code = 500
	}

	return &NotificationEndpointListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the notification endpoint list default response
func (o *NotificationEndpointListDefault) WithStatusCode(code int) *NotificationEndpointListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the notification endpoint list default response
func (o *NotificationEndpointListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the notification endpoint list default response
func (o *NotificationEndpointListDefault) WithPayload(payload *models.Error) *NotificationEndpointListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the notification endpoint list default response
func (o *NotificationEndpointListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NotificationEndpointListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
