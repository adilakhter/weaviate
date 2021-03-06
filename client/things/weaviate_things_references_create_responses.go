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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateThingsReferencesCreateReader is a Reader for the WeaviateThingsReferencesCreate structure.
type WeaviateThingsReferencesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsReferencesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateThingsReferencesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsReferencesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsReferencesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateThingsReferencesCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateThingsReferencesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsReferencesCreateOK creates a WeaviateThingsReferencesCreateOK with default headers values
func NewWeaviateThingsReferencesCreateOK() *WeaviateThingsReferencesCreateOK {
	return &WeaviateThingsReferencesCreateOK{}
}

/*WeaviateThingsReferencesCreateOK handles this case with default header values.

Successfully added the reference.
*/
type WeaviateThingsReferencesCreateOK struct {
}

func (o *WeaviateThingsReferencesCreateOK) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesCreateOK ", 200)
}

func (o *WeaviateThingsReferencesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesCreateUnauthorized creates a WeaviateThingsReferencesCreateUnauthorized with default headers values
func NewWeaviateThingsReferencesCreateUnauthorized() *WeaviateThingsReferencesCreateUnauthorized {
	return &WeaviateThingsReferencesCreateUnauthorized{}
}

/*WeaviateThingsReferencesCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsReferencesCreateUnauthorized struct {
}

func (o *WeaviateThingsReferencesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesCreateUnauthorized ", 401)
}

func (o *WeaviateThingsReferencesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesCreateForbidden creates a WeaviateThingsReferencesCreateForbidden with default headers values
func NewWeaviateThingsReferencesCreateForbidden() *WeaviateThingsReferencesCreateForbidden {
	return &WeaviateThingsReferencesCreateForbidden{}
}

/*WeaviateThingsReferencesCreateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateThingsReferencesCreateForbidden struct {
}

func (o *WeaviateThingsReferencesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesCreateForbidden ", 403)
}

func (o *WeaviateThingsReferencesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsReferencesCreateUnprocessableEntity creates a WeaviateThingsReferencesCreateUnprocessableEntity with default headers values
func NewWeaviateThingsReferencesCreateUnprocessableEntity() *WeaviateThingsReferencesCreateUnprocessableEntity {
	return &WeaviateThingsReferencesCreateUnprocessableEntity{}
}

/*WeaviateThingsReferencesCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type WeaviateThingsReferencesCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsReferencesCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateThingsReferencesCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateThingsReferencesCreateInternalServerError creates a WeaviateThingsReferencesCreateInternalServerError with default headers values
func NewWeaviateThingsReferencesCreateInternalServerError() *WeaviateThingsReferencesCreateInternalServerError {
	return &WeaviateThingsReferencesCreateInternalServerError{}
}

/*WeaviateThingsReferencesCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateThingsReferencesCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsReferencesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /things/{id}/references/{propertyName}][%d] weaviateThingsReferencesCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateThingsReferencesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
