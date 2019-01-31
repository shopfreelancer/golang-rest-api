package main

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/article/:articleID", showArticle)
	router.POST("/article", createArticle)
	http.ListenAndServe(":8080", router)

}

// showArticle show one article resource
// this site is super useful when it comes to creating structs based on json
// https://mholt.github.io/json-to-go/
func showArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	articleID, _ := strconv.Atoi(ps.ByName("articleID"))

	a1 := models.Article{
		ID:    articleID,
		Title: "asdasd",
	}

	j, err := json.Marshal(a1)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	fmt.Printf("%s", j)
}

// Create a new article resource
// curl -X POST -H "Content-Type: application/json" -d '{"id":33,"title":"asdasd"}' localhost:8080/article/
func createArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
