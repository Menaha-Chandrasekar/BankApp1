package controllers

import (
	"bankapp/interfaces"
	"bankapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{
	TransactionService interfaces.Icustomer

}

func InitTransController(transactionService interfaces.Icustomer)TransactionController{
	return TransactionController{transactionService}
}

func (t* TransactionController)CreateCustomer(ctx *gin.Context){
	var trans *models.Customer
	if err:= ctx.ShouldBindJSON(&trans);err!= nil{
		ctx.JSON(http.StatusBadRequest, err.Error())
		return 
	}
	newtrans,err:= t.TransactionService.CreateCustomer(trans)
	if err!= nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})
}