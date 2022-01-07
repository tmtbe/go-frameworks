// Code generated by gin-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"

	"test/internal/gen/models"
	"test/internal/pkg/app"
)

// CreateUsersWithArrayInputEndpoint executes the core logic of the related
// route endpoint.
func CreateUsersWithArrayInputEndpoint(handler func(ctx *gin.Context, params *CreateUsersWithArrayInputParams) *app.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewCreateUsersWithArrayInputParams()
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

// NewCreateUsersWithArrayInputParams creates a new CreateUsersWithArrayInputParams object
// with the default values initialized.
func NewCreateUsersWithArrayInputParams() *CreateUsersWithArrayInputParams {
	var ()
	return &CreateUsersWithArrayInputParams{}
}

// CreateUsersWithArrayInputParams contains all the bound params for the create users with array input operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUsersWithArrayInput
type CreateUsersWithArrayInputParams struct {

	/*List of user object
	  Required: true
	  In: body
	*/
	Body []*models.User
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateUsersWithArrayInputParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body []*models.User
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			for _, io := range o.Body {
				if err := io.Validate(formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// vim: ft=go
