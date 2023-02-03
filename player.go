package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type JsonPlayer struct {
	Get        string `json:"get"`
	Parameters struct {
		Search string `json:"search"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Response []struct {
		ID        int    `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Birth     struct {
			Date    string `json:"date"`
			Country string `json:"country"`
		} `json:"birth"`
		Nba struct {
			Start int `json:"start"`
			Pro   int `json:"pro"`
		} `json:"nba"`
		Height struct {
			Feets  string `json:"feets"`
			Inches string `json:"inches"`
			Meters string `json:"meters"`
		} `json:"height"`
		Weight struct {
			Pounds    string `json:"pounds"`
			Kilograms string `json:"kilograms"`
		} `json:"weight"`
		College     string `json:"college"`
		Affiliation string `json:"affiliation"`
		Leagues     struct {
			Standard struct {
				Jersey int    `json:"jersey"`
				Active bool   `json:"active"`
				Pos    string `json:"pos"`
			} `json:"standard"`
		} `json:"leagues"`
	} `json:"response"`
}
//var client *http.Client

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func main(){


	var player string
	var firstName string
	fmt.Println("Prenom du joueur :")
	fmt.Scan(&firstName)
	fmt.Println("Nom du joueur :")
	fmt.Scan(&player)
	fmt.Println(firstName, player) 

	url := "https://api-nba-v1.p.rapidapi.com/players?search=" + player

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "de18451645msh049760fce527e8fp1e555ejsn4eb253213e47")
	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var result JsonPlayer
    if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
        fmt.Println("Can not unmarshal JSON")
    }

	//fmt.Println(PrettyPrint(result))
	var playerId int
	var playerFirstName string
	var playerLastName string

	for _, rec := range result.Response{		
		v := strings.EqualFold(firstName, rec.Firstname)
		if v != false {
			playerFirstName = rec.Firstname
			playerId = rec.ID
			playerLastName = rec.Lastname
		}
	}
	fmt.Println(playerId, playerFirstName, playerLastName)

	
	//recherche les stats dans l'api

	//affiche les stats

	//ecoute le chat

	//affiche le pseudo du viewer 

	//affiche la photo du joueur
}