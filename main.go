package main

import (
	"encoding/json"
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
	router.DELETE("/article/:articleID", deleteArticle)
	http.ListenAndServe(":8080", router)

}

// showArticle show one article resource
// this site is super useful when it comes to creating structs based on json
// https://mholt.github.io/json-to-go/
func showArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	articleID, _ := strconv.Atoi(ps.ByName("articleID"))

	a := models.Article{
		ID:    articleID,
		Title: "asdasd",
	}

	j, err := json.Marshal(a)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

	//fmt.Printf("%s", j)
}

// Create a new article resource
// curl -X POST -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article
func createArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	a := models.Article{}

	json.NewDecoder(r.Body).Decode(&a)

	a.ID = 999

	j, err := json.Marshal(a)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// delete article
// // curl -X DELETE -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article/1
func deleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
