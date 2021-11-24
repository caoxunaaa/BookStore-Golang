package order

import (
	"WebApi/Svc"
	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(c *gin.Context) {
	Svc.SvcContext.Model.ConfirmInventoryEnough()
}
