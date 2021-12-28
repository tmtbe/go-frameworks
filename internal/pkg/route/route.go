package route

import "github.com/gin-gonic/gin"

type Controller interface {
	GetRoute(r *gin.Engine)
}
