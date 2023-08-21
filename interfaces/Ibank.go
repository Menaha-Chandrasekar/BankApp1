package interfaces

import (
	"bankapp/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface{
	CreateCustomer(*models.Customer)(*mongo.InsertOneResult,error)
	FetchCustomer(*models.Customer)()
}