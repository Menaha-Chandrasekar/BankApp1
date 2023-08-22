package controllers

import (
	"bankapp/interfaces"
	"bankapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (f *TransactionController)FetchCustomer(ctx *gin.Context){
	id,_:=primitive.ObjectIDFromHex("64e33245a2fd885e98aa8c17")
    new,err:=f.TransactionService.FetchCustomer(id)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway,gin.H{"status":"fail","message":err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated,gin.H{"status":"success","message":new})
    ctx.String(http.StatusOK,"got the data")
}
func (c * TransactionController)UpdateCustomer(ctx *gin.Context){
    
    new,err:=c.TransactionService.UpdateCustomer(718174533,9876543)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway,gin.H{"status":"fail","message":err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated,gin.H{"status":"success","message":new})
    ctx.String(http.StatusOK,"got the data")

}
func (c * TransactionController)DeleteCustomer(ctx *gin.Context){
    
    err:=c.TransactionService.DeleteCustomer(718174533)
    if err!=nil{
        ctx.JSON(http.StatusBadGateway,gin.H{"status":"fail","message":err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated,gin.H{"status":"success"})
    ctx.String(http.StatusOK,"got the data")

}
