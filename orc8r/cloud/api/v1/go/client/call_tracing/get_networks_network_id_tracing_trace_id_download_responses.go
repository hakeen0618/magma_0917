// Code generated by go-swagger; DO NOT EDIT.

package call_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetNetworksNetworkIDTracingTraceIDDownloadReader is a Reader for the GetNetworksNetworkIDTracingTraceIDDownload structure.
type GetNetworksNetworkIDTracingTraceIDDownloadReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksNetworkIDTracingTraceIDDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksNetworkIDTracingTraceIDDownloadOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworksNetworkIDTracingTraceIDDownloadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworksNetworkIDTracingTraceIDDownloadOK creates a GetNetworksNetworkIDTracingTraceIDDownloadOK with default headers values
func NewGetNetworksNetworkIDTracingTraceIDDownloadOK(writer io.Writer) *GetNetworksNetworkIDTracingTraceIDDownloadOK {
	return &GetNetworksNetworkIDTracingTraceIDDownloadOK{
		Payload: writer,
	}
}

/*GetNetworksNetworkIDTracingTraceIDDownloadOK handles this case with default header values.

Show tracing status
*/
type GetNetworksNetworkIDTracingTraceIDDownloadOK struct {
	ContentDisposition string

	Payload io.Writer
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadOK) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing/{trace_id}/download][%d] getNetworksNetworkIdTracingTraceIdDownloadOK  %+v", 200, o.Payload)
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksNetworkIDTracingTraceIDDownloadDefault creates a GetNetworksNetworkIDTracingTraceIDDownloadDefault with default headers values
func NewGetNetworksNetworkIDTracingTraceIDDownloadDefault(code int) *GetNetworksNetworkIDTracingTraceIDDownloadDefault {
	return &GetNetworksNetworkIDTracingTraceIDDownloadDefault{
		_statusCode: code,
	}
}

/*GetNetworksNetworkIDTracingTraceIDDownloadDefault handles this case with default header values.

Unexpected Error
*/
type GetNetworksNetworkIDTracingTraceIDDownloadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get networks network ID tracing trace ID download default response
func (o *GetNetworksNetworkIDTracingTraceIDDownloadDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadDefault) Error() string {
	return fmt.Sprintf("[GET /networks/{network_id}/tracing/{trace_id}/download][%d] GetNetworksNetworkIDTracingTraceIDDownload default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworksNetworkIDTracingTraceIDDownloadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
