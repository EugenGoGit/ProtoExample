package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"../../client/models"
)

// KeyServiceGetPageReader is a Reader for the KeyServiceGetPage structure.
type KeyServiceGetPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeyServiceGetPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewKeyServiceGetPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewKeyServiceGetPageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKeyServiceGetPageOK creates a KeyServiceGetPageOK with default headers values
func NewKeyServiceGetPageOK() *KeyServiceGetPageOK {
	return &KeyServiceGetPageOK{}
}

/*KeyServiceGetPageOK handles this case with default header values.

A successful response.
*/
type KeyServiceGetPageOK struct {
	Payload *models.ExampleKeyPageResponse
}

func (o *KeyServiceGetPageOK) Error() string {
	return fmt.Sprintf("[GET /v1/keys][%d] keyServiceGetPageOK  %+v", 200, o.Payload)
}

func (o *KeyServiceGetPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExampleKeyPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeyServiceGetPageDefault creates a KeyServiceGetPageDefault with default headers values
func NewKeyServiceGetPageDefault(code int) *KeyServiceGetPageDefault {
	return &KeyServiceGetPageDefault{
		_statusCode: code,
	}
}

/*KeyServiceGetPageDefault handles this case with default header values.

An unexpected error response.
*/
type KeyServiceGetPageDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the key service get page default response
func (o *KeyServiceGetPageDefault) Code() int {
	return o._statusCode
}

func (o *KeyServiceGetPageDefault) Error() string {
	return fmt.Sprintf("[GET /v1/keys][%d] KeyService_GetPage default  %+v", o._statusCode, o.Payload)
}

func (o *KeyServiceGetPageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
