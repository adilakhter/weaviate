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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewWeaviateSchemaThingsPropertiesUpdateParams creates a new WeaviateSchemaThingsPropertiesUpdateParams object
// with the default values initialized.
func NewWeaviateSchemaThingsPropertiesUpdateParams() *WeaviateSchemaThingsPropertiesUpdateParams {
	var ()
	return &WeaviateSchemaThingsPropertiesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaThingsPropertiesUpdateParamsWithTimeout creates a new WeaviateSchemaThingsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaThingsPropertiesUpdateParamsWithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesUpdateParams {
	var ()
	return &WeaviateSchemaThingsPropertiesUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaThingsPropertiesUpdateParamsWithContext creates a new WeaviateSchemaThingsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaThingsPropertiesUpdateParamsWithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesUpdateParams {
	var ()
	return &WeaviateSchemaThingsPropertiesUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaThingsPropertiesUpdateParamsWithHTTPClient creates a new WeaviateSchemaThingsPropertiesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaThingsPropertiesUpdateParamsWithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesUpdateParams {
	var ()
	return &WeaviateSchemaThingsPropertiesUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaThingsPropertiesUpdateParams contains all the parameters to send to the API endpoint
for the weaviate schema things properties update operation typically these are written to a http.Request
*/
type WeaviateSchemaThingsPropertiesUpdateParams struct {

	/*Body*/
	Body *models.SemanticSchemaClassProperty
	/*ClassName*/
	ClassName string
	/*PropertyName*/
	PropertyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithBody(body *models.SemanticSchemaClassProperty) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetBody(body *models.SemanticSchemaClassProperty) {
	o.Body = body
}

// WithClassName adds the className to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithClassName(className string) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetClassName(className string) {
	o.ClassName = className
}

// WithPropertyName adds the propertyName to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WithPropertyName(propertyName string) *WeaviateSchemaThingsPropertiesUpdateParams {
	o.SetPropertyName(propertyName)
	return o
}

// SetPropertyName adds the propertyName to the weaviate schema things properties update params
func (o *WeaviateSchemaThingsPropertiesUpdateParams) SetPropertyName(propertyName string) {
	o.PropertyName = propertyName
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaThingsPropertiesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param propertyName
	if err := r.SetPathParam("propertyName", o.PropertyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
