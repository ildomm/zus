// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ildomm/zus/models"
)

// GetHashOKCode is the HTTP code returned for type GetHashOK
const GetHashOKCode int = 200

/*GetHashOK Successful operation

swagger:response getHashOK
*/
type GetHashOK struct {

	/*
	  In: Body
	*/
	Payload *models.TokenInfo `json:"body,omitempty"`
}

// NewGetHashOK creates GetHashOK with default headers values
func NewGetHashOK() *GetHashOK {

	return &GetHashOK{}
}

// WithPayload adds the payload to the get hash o k response
func (o *GetHashOK) WithPayload(payload *models.TokenInfo) *GetHashOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hash o k response
func (o *GetHashOK) SetPayload(payload *models.TokenInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHashOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetHashDefault error

swagger:response getHashDefault
*/
type GetHashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.GeneralError `json:"body,omitempty"`
}

// NewGetHashDefault creates GetHashDefault with default headers values
func NewGetHashDefault(code int) *GetHashDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get hash default response
func (o *GetHashDefault) WithStatusCode(code int) *GetHashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get hash default response
func (o *GetHashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get hash default response
func (o *GetHashDefault) WithPayload(payload *models.GeneralError) *GetHashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hash default response
func (o *GetHashDefault) SetPayload(payload *models.GeneralError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
