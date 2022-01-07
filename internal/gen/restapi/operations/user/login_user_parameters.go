// Code generated by gin-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"test/internal/pkg/app"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// LoginUserEndpoint executes the core logic of the related
// route endpoint.
func LoginUserEndpoint(handler func(ctx *gin.Context, params *LoginUserParams) *app.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// generate params from request
		params := NewLoginUserParams()
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

// NewLoginUserParams creates a new LoginUserParams object
// with the default values initialized.
func NewLoginUserParams() *LoginUserParams {
	var ()
	return &LoginUserParams{}
}

// LoginUserParams contains all the bound params for the login user operation
// typically these are obtained from a http.Request
//
// swagger:parameters loginUser
type LoginUserParams struct {

	/*The password for login in clear text
	  Required: true
	  In: query
	*/
	Password string
	/*The user name for login
	  Required: true
	  In: query
	*/
	Username string
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LoginUserParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	qs := runtime.Values(ctx.Request.URL.Query())

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, formats); err != nil {
		res = append(res, err)
	}

	qUsername, qhkUsername, _ := qs.GetOK("username")
	if err := o.bindUsername(qUsername, qhkUsername, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LoginUserParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("password", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("password", "query", raw); err != nil {
		return err
	}

	o.Password = raw

	return nil
}

func (o *LoginUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("username", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("username", "query", raw); err != nil {
		return err
	}

	o.Username = raw

	return nil
}

// vim: ft=go