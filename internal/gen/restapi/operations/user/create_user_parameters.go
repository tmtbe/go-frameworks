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

// CreateUserEndpoint executes the core logic of the related
// route endpoint.
func CreateUserEndpoint(handler func(ctx *gin.Context, params *CreateUserParams) *app.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewCreateUserParams()
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

// NewCreateUserParams creates a new CreateUserParams object
// with the default values initialized.
func NewCreateUserParams() *CreateUserParams {
	var ()
	return &CreateUserParams{}
}

// CreateUserParams contains all the bound params for the create user operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUser
type CreateUserParams struct {

	/*Created user object
	  Required: true
	  In: body
	*/
	Body *models.User
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateUserParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body models.User
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
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
