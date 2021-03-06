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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateSchemaActionsPropertiesUpdateReader is a Reader for the WeaviateSchemaActionsPropertiesUpdate structure.
type WeaviateSchemaActionsPropertiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaActionsPropertiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaActionsPropertiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaActionsPropertiesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaActionsPropertiesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaActionsPropertiesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaActionsPropertiesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaActionsPropertiesUpdateOK creates a WeaviateSchemaActionsPropertiesUpdateOK with default headers values
func NewWeaviateSchemaActionsPropertiesUpdateOK() *WeaviateSchemaActionsPropertiesUpdateOK {
	return &WeaviateSchemaActionsPropertiesUpdateOK{}
}

/*WeaviateSchemaActionsPropertiesUpdateOK handles this case with default header values.

Changes applied.
*/
type WeaviateSchemaActionsPropertiesUpdateOK struct {
}

func (o *WeaviateSchemaActionsPropertiesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesUpdateOK ", 200)
}

func (o *WeaviateSchemaActionsPropertiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesUpdateUnauthorized creates a WeaviateSchemaActionsPropertiesUpdateUnauthorized with default headers values
func NewWeaviateSchemaActionsPropertiesUpdateUnauthorized() *WeaviateSchemaActionsPropertiesUpdateUnauthorized {
	return &WeaviateSchemaActionsPropertiesUpdateUnauthorized{}
}

/*WeaviateSchemaActionsPropertiesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaActionsPropertiesUpdateUnauthorized struct {
}

func (o *WeaviateSchemaActionsPropertiesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesUpdateUnauthorized ", 401)
}

func (o *WeaviateSchemaActionsPropertiesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesUpdateForbidden creates a WeaviateSchemaActionsPropertiesUpdateForbidden with default headers values
func NewWeaviateSchemaActionsPropertiesUpdateForbidden() *WeaviateSchemaActionsPropertiesUpdateForbidden {
	return &WeaviateSchemaActionsPropertiesUpdateForbidden{}
}

/*WeaviateSchemaActionsPropertiesUpdateForbidden handles this case with default header values.

Could not find the Action class or property.
*/
type WeaviateSchemaActionsPropertiesUpdateForbidden struct {
}

func (o *WeaviateSchemaActionsPropertiesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesUpdateForbidden ", 403)
}

func (o *WeaviateSchemaActionsPropertiesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaActionsPropertiesUpdateUnprocessableEntity creates a WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity with default headers values
func NewWeaviateSchemaActionsPropertiesUpdateUnprocessableEntity() *WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity {
	return &WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity{}
}

/*WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity handles this case with default header values.

Invalid update.
*/
type WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaActionsPropertiesUpdateInternalServerError creates a WeaviateSchemaActionsPropertiesUpdateInternalServerError with default headers values
func NewWeaviateSchemaActionsPropertiesUpdateInternalServerError() *WeaviateSchemaActionsPropertiesUpdateInternalServerError {
	return &WeaviateSchemaActionsPropertiesUpdateInternalServerError{}
}

/*WeaviateSchemaActionsPropertiesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaActionsPropertiesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaActionsPropertiesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schema/actions/{className}/properties/{propertyName}][%d] weaviateSchemaActionsPropertiesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaActionsPropertiesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
