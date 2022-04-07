// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// DownloadModelVersionByIDOKCode is the HTTP code returned for type DownloadModelVersionByIDOK
const DownloadModelVersionByIDOKCode int = 200

/*DownloadModelVersionByIDOK OK

swagger:response downloadModelVersionByIdOK
*/
type DownloadModelVersionByIDOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadModelVersionByIDOK creates DownloadModelVersionByIDOK with default headers values
func NewDownloadModelVersionByIDOK() *DownloadModelVersionByIDOK {

	return &DownloadModelVersionByIDOK{}
}

// WithPayload adds the payload to the download model version by Id o k response
func (o *DownloadModelVersionByIDOK) WithPayload(payload io.ReadCloser) *DownloadModelVersionByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model version by Id o k response
func (o *DownloadModelVersionByIDOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelVersionByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadModelVersionByIDUnauthorizedCode is the HTTP code returned for type DownloadModelVersionByIDUnauthorized
const DownloadModelVersionByIDUnauthorizedCode int = 401

/*DownloadModelVersionByIDUnauthorized Unauthorized

swagger:response downloadModelVersionByIdUnauthorized
*/
type DownloadModelVersionByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadModelVersionByIDUnauthorized creates DownloadModelVersionByIDUnauthorized with default headers values
func NewDownloadModelVersionByIDUnauthorized() *DownloadModelVersionByIDUnauthorized {

	return &DownloadModelVersionByIDUnauthorized{}
}

// WithPayload adds the payload to the download model version by Id unauthorized response
func (o *DownloadModelVersionByIDUnauthorized) WithPayload(payload *models.Error) *DownloadModelVersionByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model version by Id unauthorized response
func (o *DownloadModelVersionByIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelVersionByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadModelVersionByIDNotFoundCode is the HTTP code returned for type DownloadModelVersionByIDNotFound
const DownloadModelVersionByIDNotFoundCode int = 404

/*DownloadModelVersionByIDNotFound Model version download fail

swagger:response downloadModelVersionByIdNotFound
*/
type DownloadModelVersionByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadModelVersionByIDNotFound creates DownloadModelVersionByIDNotFound with default headers values
func NewDownloadModelVersionByIDNotFound() *DownloadModelVersionByIDNotFound {

	return &DownloadModelVersionByIDNotFound{}
}

// WithPayload adds the payload to the download model version by Id not found response
func (o *DownloadModelVersionByIDNotFound) WithPayload(payload *models.Error) *DownloadModelVersionByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model version by Id not found response
func (o *DownloadModelVersionByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelVersionByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}