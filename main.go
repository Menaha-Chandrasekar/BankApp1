package main

import (
	"bankapp/config"
	"bankapp/constants"
	"bankapp/controllers"
	"bankapp/routes"
	"bankapp/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
	routes.Default(server)
}

func initApp(mongoClient *mongo.Client){
	ctx = context.TODO()
	profileCollection := mongoClient.Database("banking").Collection("customer")
	profileService := services.InitCustomer(profileCollection, ctx)
	profileController := controllers.InitTransController(profileService)
	routes.CustRoute(server,profileController)
}


func main(){
	server = gin.Default()
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}