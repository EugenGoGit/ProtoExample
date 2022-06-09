package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"../../server/models"
)

// KeyServiceAddKeyOKCode is the HTTP code returned for type KeyServiceAddKeyOK
const KeyServiceAddKeyOKCode int = 200

/*KeyServiceAddKeyOK A successful response.

swagger:response keyServiceAddKeyOK
*/
type KeyServiceAddKeyOK struct {

	/*
	  In: Body
	*/
	Payload *models.ExampleV1RfidDto `json:"body,omitempty"`
}

// NewKeyServiceAddKeyOK creates KeyServiceAddKeyOK with default headers values
func NewKeyServiceAddKeyOK() *KeyServiceAddKeyOK {
	return &KeyServiceAddKeyOK{}
}

// WithPayload adds the payload to the key service add key o k response
func (o *KeyServiceAddKeyOK) WithPayload(payload *models.ExampleV1RfidDto) *KeyServiceAddKeyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service add key o k response
func (o *KeyServiceAddKeyOK) SetPayload(payload *models.ExampleV1RfidDto) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceAddKeyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*KeyServiceAddKeyDefault An unexpected error response.

swagger:response keyServiceAddKeyDefault
*/
type KeyServiceAddKeyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewKeyServiceAddKeyDefault creates KeyServiceAddKeyDefault with default headers values
func NewKeyServiceAddKeyDefault(code int) *KeyServiceAddKeyDefault {
	if code <= 0 {
		code = 500
	}

	return &KeyServiceAddKeyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the key service add key default response
func (o *KeyServiceAddKeyDefault) WithStatusCode(code int) *KeyServiceAddKeyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the key service add key default response
func (o *KeyServiceAddKeyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the key service add key default response
func (o *KeyServiceAddKeyDefault) WithPayload(payload *models.RuntimeError) *KeyServiceAddKeyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service add key default response
func (o *KeyServiceAddKeyDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceAddKeyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}