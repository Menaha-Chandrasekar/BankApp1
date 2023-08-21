package services

import (
	"bankapp/interfaces"
	"bankapp/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cust struct{
	mongoCollection *mongo.Collection
	ctx context.Context
}

func InitCustomer (collection *mongo.Collection, ctx context.Context)interfaces.Icustomer{
	return &Cust{collection,ctx}
}
func (c* Cust)CreateCustomer(user *models.Customer)(*mongo.InsertOneResult,error){

	user.Customer_Id = primitive.NewObjectID()
	res,err := c.mongoCollection.InsertOne(c.ctx,&user)
	 if err!= nil{
		return nil,err
	 }
	 return res, nil
}
func (c *Cust)FetchCustomer(user *models.Customer)(){
	
}