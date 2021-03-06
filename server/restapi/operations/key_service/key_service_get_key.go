package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// KeyServiceGetKeyHandlerFunc turns a function with the right signature into a key service get key handler
type KeyServiceGetKeyHandlerFunc func(KeyServiceGetKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn KeyServiceGetKeyHandlerFunc) Handle(params KeyServiceGetKeyParams) middleware.Responder {
	return fn(params)
}

// KeyServiceGetKeyHandler interface for that can handle valid key service get key params
type KeyServiceGetKeyHandler interface {
	Handle(KeyServiceGetKeyParams) middleware.Responder
}

// NewKeyServiceGetKey creates a new http.Handler for the key service get key operation
func NewKeyServiceGetKey(ctx *middleware.Context, handler KeyServiceGetKeyHandler) *KeyServiceGetKey {
	return &KeyServiceGetKey{Context: ctx, Handler: handler}
}

/*KeyServiceGetKey swagger:route GET /v1/keys/{id} KeyService keyServiceGetKey

Get key by ID

*/
type KeyServiceGetKey struct {
	Context *middleware.Context
	Handler KeyServiceGetKeyHandler
}

func (o *KeyServiceGetKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewKeyServiceGetKeyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
