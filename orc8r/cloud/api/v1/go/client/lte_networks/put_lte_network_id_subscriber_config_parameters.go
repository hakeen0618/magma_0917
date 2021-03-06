// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

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

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutLTENetworkIDSubscriberConfigParams creates a new PutLTENetworkIDSubscriberConfigParams object
// with the default values initialized.
func NewPutLTENetworkIDSubscriberConfigParams() *PutLTENetworkIDSubscriberConfigParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutLTENetworkIDSubscriberConfigParamsWithTimeout creates a new PutLTENetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutLTENetworkIDSubscriberConfigParamsWithTimeout(timeout time.Duration) *PutLTENetworkIDSubscriberConfigParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigParams{

		timeout: timeout,
	}
}

// NewPutLTENetworkIDSubscriberConfigParamsWithContext creates a new PutLTENetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutLTENetworkIDSubscriberConfigParamsWithContext(ctx context.Context) *PutLTENetworkIDSubscriberConfigParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigParams{

		Context: ctx,
	}
}

// NewPutLTENetworkIDSubscriberConfigParamsWithHTTPClient creates a new PutLTENetworkIDSubscriberConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutLTENetworkIDSubscriberConfigParamsWithHTTPClient(client *http.Client) *PutLTENetworkIDSubscriberConfigParams {
	var ()
	return &PutLTENetworkIDSubscriberConfigParams{
		HTTPClient: client,
	}
}

/*PutLTENetworkIDSubscriberConfigParams contains all the parameters to send to the API endpoint
for the put LTE network ID subscriber config operation typically these are written to a http.Request
*/
type PutLTENetworkIDSubscriberConfigParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record *models.NetworkSubscriberConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) WithTimeout(timeout time.Duration) *PutLTENetworkIDSubscriberConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) WithContext(ctx context.Context) *PutLTENetworkIDSubscriberConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) WithHTTPClient(client *http.Client) *PutLTENetworkIDSubscriberConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) WithNetworkID(networkID string) *PutLTENetworkIDSubscriberConfigParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) WithRecord(record *models.NetworkSubscriberConfig) *PutLTENetworkIDSubscriberConfigParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put LTE network ID subscriber config params
func (o *PutLTENetworkIDSubscriberConfigParams) SetRecord(record *models.NetworkSubscriberConfig) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutLTENetworkIDSubscriberConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
