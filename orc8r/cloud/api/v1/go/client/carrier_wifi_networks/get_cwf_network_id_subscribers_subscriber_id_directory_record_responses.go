// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordReader is a Reader for the GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecord structure.
type GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK creates a GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK with default headers values
func NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK() *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK {
	return &GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK{}
}

/*GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK handles this case with default header values.

The directory record of a subscriber
*/
type GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK struct {
	Payload *models.CwfSubscriberDirectoryRecord
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/subscribers/{subscriber_id}/directory_record][%d] getCwfNetworkIdSubscribersSubscriberIdDirectoryRecordOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK) GetPayload() *models.CwfSubscriberDirectoryRecord {
	return o.Payload
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CwfSubscriberDirectoryRecord)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault creates a GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault with default headers values
func NewGetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault(code int) *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault {
	return &GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID subscribers subscriber ID directory record default response
func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/subscribers/{subscriber_id}/directory_record][%d] GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecord default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDSubscribersSubscriberIDDirectoryRecordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
