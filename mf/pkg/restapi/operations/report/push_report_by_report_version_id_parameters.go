// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "mlss-mf/pkg/models"
)

// NewPushReportByReportVersionIDParams creates a new PushReportByReportVersionIDParams object
// no default values defined in spec.
func NewPushReportByReportVersionIDParams() PushReportByReportVersionIDParams {

	return PushReportByReportVersionIDParams{}
}

// PushReportByReportVersionIDParams contains all the bound params for the push report by report version Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters pushReportByReportVersionId
type PushReportByReportVersionIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The report push event Request
	  Required: true
	  In: body
	*/
	ReportPushEvent *models.PushReportRequest
	/*
	  Required: true
	  In: path
	*/
	ReportVersionID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPushReportByReportVersionIDParams() beforehand.
func (o *PushReportByReportVersionIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PushReportRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("reportPushEvent", "body"))
			} else {
				res = append(res, errors.NewParseError("reportPushEvent", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ReportPushEvent = &body
			}
		}
	} else {
		res = append(res, errors.Required("reportPushEvent", "body"))
	}
	rReportVersionID, rhkReportVersionID, _ := route.Params.GetOK("reportVersionId")
	if err := o.bindReportVersionID(rReportVersionID, rhkReportVersionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindReportVersionID binds and validates parameter ReportVersionID from path.
func (o *PushReportByReportVersionIDParams) bindReportVersionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("reportVersionId", "path", "int64", raw)
	}
	o.ReportVersionID = value

	return nil
}