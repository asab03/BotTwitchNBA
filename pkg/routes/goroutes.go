package routes

import (
	controller "twitchbot/nbabot/pkg/controllers"

	"github.com/gorilla/mux"
)

//Route table

var RegisterGoRoutes = func(router *mux.Router){
	// route ruches
	router.HandleFunc("/stats", controller.GetStats).Methods("GET")
	
}