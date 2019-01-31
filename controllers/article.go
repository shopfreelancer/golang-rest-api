package controllers

import (
	"encoding/json"
	"go-rest-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// ArticleController struct
type ArticleController struct{}

// NewArticleController handles methods for article resource
func NewArticleController() *ArticleController {
	return &ArticleController{}
}

// ShowArticle show one article resource
// curl localhost:8080/article/23
func (ac ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

// CreateArticle - Create a new article resource
// curl -X POST -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article
func (ac ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

// DeleteArticle  - delete one article by id
// curl -X DELETE -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article/1
func (ac ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
