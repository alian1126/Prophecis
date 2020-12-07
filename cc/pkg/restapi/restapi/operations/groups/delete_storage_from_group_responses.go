// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// DeleteStorageFromGroupOKCode is the HTTP code returned for type DeleteStorageFromGroupOK
const DeleteStorageFromGroupOKCode int = 200

/*DeleteStorageFromGroupOK Detailed GroupStorage and GroupStorage information.

swagger:response deleteStorageFromGroupOK
*/
type DeleteStorageFromGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.GroupStorage `json:"body,omitempty"`
}

// NewDeleteStorageFromGroupOK creates DeleteStorageFromGroupOK with default headers values
func NewDeleteStorageFromGroupOK() *DeleteStorageFromGroupOK {

	return &DeleteStorageFromGroupOK{}
}

// WithPayload adds the payload to the delete storage from group o k response
func (o *DeleteStorageFromGroupOK) WithPayload(payload *models.GroupStorage) *DeleteStorageFromGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage from group o k response
func (o *DeleteStorageFromGroupOK) SetPayload(payload *models.GroupStorage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageFromGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageFromGroupUnauthorizedCode is the HTTP code returned for type DeleteStorageFromGroupUnauthorized
const DeleteStorageFromGroupUnauthorizedCode int = 401

/*DeleteStorageFromGroupUnauthorized Unauthorized

swagger:response deleteStorageFromGroupUnauthorized
*/
type DeleteStorageFromGroupUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageFromGroupUnauthorized creates DeleteStorageFromGroupUnauthorized with default headers values
func NewDeleteStorageFromGroupUnauthorized() *DeleteStorageFromGroupUnauthorized {

	return &DeleteStorageFromGroupUnauthorized{}
}

// WithPayload adds the payload to the delete storage from group unauthorized response
func (o *DeleteStorageFromGroupUnauthorized) WithPayload(payload *models.Error) *DeleteStorageFromGroupUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage from group unauthorized response
func (o *DeleteStorageFromGroupUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageFromGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageFromGroupNotFoundCode is the HTTP code returned for type DeleteStorageFromGroupNotFound
const DeleteStorageFromGroupNotFoundCode int = 404

/*DeleteStorageFromGroupNotFound Model with the given ID not found.

swagger:response deleteStorageFromGroupNotFound
*/
type DeleteStorageFromGroupNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStorageFromGroupNotFound creates DeleteStorageFromGroupNotFound with default headers values
func NewDeleteStorageFromGroupNotFound() *DeleteStorageFromGroupNotFound {

	return &DeleteStorageFromGroupNotFound{}
}

// WithPayload adds the payload to the delete storage from group not found response
func (o *DeleteStorageFromGroupNotFound) WithPayload(payload *models.Error) *DeleteStorageFromGroupNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage from group not found response
func (o *DeleteStorageFromGroupNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageFromGroupNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}