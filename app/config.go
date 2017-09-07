package app

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
)

var (
	// DB is global variable for Database
	DB *gorm.DB

	// Render global variable for view template
	Render *render.Render

	// Router global variable for routing
	Router *mux.Router
)

func InitRender() {
	Render = render.New(render.Options{
		Directory: "view",
	})
}

func InitRouter() {
	Router = mux.NewRouter()
}

func InitDB(dataSource string) {
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		log.Panic(err)
	}
}
