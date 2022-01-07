package usercase

import (
	"github.com/gin-gonic/gin"
	"test/gen/models"
	"test/gen/restapi/operations/pet"
	"test/internal/pkg/app"
)

type PetUsercaseImpl struct {
}

func NewPetUsercaseImpl() *PetUsercaseImpl {
	return &PetUsercaseImpl{}
}

func (p PetUsercaseImpl) GetPets(ctx *gin.Context, params *pet.GetPetsParams) *app.Response {
	return &app.Response{
		Code: 200,
		Body: models.Pet{
			Age:  123,
			ID:   0,
			Name: "",
			Sex:  false,
		},
	}
}
