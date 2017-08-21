package router

import (
	"github.com/kevinongko/go-blog/controller"

	"github.com/gorilla/mux"
)

// SetWebRoutes is for page based routes
func SetWebRoutes(router *mux.Router) {
	router.HandleFunc("/", controller.HomeIndex).Methods("GET")
}
