package usercase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"test/gen/models"
	user2 "test/gen/restapi/operations/user"
	"test/internal/pkg/app"
)

type UserUsercaseImpl struct {
}

func (u UserUsercaseImpl) GetUsersUserID(ctx *gin.Context, params *user2.GetUsersUserIDParams) *app.Response {
	return &app.Response{
		Code: 200,
		Body: &models.User{
			CreateDate:    strfmt.Date{},
			DateOfBirth:   nil,
			Email:         nil,
			EmailVerified: nil,
			FirstName:     nil,
			ID:            0,
			LastName:      nil,
		},
	}
}

func (u UserUsercaseImpl) PostUser(ctx *gin.Context, params *user2.PostUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func NewUserUsercaseImpl() *UserUsercaseImpl {
	return &UserUsercaseImpl{}
}
