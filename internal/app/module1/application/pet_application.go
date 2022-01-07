package application

import (
	"github.com/gin-gonic/gin"
	pet2 "test/internal/gen/restapi/operations/pet"
	"test/internal/pkg/app"
)

type PetApplicationImpl struct {
}

func NewPetApplicationImpl() *PetApplicationImpl {
	return &PetApplicationImpl{}
}

func (p PetApplicationImpl) AddPet(ctx *gin.Context, params *pet2.AddPetParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) DeletePet(ctx *gin.Context, params *pet2.DeletePetParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) FindPetsByStatus(ctx *gin.Context, params *pet2.FindPetsByStatusParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) FindPetsByTags(ctx *gin.Context, params *pet2.FindPetsByTagsParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) GetPetByID(ctx *gin.Context, params *pet2.GetPetByIDParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) UpdatePet(ctx *gin.Context, params *pet2.UpdatePetParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) UpdatePetWithForm(ctx *gin.Context, params *pet2.UpdatePetWithFormParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (p PetApplicationImpl) UploadFile(ctx *gin.Context, params *pet2.UploadFileParams) *app.Response {
	//TODO implement me
	panic("implement me")
}
