// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// GetOpsNoteByIDReader is a Reader for the GetOpsNoteByID structure.
type GetOpsNoteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOpsNoteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOpsNoteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetOpsNoteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOpsNoteByIDOK creates a GetOpsNoteByIDOK with default headers values
func NewGetOpsNoteByIDOK() *GetOpsNoteByIDOK {
	return &GetOpsNoteByIDOK{}
}

/*GetOpsNoteByIDOK handles this case with default header values.

successful operation
*/
type GetOpsNoteByIDOK struct {
	Payload *models.OpsNote
}

func (o *GetOpsNoteByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/opsnotes/{id}][%d] getOpsNoteByIdOK  %+v", 200, o.Payload)
}

func (o *GetOpsNoteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpsNote)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOpsNoteByIDDefault creates a GetOpsNoteByIDDefault with default headers values
func NewGetOpsNoteByIDDefault(code int) *GetOpsNoteByIDDefault {
	return &GetOpsNoteByIDDefault{
		_statusCode: code,
	}
}

/*GetOpsNoteByIDDefault handles this case with default header values.

Error
*/
type GetOpsNoteByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get ops note by Id default response
func (o *GetOpsNoteByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetOpsNoteByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/opsnotes/{id}][%d] getOpsNoteById default  %+v", o._statusCode, o.Payload)
}

func (o *GetOpsNoteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
