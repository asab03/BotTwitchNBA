package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"twitchbot/nbabot/pkg/models"

	"github.com/gorilla/mux"
)

func StatsIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
  
    json.NewEncoder(w).Encode(models.GetStats())
}

func CreateStats(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
  
    body, err := ioutil.ReadAll(r.Body)
  
    if err != nil {
      log.Fatal(err)
    }
  
    var stat models.MyStat
  
    err = json.Unmarshal(body, &stat)
  
    if err != nil {
      log.Fatal(err)
    }
  
    models.NewStats(&stat)
  
    json.NewEncoder(w).Encode(stat)
}
func GetStats(w http.ResponseWriter, r *http.Request){
    stats := models.GetStats()
    res, _ :=json.Marshal(stats)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func GetStatsById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
  
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
  
    if err != nil {
      log.Fatal(err)
    }
  
    stat := models.GetStatById(id)
  
    json.NewEncoder(w).Encode(stat)
  }