// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateActionsGetParams creates a new WeaviateActionsGetParams object
// with the default values initialized.
func NewWeaviateActionsGetParams() *WeaviateActionsGetParams {
	var ()
	return &WeaviateActionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsGetParamsWithTimeout creates a new WeaviateActionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsGetParamsWithTimeout(timeout time.Duration) *WeaviateActionsGetParams {
	var ()
	return &WeaviateActionsGetParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsGetParamsWithContext creates a new WeaviateActionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsGetParamsWithContext(ctx context.Context) *WeaviateActionsGetParams {
	var ()
	return &WeaviateActionsGetParams{

		Context: ctx,
	}
}

// NewWeaviateActionsGetParamsWithHTTPClient creates a new WeaviateActionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsGetParamsWithHTTPClient(client *http.Client) *WeaviateActionsGetParams {
	var ()
	return &WeaviateActionsGetParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsGetParams contains all the parameters to send to the API endpoint
for the weaviate actions get operation typically these are written to a http.Request
*/
type WeaviateActionsGetParams struct {

	/*ActionID
	  Unique ID of the action.

	*/
	ActionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions get params
func (o *WeaviateActionsGetParams) WithTimeout(timeout time.Duration) *WeaviateActionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions get params
func (o *WeaviateActionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions get params
func (o *WeaviateActionsGetParams) WithContext(ctx context.Context) *WeaviateActionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions get params
func (o *WeaviateActionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions get params
func (o *WeaviateActionsGetParams) WithHTTPClient(client *http.Client) *WeaviateActionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions get params
func (o *WeaviateActionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the weaviate actions get params
func (o *WeaviateActionsGetParams) WithActionID(actionID strfmt.UUID) *WeaviateActionsGetParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the weaviate actions get params
func (o *WeaviateActionsGetParams) SetActionID(actionID strfmt.UUID) {
	o.ActionID = actionID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param actionId
	if err := r.SetPathParam("actionId", o.ActionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}