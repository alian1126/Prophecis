// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// GetDashboardsReader is a Reader for the GetDashboards structure.
type GetDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetDashboardsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDashboardsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDashboardsOK creates a GetDashboardsOK with default headers values
func NewGetDashboardsOK() *GetDashboardsOK {
	return &GetDashboardsOK{}
}

/*GetDashboardsOK handles this case with default header values.

OK
*/
type GetDashboardsOK struct {
	Payload *restmodels.GetDashboardsResponse
}

func (o *GetDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboards][%d] getDashboardsOK  %+v", 200, o.Payload)
}

func (o *GetDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.GetDashboardsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardsForbidden creates a GetDashboardsForbidden with default headers values
func NewGetDashboardsForbidden() *GetDashboardsForbidden {
	return &GetDashboardsForbidden{}
}

/*GetDashboardsForbidden handles this case with default header values.

Forbidden
*/
type GetDashboardsForbidden struct {
	Payload *restmodels.Error
}

func (o *GetDashboardsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/dashboards][%d] getDashboardsForbidden  %+v", 403, o.Payload)
}

func (o *GetDashboardsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardsNotFound creates a GetDashboardsNotFound with default headers values
func NewGetDashboardsNotFound() *GetDashboardsNotFound {
	return &GetDashboardsNotFound{}
}

/*GetDashboardsNotFound handles this case with default header values.

The dashboard cannot be found
*/
type GetDashboardsNotFound struct {
	Payload *restmodels.Error
}

func (o *GetDashboardsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/dashboards][%d] getDashboardsNotFound  %+v", 404, o.Payload)
}

func (o *GetDashboardsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}