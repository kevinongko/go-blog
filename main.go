package main

import (
	"fmt"
	"net/http"

	"github.com/kevinongko/go-blog/database"
	"github.com/kevinongko/go-blog/router"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database.Init("root:secret@/goblog?charset=utf8&parseTime=True&loc=Local")

	Router := mux.NewRouter()
	router.SetWebRoutes(Router)
	router.SetAPIRoutes(Router)

	fmt.Println("webserver start at http://localhost:3000")
	http.ListenAndServe(":3000", Router)
}
