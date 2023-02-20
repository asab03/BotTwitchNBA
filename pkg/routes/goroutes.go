package routes

import (
	controller "twitchbot/nbabot/pkg/controllers"

	"github.com/gorilla/mux"
)

//Route table

var RegisterGoRoutes = func(router *mux.Router){
	// route stats
	router.HandleFunc("/stats", controller.GetStats).Methods("GET")
	router.HandleFunc("/stats/{Id}", controller.GetStats).Methods("GET")
	
}