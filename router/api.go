package router

import (
	"github.com/kevinongko/go-blog/app"
	"github.com/kevinongko/go-blog/controller"
)

// SetAPIRoutes is for api based routes
func SetAPIRoutes() {
	app.Router.HandleFunc("/api/article", controller.HomeIndex).Methods("GET")
}
