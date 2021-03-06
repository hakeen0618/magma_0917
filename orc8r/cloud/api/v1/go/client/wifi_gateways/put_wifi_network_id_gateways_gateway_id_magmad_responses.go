// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutWifiNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the PutWifiNetworkIDGatewaysGatewayIDMagmad structure.
type PutWifiNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutWifiNetworkIDGatewaysGatewayIDMagmadNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutWifiNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadNoContent creates a PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadNoContent() *PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent {
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent{}
}

/*PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent handles this case with default header values.

Success
*/
type PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent struct {
}

func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/magmad][%d] putWifiNetworkIdGatewaysGatewayIdMagmadNoContent ", 204)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutWifiNetworkIDGatewaysGatewayIDMagmadDefault creates a PutWifiNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewPutWifiNetworkIDGatewaysGatewayIDMagmadDefault(code int) *PutWifiNetworkIDGatewaysGatewayIDMagmadDefault {
	return &PutWifiNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/*PutWifiNetworkIDGatewaysGatewayIDMagmadDefault handles this case with default header values.

Unexpected Error
*/
type PutWifiNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put wifi network ID gateways gateway ID magmad default response
func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[PUT /wifi/{network_id}/gateways/{gateway_id}/magmad][%d] PutWifiNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}

func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutWifiNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
