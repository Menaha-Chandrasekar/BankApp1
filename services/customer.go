package services

import (
	"bankapp/interfaces"
	"bankapp/models"
	"context"
	
	

	"go.mongodb.org/mongo-driver/bson"
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
func (c *Cust)FetchCustomer(id primitive.ObjectID)([]*models.Customer,error){
	match:=bson.D{{Key:"customer_id",Value: id}}
	result,err:=c.mongoCollection.Find(c.ctx,match)
	if err!=nil{
		return nil, err
	}else{
		var customer_detail[] *models.Customer
		for result.Next(c.ctx){
			detail:=&models.Customer{}
			err:=result.Decode(detail)
			if err!=nil{
				return nil, err
			}
			customer_detail=append(customer_detail, detail)
		}
		return customer_detail,nil
	}
}

func (c *Cust) UpdateCustomer(initialValue int, newValue int)(*mongo.UpdateResult,error){
	filter:=bson.D{{"account_id",initialValue}}
	update:=bson.D{{"$set",bson.D{{"account_id",newValue}}}}
	result,err:=c.mongoCollection.UpdateOne(c.ctx,filter,update)
	if err!=nil{
		
		return nil,err
	}
	return result,nil
}
func (c *Cust) DeleteCustomer(initialValue int)(error){	
	options:=bson.D{{Key: "account_id",Value: initialValue}}
	_,err:=c.mongoCollection.DeleteOne(c.ctx,options)
	if err!=nil{
		return err
	}
	return nil
}