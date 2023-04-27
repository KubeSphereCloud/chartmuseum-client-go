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

// GetChartsFilenameReader is a Reader for the GetChartsFilename structure.
type GetChartsFilenameReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetChartsFilenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChartsFilenameOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetChartsFilenameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetChartsFilenameOK creates a GetChartsFilenameOK with default headers values
func NewGetChartsFilenameOK(writer io.Writer) *GetChartsFilenameOK {
	return &GetChartsFilenameOK{

		Payload: writer,
	}
}

/*
GetChartsFilenameOK describes a response with status code 200, with default header values.

OK
*/
type GetChartsFilenameOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this get charts filename o k response has a 2xx status code
func (o *GetChartsFilenameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get charts filename o k response has a 3xx status code
func (o *GetChartsFilenameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get charts filename o k response has a 4xx status code
func (o *GetChartsFilenameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get charts filename o k response has a 5xx status code
func (o *GetChartsFilenameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get charts filename o k response a status code equal to that given
func (o *GetChartsFilenameOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetChartsFilenameOK) Error() string {
	return fmt.Sprintf("[GET /charts/{filename}][%d] getChartsFilenameOK  %+v", 200, o.Payload)
}

func (o *GetChartsFilenameOK) String() string {
	return fmt.Sprintf("[GET /charts/{filename}][%d] getChartsFilenameOK  %+v", 200, o.Payload)
}

func (o *GetChartsFilenameOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetChartsFilenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChartsFilenameDefault creates a GetChartsFilenameDefault with default headers values
func NewGetChartsFilenameDefault(code int) *GetChartsFilenameDefault {
	return &GetChartsFilenameDefault{
		_statusCode: code,
	}
}

/*
GetChartsFilenameDefault describes a response with status code -1, with default header values.

GetChartsFilenameDefault get charts filename default
*/
type GetChartsFilenameDefault struct {
	_statusCode int

	Payload models.GinH
}

// Code gets the status code for the get charts filename default response
func (o *GetChartsFilenameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get charts filename default response has a 2xx status code
func (o *GetChartsFilenameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get charts filename default response has a 3xx status code
func (o *GetChartsFilenameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get charts filename default response has a 4xx status code
func (o *GetChartsFilenameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get charts filename default response has a 5xx status code
func (o *GetChartsFilenameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get charts filename default response a status code equal to that given
func (o *GetChartsFilenameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetChartsFilenameDefault) Error() string {
	return fmt.Sprintf("[GET /charts/{filename}][%d] GetChartsFilename default  %+v", o._statusCode, o.Payload)
}

func (o *GetChartsFilenameDefault) String() string {
	return fmt.Sprintf("[GET /charts/{filename}][%d] GetChartsFilename default  %+v", o._statusCode, o.Payload)
}

func (o *GetChartsFilenameDefault) GetPayload() models.GinH {
	return o.Payload
}

func (o *GetChartsFilenameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
