package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/kevinongko/go-blog/database"
	"github.com/kevinongko/go-blog/model"
)

// HomeIndex to show homepage
func HomeIndex(res http.ResponseWriter, req *http.Request) {
	var article model.Article
	database.ORM.First(&article, 1)

	html, err := template.ParseFiles("view/home.html")

	if err != nil {
		fmt.Println(err.Error())
	}

	html.Execute(res, article)
}
