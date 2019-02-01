package main

import (
	"context"
	"go-rest-api/controllers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	mongo "github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	godotenv.Load()
	initRoutes()
}

// init the controller and the routing
func initRoutes() {
	ac := controllers.NewArticleController(getMongoDatabase())
	router := httprouter.New()
	router.GET("/articles", ac.Index)
	router.GET("/article/:articleID", ac.ShowArticle)
	router.POST("/article", ac.CreateArticle)
	router.DELETE("/article/:articleID", ac.DeleteArticle)
	http.ListenAndServe(":8080", router)
}

// connect to mongo db or atlas cluster atlas.mongodb.com
// this wasn´t so easy at first as the mgo driver is outdated
// but mongodb are developing an official driver which is currently beta
// and can deal with the srv connection to the atlas cluster
// read the docs https://godoc.org/github.com/mongodb/mongo-go-driver/mongo
func getMongoDatabase() *mongo.Database {
	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI == "" {
		log.Fatal("No database connection given")
	}

	client, err := mongo.NewClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("golang")
	return database
}
