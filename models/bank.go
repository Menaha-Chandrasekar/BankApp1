package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct{
	Customer_Id primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Bank_ID int64 `json:"bank_id" bson:"bank_id"`
	Password string `json: "password" bson:"password"`
	Name string `json:"name" bson:"name"`
	Account_Id int64 `json:"account_id" bson:"account_id"`

}