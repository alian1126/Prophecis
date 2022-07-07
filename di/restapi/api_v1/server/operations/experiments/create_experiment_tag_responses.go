// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// CreateExperimentTagOKCode is the HTTP code returned for type CreateExperimentTagOK
const CreateExperimentTagOKCode int = 200

/*CreateExperimentTagOK OK

swagger:response createExperimentTagOK
*/
type CreateExperimentTagOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.ProphecisExperimentTag `json:"body,omitempty"`
}

// NewCreateExperimentTagOK creates CreateExperimentTagOK with default headers values
func NewCreateExperimentTagOK() *CreateExperimentTagOK {

	return &CreateExperimentTagOK{}
}

// WithPayload adds the payload to the create experiment tag o k response
func (o *CreateExperimentTagOK) WithPayload(payload *restmodels.ProphecisExperimentTag) *CreateExperimentTagOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment tag o k response
func (o *CreateExperimentTagOK) SetPayload(payload *restmodels.ProphecisExperimentTag) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExperimentTagUnauthorizedCode is the HTTP code returned for type CreateExperimentTagUnauthorized
const CreateExperimentTagUnauthorizedCode int = 401

/*CreateExperimentTagUnauthorized Unauthorized

swagger:response createExperimentTagUnauthorized
*/
type CreateExperimentTagUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewCreateExperimentTagUnauthorized creates CreateExperimentTagUnauthorized with default headers values
func NewCreateExperimentTagUnauthorized() *CreateExperimentTagUnauthorized {

	return &CreateExperimentTagUnauthorized{}
}

// WithPayload adds the payload to the create experiment tag unauthorized response
func (o *CreateExperimentTagUnauthorized) WithPayload(payload *restmodels.Error) *CreateExperimentTagUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment tag unauthorized response
func (o *CreateExperimentTagUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentTagUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExperimentTagNotFoundCode is the HTTP code returned for type CreateExperimentTagNotFound
const CreateExperimentTagNotFoundCode int = 404

/*CreateExperimentTagNotFound The Models cannot be found

swagger:response createExperimentTagNotFound
*/
type CreateExperimentTagNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewCreateExperimentTagNotFound creates CreateExperimentTagNotFound with default headers values
func NewCreateExperimentTagNotFound() *CreateExperimentTagNotFound {

	return &CreateExperimentTagNotFound{}
}

// WithPayload adds the payload to the create experiment tag not found response
func (o *CreateExperimentTagNotFound) WithPayload(payload *restmodels.Error) *CreateExperimentTagNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create experiment tag not found response
func (o *CreateExperimentTagNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExperimentTagNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}