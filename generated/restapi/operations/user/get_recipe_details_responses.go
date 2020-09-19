// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/gangozero/hackzurich2020-be/generated/models"
)

// GetRecipeDetailsOKCode is the HTTP code returned for type GetRecipeDetailsOK
const GetRecipeDetailsOKCode int = 200

/*GetRecipeDetailsOK Recipe details

swagger:response getRecipeDetailsOK
*/
type GetRecipeDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.RecipeDetails `json:"body,omitempty"`
}

// NewGetRecipeDetailsOK creates GetRecipeDetailsOK with default headers values
func NewGetRecipeDetailsOK() *GetRecipeDetailsOK {

	return &GetRecipeDetailsOK{}
}

// WithPayload adds the payload to the get recipe details o k response
func (o *GetRecipeDetailsOK) WithPayload(payload *models.RecipeDetails) *GetRecipeDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipe details o k response
func (o *GetRecipeDetailsOK) SetPayload(payload *models.RecipeDetails) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipeDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRecipeDetailsDefault Error

swagger:response getRecipeDetailsDefault
*/
type GetRecipeDetailsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRecipeDetailsDefault creates GetRecipeDetailsDefault with default headers values
func NewGetRecipeDetailsDefault(code int) *GetRecipeDetailsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRecipeDetailsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get recipe details default response
func (o *GetRecipeDetailsDefault) WithStatusCode(code int) *GetRecipeDetailsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get recipe details default response
func (o *GetRecipeDetailsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get recipe details default response
func (o *GetRecipeDetailsDefault) WithPayload(payload *models.Error) *GetRecipeDetailsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipe details default response
func (o *GetRecipeDetailsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipeDetailsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}