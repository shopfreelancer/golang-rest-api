package main

import (
	"go-rest-api/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	ac := controllers.NewArticleController()
	router.GET("/article/:articleID", ac.ShowArticle)
	router.POST("/article", ac.CreateArticle)
	router.DELETE("/article/:articleID", ac.DeleteArticle)
	http.ListenAndServe(":8080", router)
}
