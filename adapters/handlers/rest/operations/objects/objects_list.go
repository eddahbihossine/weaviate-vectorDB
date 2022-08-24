//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsListHandlerFunc turns a function with the right signature into a objects list handler
type ObjectsListHandlerFunc func(ObjectsListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ObjectsListHandlerFunc) Handle(params ObjectsListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ObjectsListHandler interface for that can handle valid objects list params
type ObjectsListHandler interface {
	Handle(ObjectsListParams, *models.Principal) middleware.Responder
}

// NewObjectsList creates a new http.Handler for the objects list operation
func NewObjectsList(ctx *middleware.Context, handler ObjectsListHandler) *ObjectsList {
	return &ObjectsList{Context: ctx, Handler: handler}
}

/*
ObjectsList swagger:route GET /objects objects objectsList

Get a list of Objects.

Lists all Objects in reverse order of creation, owned by the user that belongs to the used token.
*/
type ObjectsList struct {
	Context *middleware.Context
	Handler ObjectsListHandler
}

func (o *ObjectsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	Params := NewObjectsListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
