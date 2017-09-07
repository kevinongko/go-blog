package main

import (
	"fmt"
	"net/http"

	"github.com/kevinongko/go-blog/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kevinongko/go-blog/app"
	"github.com/kevinongko/go-blog/model"
)

func main() {
	app.InitRouter()
	app.InitRender()
	app.InitDB("root:secret@/goblog?charset=utf8&parseTime=True&loc=Local")
	app.DB.AutoMigrate(
		&model.Article{},
		&model.User{},
	)

	router.SetWebRoutes()
	router.SetAPIRoutes()

	fmt.Println("webserver start at http://localhost:3000")
	http.ListenAndServe(":3000", app.Router)
}
