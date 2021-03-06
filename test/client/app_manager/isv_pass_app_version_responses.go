// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// IsvPassAppVersionReader is a Reader for the IsvPassAppVersion structure.
type IsvPassAppVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsvPassAppVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIsvPassAppVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIsvPassAppVersionOK creates a IsvPassAppVersionOK with default headers values
func NewIsvPassAppVersionOK() *IsvPassAppVersionOK {
	return &IsvPassAppVersionOK{}
}

/*IsvPassAppVersionOK handles this case with default header values.

A successful response.
*/
type IsvPassAppVersionOK struct {
	Payload *models.OpenpitrixPassAppVersionResponse
}

func (o *IsvPassAppVersionOK) Error() string {
	return fmt.Sprintf("[POST /v1/app_version/action/pass/isv][%d] isvPassAppVersionOK  %+v", 200, o.Payload)
}

func (o *IsvPassAppVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixPassAppVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
