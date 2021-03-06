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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateThingsReferencesDeleteNoContentCode is the HTTP code returned for type WeaviateThingsReferencesDeleteNoContent
const WeaviateThingsReferencesDeleteNoContentCode int = 204

/*WeaviateThingsReferencesDeleteNoContent Successfully deleted.

swagger:response weaviateThingsReferencesDeleteNoContent
*/
type WeaviateThingsReferencesDeleteNoContent struct {
}

// NewWeaviateThingsReferencesDeleteNoContent creates WeaviateThingsReferencesDeleteNoContent with default headers values
func NewWeaviateThingsReferencesDeleteNoContent() *WeaviateThingsReferencesDeleteNoContent {

	return &WeaviateThingsReferencesDeleteNoContent{}
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// WeaviateThingsReferencesDeleteUnauthorizedCode is the HTTP code returned for type WeaviateThingsReferencesDeleteUnauthorized
const WeaviateThingsReferencesDeleteUnauthorizedCode int = 401

/*WeaviateThingsReferencesDeleteUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsReferencesDeleteUnauthorized
*/
type WeaviateThingsReferencesDeleteUnauthorized struct {
}

// NewWeaviateThingsReferencesDeleteUnauthorized creates WeaviateThingsReferencesDeleteUnauthorized with default headers values
func NewWeaviateThingsReferencesDeleteUnauthorized() *WeaviateThingsReferencesDeleteUnauthorized {

	return &WeaviateThingsReferencesDeleteUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesDeleteUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateThingsReferencesDeleteForbiddenCode is the HTTP code returned for type WeaviateThingsReferencesDeleteForbidden
const WeaviateThingsReferencesDeleteForbiddenCode int = 403

/*WeaviateThingsReferencesDeleteForbidden Insufficient permissions.

swagger:response weaviateThingsReferencesDeleteForbidden
*/
type WeaviateThingsReferencesDeleteForbidden struct {
}

// NewWeaviateThingsReferencesDeleteForbidden creates WeaviateThingsReferencesDeleteForbidden with default headers values
func NewWeaviateThingsReferencesDeleteForbidden() *WeaviateThingsReferencesDeleteForbidden {

	return &WeaviateThingsReferencesDeleteForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesDeleteForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateThingsReferencesDeleteNotFoundCode is the HTTP code returned for type WeaviateThingsReferencesDeleteNotFound
const WeaviateThingsReferencesDeleteNotFoundCode int = 404

/*WeaviateThingsReferencesDeleteNotFound Successful query result but no resource was found.

swagger:response weaviateThingsReferencesDeleteNotFound
*/
type WeaviateThingsReferencesDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsReferencesDeleteNotFound creates WeaviateThingsReferencesDeleteNotFound with default headers values
func NewWeaviateThingsReferencesDeleteNotFound() *WeaviateThingsReferencesDeleteNotFound {

	return &WeaviateThingsReferencesDeleteNotFound{}
}

// WithPayload adds the payload to the weaviate things references delete not found response
func (o *WeaviateThingsReferencesDeleteNotFound) WithPayload(payload *models.ErrorResponse) *WeaviateThingsReferencesDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things references delete not found response
func (o *WeaviateThingsReferencesDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsReferencesDeleteInternalServerErrorCode is the HTTP code returned for type WeaviateThingsReferencesDeleteInternalServerError
const WeaviateThingsReferencesDeleteInternalServerErrorCode int = 500

/*WeaviateThingsReferencesDeleteInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateThingsReferencesDeleteInternalServerError
*/
type WeaviateThingsReferencesDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateThingsReferencesDeleteInternalServerError creates WeaviateThingsReferencesDeleteInternalServerError with default headers values
func NewWeaviateThingsReferencesDeleteInternalServerError() *WeaviateThingsReferencesDeleteInternalServerError {

	return &WeaviateThingsReferencesDeleteInternalServerError{}
}

// WithPayload adds the payload to the weaviate things references delete internal server error response
func (o *WeaviateThingsReferencesDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateThingsReferencesDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things references delete internal server error response
func (o *WeaviateThingsReferencesDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsReferencesDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
