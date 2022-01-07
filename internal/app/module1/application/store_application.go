package application

import (
	"github.com/gin-gonic/gin"
	store2 "test/internal/gen/restapi/operations/store"
	"test/internal/pkg/app"
)

type StoreApplication struct {
}

func NewStoreApplication() *StoreApplication {
	return &StoreApplication{}
}

func (s StoreApplication) DeleteOrder(ctx *gin.Context, params *store2.DeleteOrderParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (s StoreApplication) GetInventory(ctx *gin.Context) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (s StoreApplication) GetOrderByID(ctx *gin.Context, params *store2.GetOrderByIDParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (s StoreApplication) PlaceOrder(ctx *gin.Context, params *store2.PlaceOrderParams) *app.Response {
	//TODO implement me
	panic("implement me")
}
