package routes

import (
	controller "twitchbot/nbabot/pkg/controllers"

	"github.com/gorilla/mux"
)

//Route table

func InitializeRouter() *mux.Router {
	
	router := mux.NewRouter()
  
	router.HandleFunc("/stats", controller.GetStats).Methods("GET")
	router.HandleFunc("/stats/{id}", controller.GetStatsById).Methods("GET")
	router.HandleFunc("/stats", controller.CreateStats).Methods("POST")
	
	
	return router
  }