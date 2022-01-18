package usercase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"test/gen/models"
	"test/gen/restapi"
	user2 "test/gen/restapi/operations/user"
)

type UserUsercaseImpl struct {
}

func (u UserUsercaseImpl) GetUsersUserID(ctx *gin.Context, params *user2.GetUsersUserIDParams) *restapi.GetUsersUserIDResponse {
	g := &restapi.GetUsersUserIDOK{
		Model: models.User{
			CreateDate:    strfmt.Date{},
			DateOfBirth:   nil,
			Email:         nil,
			EmailVerified: nil,
			FirstName:     nil,
			ID:            0,
			LastName:      nil,
		},
	}
	return g.ToResponse()
}

func (u UserUsercaseImpl) PostUser(ctx *gin.Context, params *user2.PostUserParams) *restapi.PostUserResponse {
	//TODO implement me
	panic("implement me")
}

func NewUserUsercaseImpl() *UserUsercaseImpl {
	return &UserUsercaseImpl{}
}
