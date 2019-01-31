package main

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/article/:articleId", ShowArticle)
	http.ListenAndServe(":8080", router)

}

func ShowArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	articleId := ps.ByName("articleId")

	a1 := models.Article{
		ID:    33,
		Title: "asdasd",
	}

	j, err := json.Marshal(a1)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s", j)
	fmt.Printf("%s", articleId)
}
