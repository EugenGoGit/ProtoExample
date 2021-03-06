package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"../../server/models"
)

// KeyServiceGetPageOKCode is the HTTP code returned for type KeyServiceGetPageOK
const KeyServiceGetPageOKCode int = 200

/*KeyServiceGetPageOK A successful response.

swagger:response keyServiceGetPageOK
*/
type KeyServiceGetPageOK struct {

	/*
	  In: Body
	*/
	Payload *models.ExampleKeyPageResponse `json:"body,omitempty"`
}

// NewKeyServiceGetPageOK creates KeyServiceGetPageOK with default headers values
func NewKeyServiceGetPageOK() *KeyServiceGetPageOK {
	return &KeyServiceGetPageOK{}
}

// WithPayload adds the payload to the key service get page o k response
func (o *KeyServiceGetPageOK) WithPayload(payload *models.ExampleKeyPageResponse) *KeyServiceGetPageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service get page o k response
func (o *KeyServiceGetPageOK) SetPayload(payload *models.ExampleKeyPageResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceGetPageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*KeyServiceGetPageDefault An unexpected error response.

swagger:response keyServiceGetPageDefault
*/
type KeyServiceGetPageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.RuntimeError `json:"body,omitempty"`
}

// NewKeyServiceGetPageDefault creates KeyServiceGetPageDefault with default headers values
func NewKeyServiceGetPageDefault(code int) *KeyServiceGetPageDefault {
	if code <= 0 {
		code = 500
	}

	return &KeyServiceGetPageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the key service get page default response
func (o *KeyServiceGetPageDefault) WithStatusCode(code int) *KeyServiceGetPageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the key service get page default response
func (o *KeyServiceGetPageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the key service get page default response
func (o *KeyServiceGetPageDefault) WithPayload(payload *models.RuntimeError) *KeyServiceGetPageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the key service get page default response
func (o *KeyServiceGetPageDefault) SetPayload(payload *models.RuntimeError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *KeyServiceGetPageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
