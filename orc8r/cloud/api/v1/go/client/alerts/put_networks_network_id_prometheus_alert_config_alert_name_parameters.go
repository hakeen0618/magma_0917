// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParams creates a new PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams object
// with the default values initialized.
func NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParams() *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	var ()
	return &PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithTimeout creates a new PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithTimeout(timeout time.Duration) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	var ()
	return &PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams{

		timeout: timeout,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithContext creates a new PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithContext(ctx context.Context) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	var ()
	return &PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams{

		Context: ctx,
	}
}

// NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithHTTPClient creates a new PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutNetworksNetworkIDPrometheusAlertConfigAlertNameParamsWithHTTPClient(client *http.Client) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	var ()
	return &PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams{
		HTTPClient: client,
	}
}

/*PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams contains all the parameters to send to the API endpoint
for the put networks network ID prometheus alert config alert name operation typically these are written to a http.Request
*/
type PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams struct {

	/*AlertConfig
	  Alerting rule that is to be added

	*/
	AlertConfig *models.PromAlertConfig
	/*AlertName
	  Name of alert to be updated

	*/
	AlertName string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithTimeout(timeout time.Duration) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithContext(ctx context.Context) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithHTTPClient(client *http.Client) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertConfig adds the alertConfig to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithAlertConfig(alertConfig *models.PromAlertConfig) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetAlertConfig(alertConfig)
	return o
}

// SetAlertConfig adds the alertConfig to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetAlertConfig(alertConfig *models.PromAlertConfig) {
	o.AlertConfig = alertConfig
}

// WithAlertName adds the alertName to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithAlertName(alertName string) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetAlertName(alertName)
	return o
}

// SetAlertName adds the alertName to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetAlertName(alertName string) {
	o.AlertName = alertName
}

// WithNetworkID adds the networkID to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WithNetworkID(networkID string) *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put networks network ID prometheus alert config alert name params
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PutNetworksNetworkIDPrometheusAlertConfigAlertNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlertConfig != nil {
		if err := r.SetBodyParam(o.AlertConfig); err != nil {
			return err
		}
	}

	// path param alert_name
	if err := r.SetPathParam("alert_name", o.AlertName); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
