// Code generated by gin-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"test/internal/pkg/app"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPetByIDEndpoint executes the core logic of the related
// route endpoint.
func GetPetByIDEndpoint(handler func(ctx *gin.Context, params *GetPetByIDParams) *app.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewGetPetByIDParams()
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := app.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}

			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)
		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewGetPetByIDParams creates a new GetPetByIDParams object
// with the default values initialized.
func NewGetPetByIDParams() *GetPetByIDParams {
	var ()
	return &GetPetByIDParams{}
}

// GetPetByIDParams contains all the bound params for the get pet by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getPetById
type GetPetByIDParams struct {

	/*ID of pet to return
	  Required: true
	  In: path
	*/
	PetID int64
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetPetByIDParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	rPetID := []string{ctx.Param("petId")}
	if err := o.bindPetID(rPetID, true, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPetByIDParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("petId", "path", "int64", raw)
	}
	o.PetID = value

	return nil
}

// vim: ft=go
