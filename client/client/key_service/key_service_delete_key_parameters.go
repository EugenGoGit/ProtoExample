package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewKeyServiceDeleteKeyParams creates a new KeyServiceDeleteKeyParams object
// with the default values initialized.
func NewKeyServiceDeleteKeyParams() *KeyServiceDeleteKeyParams {
	var ()
	return &KeyServiceDeleteKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKeyServiceDeleteKeyParamsWithTimeout creates a new KeyServiceDeleteKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKeyServiceDeleteKeyParamsWithTimeout(timeout time.Duration) *KeyServiceDeleteKeyParams {
	var ()
	return &KeyServiceDeleteKeyParams{

		timeout: timeout,
	}
}

// NewKeyServiceDeleteKeyParamsWithContext creates a new KeyServiceDeleteKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewKeyServiceDeleteKeyParamsWithContext(ctx context.Context) *KeyServiceDeleteKeyParams {
	var ()
	return &KeyServiceDeleteKeyParams{

		Context: ctx,
	}
}

// NewKeyServiceDeleteKeyParamsWithHTTPClient creates a new KeyServiceDeleteKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKeyServiceDeleteKeyParamsWithHTTPClient(client *http.Client) *KeyServiceDeleteKeyParams {
	var ()
	return &KeyServiceDeleteKeyParams{
		HTTPClient: client,
	}
}

/*KeyServiceDeleteKeyParams contains all the parameters to send to the API endpoint
for the key service delete key operation typically these are written to a http.Request
*/
type KeyServiceDeleteKeyParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the key service delete key params
func (o *KeyServiceDeleteKeyParams) WithTimeout(timeout time.Duration) *KeyServiceDeleteKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key service delete key params
func (o *KeyServiceDeleteKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key service delete key params
func (o *KeyServiceDeleteKeyParams) WithContext(ctx context.Context) *KeyServiceDeleteKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key service delete key params
func (o *KeyServiceDeleteKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key service delete key params
func (o *KeyServiceDeleteKeyParams) WithHTTPClient(client *http.Client) *KeyServiceDeleteKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key service delete key params
func (o *KeyServiceDeleteKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the key service delete key params
func (o *KeyServiceDeleteKeyParams) WithID(id int32) *KeyServiceDeleteKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the key service delete key params
func (o *KeyServiceDeleteKeyParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *KeyServiceDeleteKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
