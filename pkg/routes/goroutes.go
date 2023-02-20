package routes

import (
	controller "twitchbot/nbabot/pkg/controllers"

	"github.com/gorilla/mux"
)

//Route table

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)
  
	router.Methods("GET").Path("/stats").HandlerFunc(controller.GetStats)
	router.Methods("POST").Path("/stats").Name("Create").HandlerFunc(controller.CreateStats)
	router.Methods("GET").Path("/stats/{id}").Name("Show").HandlerFunc(controller.GetStatsById)
	
	return router
  }