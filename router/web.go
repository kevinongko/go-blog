package router

import (
	"github.com/kevinongko/go-blog/app"
	"github.com/kevinongko/go-blog/controller"
)

// SetWebRoutes is for page based routes
func SetWebRoutes() {
	app.Router.HandleFunc("/", controller.HomeIndex).Methods("GET")
}
