package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

type JsonMatchOfTheDay struct {
	Get        string `json:"get"`
	Parameters struct {
		Date string `json:"date"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Response []struct {
		ID     int    `json:"id"`
		League string `json:"league"`
		Season int    `json:"season"`
		Date   struct {
			Start    time.Time   `json:"start"`
			End      interface{} `json:"end"`
			Duration interface{} `json:"duration"`
		} `json:"date"`
		Stage  int `json:"stage"`
		Status struct {
			Clock    interface{} `json:"clock"`
			Halftime bool        `json:"halftime"`
			Short    int         `json:"short"`
			Long     string      `json:"long"`
		} `json:"status"`
		Periods struct {
			Current     int  `json:"current"`
			Total       int  `json:"total"`
			EndOfPeriod bool `json:"endOfPeriod"`
		} `json:"periods"`
		Arena struct {
			Name    interface{} `json:"name"`
			City    interface{} `json:"city"`
			State   interface{} `json:"state"`
			Country interface{} `json:"country"`
		} `json:"arena"`
		Teams struct {
			Visitors struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"visitors"`
			Home struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				Nickname string `json:"nickname"`
				Code     string `json:"code"`
				Logo     string `json:"logo"`
			} `json:"home"`
		} `json:"teams"`
		Scores struct {
			Visitors struct {
				Win    int `json:"win"`
				Loss   int `json:"loss"`
				Series struct {
					Win  int `json:"win"`
					Loss int `json:"loss"`
				} `json:"series"`
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"visitors"`
			Home struct {
				Win    int `json:"win"`
				Loss   int `json:"loss"`
				Series struct {
					Win  int `json:"win"`
					Loss int `json:"loss"`
				} `json:"series"`
				Linescore []string `json:"linescore"`
				Points    int      `json:"points"`
			} `json:"home"`
		} `json:"scores"`
		Officials   []interface{} `json:"officials"`
		TimesTied   interface{}   `json:"timesTied"`
		LeadChanges interface{}   `json:"leadChanges"`
		Nugget      interface{}   `json:"nugget"`
	} `json:"response"`
}


func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func main(){

	// Definir le joueur à rechercher
	var player string
	var firstName string
	fmt.Println("Prenom du joueur :")
	fmt.Scan(&firstName)
	fmt.Println("Nom du joueur :")
	fmt.Scan(&player)
	fmt.Println(firstName, player) 

	//recherche du joueur dans l'api pour definir son Id

	url := "https://api-nba-v1.p.rapidapi.com/players?search=" + player

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "de18451645msh049760fce527e8fp1e555ejsn4eb253213e47")
	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	var result JsonPlayer
    if err := json.Unmarshal(body, &result); err != nil { 
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

	//recherche du match dans l'api pour obtenir Id
	
	dt :=  time.Now().Local().Format("2006-01-02")
	

	fmt.Println(dt)
	

	url2 := "https://api-nba-v1.p.rapidapi.com/games?date=" + dt

	fmt.Println(url2)

	req2, _ := http.NewRequest("GET", url2, nil)

	req2.Header.Add("X-RapidAPI-Key", "de18451645msh049760fce527e8fp1e555ejsn4eb253213e47")
	req2.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	res2, _ := http.DefaultClient.Do(req2)

	defer res2.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)

	//fmt.Println(res2)
	//fmt.Println(string(body2))

	var result2 JsonMatchOfTheDay
    if err2 := json.Unmarshal(body2, &result2); err2 != nil { 
        fmt.Println("Can not unmarshal JSON")
    }

	for _, rec2 := range result2.Response{
		fmt.Println("L'Id du match :", rec2.ID)		
		fmt.Println("Visiteur :" , rec2.Teams.Visitors.Name, "score :", rec2.Scores.Visitors.Points)
		fmt.Println("Domicile :" , rec2.Teams.Home.Name, "score :", rec2.Scores.Home.Points)
	}

	//recherche les stats dans l'api

	//affiche les stats

	//ecoute le chat

	//affiche le pseudo du viewer 

	//affiche la photo du joueur
}