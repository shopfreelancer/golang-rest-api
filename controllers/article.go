package controllers

import (
	"context"
	"encoding/json"
	"go-rest-api/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// ArticleController struct
type ArticleController struct {
	database *mongo.Database
}

// ErrorMessage - simple struct for json error message
type ErrorMessage struct {
	Message string `json:"message"`
}

// NewArticleController handles methods for article resource
func NewArticleController(d *mongo.Database) *ArticleController {
	return &ArticleController{d}
}

// Index - list articles
// curl localhost:8080/articles
func (ac ArticleController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collection := ac.database.Collection("articles")
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		msg := ErrorMessage{"No matching articles found"}
		j, _ := json.Marshal(msg)
		w.Write(j)
		return
	}

	// struct for collection of articles
	var articles []models.Article

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		a := models.Article{}
		cur.Decode(&a)
		articles = append(articles, a)
	}

	j, _ := json.Marshal(articles)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

// ShowArticle show one article resource
// curl localhost:8080/article/23
func (ac ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collection := ac.database.Collection("articles")

	articleID := ps.ByName("articleID")
	articleIdHex, err1 := primitive.ObjectIDFromHex(articleID)

	if err1 != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{"Not a valid Object ID"}
		j, _ := json.Marshal(msg)
		w.Write(j)
		return
	}

	a := models.Article{}

	err := collection.FindOne(context.Background(), bson.M{"_id": articleIdHex}).Decode(&a)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		msg := ErrorMessage{"No matching article found"}
		j, _ := json.Marshal(msg)
		w.Write(j)
		return
	}

	j, _ := json.Marshal(a)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// CreateArticle - Create a new article resource
// curl -X POST -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article
func (ac ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	a := models.Article{}

	json.NewDecoder(r.Body).Decode(&a)

	collection := ac.database.Collection("articles")

	// assigning an own ID here
	a.ID = primitive.NewObjectID()

	res, err := collection.InsertOne(context.Background(), a)
	if err != nil {
		log.Println(err)
	}

	// using the internal mongodb id could work but...
	// https://stackoverflow.com/questions/37329246/how-to-convert-string-from-interface-to-string-in-golang
	id := res.InsertedID
	//a.ID = id
	log.Println(id)

	j, err := json.Marshal(a)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// DeleteArticle  - delete one article by id
// curl -X DELETE -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article/1
func (ac ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	collection := ac.database.Collection("articles")

	articleID := ps.ByName("articleID")
	articleIdHex, err1 := primitive.ObjectIDFromHex(articleID)

	if err1 != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		msg := ErrorMessage{"Not a valid Object ID"}
		j, _ := json.Marshal(msg)
		w.Write(j)
		return
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": articleIdHex})

	if err != nil || res.DeletedCount == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		msg := ErrorMessage{"No matching article found"}
		j, _ := json.Marshal(msg)
		w.Write(j)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
