package main

import (
	"fmt"
	"net/http"

	"github.com/kevinongko/go-blog/database"
	"github.com/kevinongko/go-blog/router"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App is the app duh
type App struct {
	Router *mux.Router
}

func main() {
	app := App{}

	app.Router = mux.NewRouter()
	router.SetWebRoutes(app.Router)
	router.SetAPIRoutes(app.Router)

	var err error
	database.DBCon, err = gorm.Open("mysql", "root:secret@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer database.DBCon.Close()

	fmt.Println("webserver start at http://localhost:3000")
	http.ListenAndServe(":3000", app.Router)
}
