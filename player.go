package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	parser "twitchbot/nbabot/pkg/APINBA/Parser"
	readmsg "twitchbot/nbabot/pkg/APINBA/ReadMsg"
	"twitchbot/nbabot/pkg/APINBA/connect"
	writemessages "twitchbot/nbabot/pkg/APINBA/writeMessages"
	"twitchbot/nbabot/pkg/config"
	"twitchbot/nbabot/pkg/models"
	"twitchbot/nbabot/pkg/routes"

	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type JsonStatsMVP struct {
	Get        string `json:"get"`
	Parameters struct {
		Game   string `json:"game"`
		ID     string `json:"id"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors   []interface{} `json:"errors"`
	Results  int           `json:"results"`
	Response []struct {
		Player struct {
			ID        int    `json:"id"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		} `json:"player"`
		Team struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Nickname string `json:"nickname"`
			Code     string `json:"code"`
			Logo     string `json:"logo"`
		} `json:"team"`
		Game struct {
			ID int `json:"id"`
		} `json:"game"`
		Points    int         `json:"points"`
		Pos       string      `json:"pos"`
		Min       string      `json:"min"`
		Fgm       int         `json:"fgm"`
		Fga       int         `json:"fga"`
		Fgp       string      `json:"fgp"`
		Ftm       int         `json:"ftm"`
		Fta       int         `json:"fta"`
		Ftp       string      `json:"ftp"`
		Tpm       int         `json:"tpm"`
		Tpa       int         `json:"tpa"`
		Tpp       string      `json:"tpp"`
		OffReb    int         `json:"offReb"`
		DefReb    int         `json:"defReb"`
		TotReb    int         `json:"totReb"`
		Assists   int         `json:"assists"`
		PFouls    int         `json:"pFouls"`
		Steals    int         `json:"steals"`
		Turnovers int         `json:"turnovers"`
		Blocks    int         `json:"blocks"`
		PlusMinus string      `json:"plusMinus"`
		Comment   interface{} `json:"comment"`
	} `json:"response"`
}


func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}
type MyStats struct{
	Points int
	Min string
	TotReb int
	Steals int
	Pfouls int
	Assist int
}

type Stats struct{
	Items []models.MyStat
}

func (stats *Stats) AddItem(item models.MyStat) []models.MyStat{
	stats.Items = append(stats.Items, item )
	return stats.Items
}

const api_url string = "api-nba-v1.p.rapidapi.com"

func main(){
	config.Connect()

	godotenv.Load()
	// Definir le joueur à rechercher
	var player string
	var firstName string
	fmt.Println("Prenom du joueur :")
	fmt.Scan(&firstName)
	fmt.Println("Nom du joueur :")
	fmt.Scan(&player)

	//recherche du joueur dans l'api pour definir son Id

	url := "https://api-nba-v1.p.rapidapi.com/players?search=" + player

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", api_url)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

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
	t := time.Now()
	//dt :=  time.Now().Local().Format("2006-01-02")
	

	url2 := "https://api-nba-v1.p.rapidapi.com/games?date=" + "2023-02-17"

	req2, _ := http.NewRequest("GET", url2, nil)

	req2.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req2.Header.Add("X-RapidAPI-Host", api_url)

	res2, _ := http.DefaultClient.Do(req2)

	defer res2.Body.Close()
	body2, _ := ioutil.ReadAll(res2.Body)

	
	var result2 JsonMatchOfTheDay
    if err2 := json.Unmarshal(body2, &result2); err2 != nil { 
        fmt.Println("Can not unmarshal JSON")
    }

	for _, rec2 := range result2.Response{
		fmt.Println("L'Id du match :", rec2.ID)		
		fmt.Println("Visiteur :" , rec2.Teams.Visitors.Name, "score :", rec2.Scores.Visitors.Points)
		fmt.Println("Domicile :" , rec2.Teams.Home.Name, "score :", rec2.Scores.Home.Points)
	}
	

	//recherche les stats du joueurs sur le match en question

	var gameId string
	fmt.Println("choisissez l'Id du match : ")
	fmt.Scan(&gameId)

	PlayerId := strconv.Itoa(playerId)
	y := t.Year()-1
	year := strconv.Itoa(y)
	
	url3 := "https://api-nba-v1.p.rapidapi.com/players/statistics?id=" + PlayerId + "&game=" + gameId + "&season=" + year

	req3, _ := http.NewRequest("GET", url3, nil)

	req3.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req3.Header.Add("X-RapidAPI-Host", api_url)

	res3, _ := http.DefaultClient.Do(req3)

	defer res3.Body.Close()
	body3, _ := ioutil.ReadAll(res3.Body)

	fmt.Println(res3)
	fmt.Println(string(body3))

	var result3 JsonStatsMVP
    if err3 := json.Unmarshal(body3, &result3); err3 != nil { 
        fmt.Println("Can not unmarshal JSON")
    }

	

	for _, rec3 := range result3.Response{
		fmt.Println("Points :", rec3.Points, "pts")		
		fmt.Println("Minutes :" , rec3.Min, "min")
		fmt.Println("Rebonds :" ,rec3.TotReb, "rebonds" )
		fmt.Println("Steals :", rec3.Steals, "recuperations")
		fmt.Println("Fautes :", rec3.PFouls, "fautes")
		fmt.Println("Assists :", rec3.Assists, "assists")
		
		

		stats := models.MyStat{ Firstname: rec3.Player.Firstname, LastName: rec3.Player.Lastname, Points: rec3.Points, Min: rec3.Min, TotReb: rec3.TotReb, Steals: rec3.Steals, Pfouls: rec3.PFouls, Assist: rec3.Assists }
		
		models.NewStats(&stats)
		
		
	}
	
	router := routes.InitializeRouter()
	
	
	
	//affichage web via React
	
	
	//Connection Twitch 

	conn := connect.Connect()
    msgsChan := readmsg.ReadMessages(conn)
    parsedChan := parser.Parse(msgsChan)
    writemessages.WriteMessages(conn, os.Stdin)
	

	//ecoute le chat et affiche le viewer gagnant

	wg:= sync.WaitGroup{}
    wg.Add(2)

	go func(){
		http.ListenAndServe(":9013", router)
		wg.Done()
	}()
	

    go func(){
        for msg := range parsedChan{
            numberOfColons := strings.Count(msg, ":")
            if numberOfColons >= 2 {
				s := strings.Split(msg, ":")
				user := strings.Split(s[1], ",")[0]
				messageData := s[2]
				
				if strings.Contains(messageData, firstName) && strings.Contains(messageData, player){
			        fmt.Fprintf(conn, "PRIVMSG #%s :@%s à gagner !\r\n", os.Getenv("CHANNEL"), user)
				} else {
                    fmt.Println(msg)
                }			
				
				
			}
            
            
        }
		
    }()
    

    wg.Wait()
	
	
	//affiche la photo du joueur
}