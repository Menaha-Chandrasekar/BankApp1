package routes

import (
	"bankapp/controllers"

	"github.com/gin-gonic/gin"
)

func CustRoute(router *gin.Engine,controller controllers.TransactionController){
	router.POST("/api/",controller.CreateCustomer)
}