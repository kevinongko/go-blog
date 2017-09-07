package controller

import (
	"net/http"

	"github.com/kevinongko/go-blog/app"
	"github.com/kevinongko/go-blog/model"
)

// HomeIndex to show homepage
func HomeIndex(res http.ResponseWriter, req *http.Request) {
	var article model.Article
	app.DB.First(&article, 1)
	app.Render.HTML(res, http.StatusOK, "home", article)
}
