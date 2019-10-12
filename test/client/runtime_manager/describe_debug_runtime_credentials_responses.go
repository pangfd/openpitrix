// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeDebugRuntimeCredentialsReader is a Reader for the DescribeDebugRuntimeCredentials structure.
type DescribeDebugRuntimeCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDebugRuntimeCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeDebugRuntimeCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeDebugRuntimeCredentialsOK creates a DescribeDebugRuntimeCredentialsOK with default headers values
func NewDescribeDebugRuntimeCredentialsOK() *DescribeDebugRuntimeCredentialsOK {
	return &DescribeDebugRuntimeCredentialsOK{}
}

/*DescribeDebugRuntimeCredentialsOK handles this case with default header values.

A successful response.
*/
type DescribeDebugRuntimeCredentialsOK struct {
	Payload *models.OpenpitrixDescribeRuntimeCredentialsResponse
}

func (o *DescribeDebugRuntimeCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v1/debug_runtimes/credentials][%d] describeDebugRuntimeCredentialsOK  %+v", 200, o.Payload)
}

func (o *DescribeDebugRuntimeCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeRuntimeCredentialsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}