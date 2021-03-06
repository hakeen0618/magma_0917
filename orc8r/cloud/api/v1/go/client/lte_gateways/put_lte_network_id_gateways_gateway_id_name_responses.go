// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDGatewaysGatewayIDNameReader is a Reader for the PutLTENetworkIDGatewaysGatewayIDName structure.
type PutLTENetworkIDGatewaysGatewayIDNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDGatewaysGatewayIDNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDGatewaysGatewayIDNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDGatewaysGatewayIDNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDGatewaysGatewayIDNameNoContent creates a PutLTENetworkIDGatewaysGatewayIDNameNoContent with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDNameNoContent() *PutLTENetworkIDGatewaysGatewayIDNameNoContent {
	return &PutLTENetworkIDGatewaysGatewayIDNameNoContent{}
}

/*PutLTENetworkIDGatewaysGatewayIDNameNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDGatewaysGatewayIDNameNoContent struct {
}

func (o *PutLTENetworkIDGatewaysGatewayIDNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/name][%d] putLteNetworkIdGatewaysGatewayIdNameNoContent ", 204)
}

func (o *PutLTENetworkIDGatewaysGatewayIDNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDGatewaysGatewayIDNameDefault creates a PutLTENetworkIDGatewaysGatewayIDNameDefault with default headers values
func NewPutLTENetworkIDGatewaysGatewayIDNameDefault(code int) *PutLTENetworkIDGatewaysGatewayIDNameDefault {
	return &PutLTENetworkIDGatewaysGatewayIDNameDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDGatewaysGatewayIDNameDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDGatewaysGatewayIDNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID gateways gateway ID name default response
func (o *PutLTENetworkIDGatewaysGatewayIDNameDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDGatewaysGatewayIDNameDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/gateways/{gateway_id}/name][%d] PutLTENetworkIDGatewaysGatewayIDName default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDGatewaysGatewayIDNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDGatewaysGatewayIDNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
