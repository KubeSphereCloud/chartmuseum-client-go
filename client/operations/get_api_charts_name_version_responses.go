// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/KubeSphereCloud/chartmuseum-client-go/models"
)

// GetAPIChartsNameVersionReader is a Reader for the GetAPIChartsNameVersion structure.
type GetAPIChartsNameVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIChartsNameVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIChartsNameVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAPIChartsNameVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAPIChartsNameVersionOK creates a GetAPIChartsNameVersionOK with default headers values
func NewGetAPIChartsNameVersionOK() *GetAPIChartsNameVersionOK {
	return &GetAPIChartsNameVersionOK{}
}

/*
GetAPIChartsNameVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetAPIChartsNameVersionOK struct {
	Payload *models.RepoChartVersion
}

// IsSuccess returns true when this get Api charts name version o k response has a 2xx status code
func (o *GetAPIChartsNameVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api charts name version o k response has a 3xx status code
func (o *GetAPIChartsNameVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api charts name version o k response has a 4xx status code
func (o *GetAPIChartsNameVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api charts name version o k response has a 5xx status code
func (o *GetAPIChartsNameVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api charts name version o k response a status code equal to that given
func (o *GetAPIChartsNameVersionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAPIChartsNameVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/charts/{name}/{version}][%d] getApiChartsNameVersionOK  %+v", 200, o.Payload)
}

func (o *GetAPIChartsNameVersionOK) String() string {
	return fmt.Sprintf("[GET /api/charts/{name}/{version}][%d] getApiChartsNameVersionOK  %+v", 200, o.Payload)
}

func (o *GetAPIChartsNameVersionOK) GetPayload() *models.RepoChartVersion {
	return o.Payload
}

func (o *GetAPIChartsNameVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepoChartVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIChartsNameVersionDefault creates a GetAPIChartsNameVersionDefault with default headers values
func NewGetAPIChartsNameVersionDefault(code int) *GetAPIChartsNameVersionDefault {
	return &GetAPIChartsNameVersionDefault{
		_statusCode: code,
	}
}

/*
GetAPIChartsNameVersionDefault describes a response with status code -1, with default header values.

GetAPIChartsNameVersionDefault get API charts name version default
*/
type GetAPIChartsNameVersionDefault struct {
	_statusCode int

	Payload models.GinH
}

// Code gets the status code for the get API charts name version default response
func (o *GetAPIChartsNameVersionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get API charts name version default response has a 2xx status code
func (o *GetAPIChartsNameVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get API charts name version default response has a 3xx status code
func (o *GetAPIChartsNameVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get API charts name version default response has a 4xx status code
func (o *GetAPIChartsNameVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get API charts name version default response has a 5xx status code
func (o *GetAPIChartsNameVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get API charts name version default response a status code equal to that given
func (o *GetAPIChartsNameVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAPIChartsNameVersionDefault) Error() string {
	return fmt.Sprintf("[GET /api/charts/{name}/{version}][%d] GetAPIChartsNameVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIChartsNameVersionDefault) String() string {
	return fmt.Sprintf("[GET /api/charts/{name}/{version}][%d] GetAPIChartsNameVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetAPIChartsNameVersionDefault) GetPayload() models.GinH {
	return o.Payload
}

func (o *GetAPIChartsNameVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
