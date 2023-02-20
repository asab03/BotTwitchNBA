package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"twitchbot/nbabot/pkg/models"

	"twitchbot/nbabot/pkg/utils"

	"github.com/gorilla/mux"
)

//Stats Controller
var newStats models.MyStat

//Function to get all ruches
func GetStats(w http.ResponseWriter, r *http.Request){
    newStats := models.GetStats()
    res, _ :=json.Marshal(newStats)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
//Function to get one ruche by Id
func GetStatsById(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    Id := vars["Id"]
    ID, err := strconv.ParseInt(Id,0,0)
    if err != nil{
        fmt.Println("error while parsing")
    }
    statsDetails, _ := models.GetStatsById(ID)
    res, _ :=json.Marshal(statsDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
//Function for create ruche
func CreateStat(w http.ResponseWriter, r *http.Request){
    CreateStats := &models.MyStat{}
    utils.ParseBody(r, CreateStats)
    a := CreateStats.CreateStats()
    res, _ :=json.Marshal(a)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}



