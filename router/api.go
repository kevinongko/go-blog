package router

import (
	"github.com/kevinongko/go-blog/controller"

	"github.com/gorilla/mux"
)

// SetAPIRoutes is for api based routes
func SetAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/article", controller.HomeIndex).Methods("GET")
}
