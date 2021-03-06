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

// PutNetworksNetworkIDFeaturesReader is a Reader for the PutNetworksNetworkIDFeatures structure.
type PutNetworksNetworkIDFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutNetworksNetworkIDFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutNetworksNetworkIDFeaturesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutNetworksNetworkIDFeaturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutNetworksNetworkIDFeaturesNoContent creates a PutNetworksNetworkIDFeaturesNoContent with default headers values
func NewPutNetworksNetworkIDFeaturesNoContent() *PutNetworksNetworkIDFeaturesNoContent {
	return &PutNetworksNetworkIDFeaturesNoContent{}
}

/*PutNetworksNetworkIDFeaturesNoContent handles this case with default header values.

Success
*/
type PutNetworksNetworkIDFeaturesNoContent struct {
}

func (o *PutNetworksNetworkIDFeaturesNoContent) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/features][%d] putNetworksNetworkIdFeaturesNoContent ", 204)
}

func (o *PutNetworksNetworkIDFeaturesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutNetworksNetworkIDFeaturesDefault creates a PutNetworksNetworkIDFeaturesDefault with default headers values
func NewPutNetworksNetworkIDFeaturesDefault(code int) *PutNetworksNetworkIDFeaturesDefault {
	return &PutNetworksNetworkIDFeaturesDefault{
		_statusCode: code,
	}
}

/*PutNetworksNetworkIDFeaturesDefault handles this case with default header values.

Unexpected Error
*/
type PutNetworksNetworkIDFeaturesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put networks network ID features default response
func (o *PutNetworksNetworkIDFeaturesDefault) Code() int {
	return o._statusCode
}

func (o *PutNetworksNetworkIDFeaturesDefault) Error() string {
	return fmt.Sprintf("[PUT /networks/{network_id}/features][%d] PutNetworksNetworkIDFeatures default  %+v", o._statusCode, o.Payload)
}

func (o *PutNetworksNetworkIDFeaturesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutNetworksNetworkIDFeaturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
