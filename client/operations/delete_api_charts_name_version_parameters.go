// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAPIChartsNameVersionParams creates a new DeleteAPIChartsNameVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAPIChartsNameVersionParams() *DeleteAPIChartsNameVersionParams {
	return &DeleteAPIChartsNameVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIChartsNameVersionParamsWithTimeout creates a new DeleteAPIChartsNameVersionParams object
// with the ability to set a timeout on a request.
func NewDeleteAPIChartsNameVersionParamsWithTimeout(timeout time.Duration) *DeleteAPIChartsNameVersionParams {
	return &DeleteAPIChartsNameVersionParams{
		timeout: timeout,
	}
}

// NewDeleteAPIChartsNameVersionParamsWithContext creates a new DeleteAPIChartsNameVersionParams object
// with the ability to set a context for a request.
func NewDeleteAPIChartsNameVersionParamsWithContext(ctx context.Context) *DeleteAPIChartsNameVersionParams {
	return &DeleteAPIChartsNameVersionParams{
		Context: ctx,
	}
}

// NewDeleteAPIChartsNameVersionParamsWithHTTPClient creates a new DeleteAPIChartsNameVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAPIChartsNameVersionParamsWithHTTPClient(client *http.Client) *DeleteAPIChartsNameVersionParams {
	return &DeleteAPIChartsNameVersionParams{
		HTTPClient: client,
	}
}

/*
DeleteAPIChartsNameVersionParams contains all the parameters to send to the API endpoint

	for the delete API charts name version operation.

	Typically these are written to a http.Request.
*/
type DeleteAPIChartsNameVersionParams struct {

	/* Name.

	   chart name
	*/
	Name string

	/* Version.

	   chart version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete API charts name version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIChartsNameVersionParams) WithDefaults() *DeleteAPIChartsNameVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete API charts name version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAPIChartsNameVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) WithTimeout(timeout time.Duration) *DeleteAPIChartsNameVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) WithContext(ctx context.Context) *DeleteAPIChartsNameVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) WithHTTPClient(client *http.Client) *DeleteAPIChartsNameVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) WithName(name string) *DeleteAPIChartsNameVersionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) SetName(name string) {
	o.Name = name
}

// WithVersion adds the version to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) WithVersion(version string) *DeleteAPIChartsNameVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete API charts name version params
func (o *DeleteAPIChartsNameVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIChartsNameVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}