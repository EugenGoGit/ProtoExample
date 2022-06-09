package key_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// KeyServiceGetPageHandlerFunc turns a function with the right signature into a key service get page handler
type KeyServiceGetPageHandlerFunc func(KeyServiceGetPageParams) middleware.Responder

// Handle executing the request and returning a response
func (fn KeyServiceGetPageHandlerFunc) Handle(params KeyServiceGetPageParams) middleware.Responder {
	return fn(params)
}

// KeyServiceGetPageHandler interface for that can handle valid key service get page params
type KeyServiceGetPageHandler interface {
	Handle(KeyServiceGetPageParams) middleware.Responder
}

// NewKeyServiceGetPage creates a new http.Handler for the key service get page operation
func NewKeyServiceGetPage(ctx *middleware.Context, handler KeyServiceGetPageHandler) *KeyServiceGetPage {
	return &KeyServiceGetPage{Context: ctx, Handler: handler}
}

/*KeyServiceGetPage swagger:route GET /v1/keys KeyService keyServiceGetPage

Get key list (with explicit page number)

*/
type KeyServiceGetPage struct {
	Context *middleware.Context
	Handler KeyServiceGetPageHandler
}

func (o *KeyServiceGetPage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewKeyServiceGetPageParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
