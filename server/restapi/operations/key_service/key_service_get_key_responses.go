package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"../../server/models"
)

// KeyServiceGetKeyOKCode is the HTTP code returned for type KeyServiceGetKeyOK
const KeyServiceGetKeyOKCode int = 200

/*KeyServiceGetKeyOK A successful response.

swagger:response keyServiceGetKeyOK
*/
type KeyServiceGetKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.ExampleV1RfidDto `json:"body,omitempty"`
}

// NewKeyServiceGetKeyOK creates KeyServiceGetKeyOK with default headers values
func NewKeyServiceGetKeyOK() *KeyServiceGetKeyOK {
	return &KeyServiceGetKeyOK{}
}

// WithPayload adds the payload to the key service get key o k response
func (o *KeyServiceGetKeyOK) WithPayload(payload *models.ExampleV1RfidDto) *KeyServiceGetKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service get key o k response
func (o *KeyServiceGetKeyOK) SetPayload(payload *models.ExampleV1RfidDto) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceGetKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*KeyServiceGetKeyDefault An unexpected error response.

swagger:response keyServiceGetKeyDefault
*/
type KeyServiceGetKeyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewKeyServiceGetKeyDefault creates KeyServiceGetKeyDefault with default headers values
func NewKeyServiceGetKeyDefault(code int) *KeyServiceGetKeyDefault {
	if code <= 0 {
		code = 500
	}

	return &KeyServiceGetKeyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the key service get key default response
func (o *KeyServiceGetKeyDefault) WithStatusCode(code int) *KeyServiceGetKeyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the key service get key default response
func (o *KeyServiceGetKeyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the key service get key default response
func (o *KeyServiceGetKeyDefault) WithPayload(payload *models.RuntimeError) *KeyServiceGetKeyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service get key default response
func (o *KeyServiceGetKeyDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceGetKeyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
