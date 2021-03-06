/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@semi.technology
 */ // Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateP2pHealthHandlerFunc turns a function with the right signature into a weaviate p2p health handler
type WeaviateP2pHealthHandlerFunc func(WeaviateP2pHealthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateP2pHealthHandlerFunc) Handle(params WeaviateP2pHealthParams) middleware.Responder {
	return fn(params)
}

// WeaviateP2pHealthHandler interface for that can handle valid weaviate p2p health params
type WeaviateP2pHealthHandler interface {
	Handle(WeaviateP2pHealthParams) middleware.Responder
}

// NewWeaviateP2pHealth creates a new http.Handler for the weaviate p2p health operation
func NewWeaviateP2pHealth(ctx *middleware.Context, handler WeaviateP2pHealthHandler) *WeaviateP2pHealth {
	return &WeaviateP2pHealth{Context: ctx, Handler: handler}
}

/*WeaviateP2pHealth swagger:route GET /p2p/health P2P weaviateP2pHealth

Check if a peer is alive.

Check if a peer is alive and healthy.

*/
type WeaviateP2pHealth struct {
	Context *middleware.Context
	Handler WeaviateP2pHealthHandler
}

func (o *WeaviateP2pHealth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateP2pHealthParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
