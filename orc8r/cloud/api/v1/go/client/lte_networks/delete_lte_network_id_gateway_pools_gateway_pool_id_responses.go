// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// DeleteLTENetworkIDGatewayPoolsGatewayPoolIDReader is a Reader for the DeleteLTENetworkIDGatewayPoolsGatewayPoolID structure.
type DeleteLTENetworkIDGatewayPoolsGatewayPoolIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent creates a DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent with default headers values
func NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent() *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent {
	return &DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent{}
}

/*DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent handles this case with default header values.

Success
*/
type DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent struct {
}

func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/gateway_pools/{gateway_pool_id}][%d] deleteLteNetworkIdGatewayPoolsGatewayPoolIdNoContent ", 204)
}

func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault creates a DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault with default headers values
func NewDeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault(code int) *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault {
	return &DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault{
		_statusCode: code,
	}
}

/*DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete LTE network ID gateway pools gateway pool ID default response
func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /lte/{network_id}/gateway_pools/{gateway_pool_id}][%d] DeleteLTENetworkIDGatewayPoolsGatewayPoolID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteLTENetworkIDGatewayPoolsGatewayPoolIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
