// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutNetworksNetworkIDDNSReader is a Reader for the PutNetworksNetworkIDDNS structure.
type PutNetworksNetworkIDDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDDNSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDDNSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDDNSNoContent creates a PutNetworksNetworkIDDNSNoContent with default headers values
func NewPutNetworksNetworkIDDNSNoContent() *PutNetworksNetworkIDDNSNoContent {
	return &PutNetworksNetworkIDDNSNoContent{}
}

/*PutNetworksNetworkIDDNSNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDDNSNoContent struct {
}

func (o *PutNetworksNetworkIDDNSNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/dns][%d] putNetworksNetworkIdDnsNoContent ", 204)
}

func (o *PutNetworksNetworkIDDNSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDDNSDefault creates a PutNetworksNetworkIDDNSDefault with default headers values
func NewPutNetworksNetworkIDDNSDefault(code int) *PutNetworksNetworkIDDNSDefault {
	return &PutNetworksNetworkIDDNSDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDDNSDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDDNSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID DNS default response
func (o *PutNetworksNetworkIDDNSDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDDNSDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/dns][%d] PutNetworksNetworkIDDNS default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDDNSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDDNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
