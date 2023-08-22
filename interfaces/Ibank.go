package interfaces

import (
	"bankapp/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface{
	CreateCustomer(*models.Customer)(*mongo.InsertOneResult,error)
	FetchCustomer(id primitive.ObjectID)([]*models.Customer,error)
	UpdateCustomer(initialValue int, newValue int)(*mongo.UpdateResult,error)
	DeleteCustomer(initialValue int)(error)
}