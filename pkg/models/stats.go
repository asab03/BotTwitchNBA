package models

import (
	"fmt"
	"log"
	"twitchbot/nbabot/pkg/config"
)



type MyStat struct{
	Id			int		`json:"id"`
	Firstname 	string 	`json:"firstname"`
	LastName 	string 	`json:"last_name"`
	Points 		int 	`json:"points"`
	Min 		string 	`json:"minutes"`
	TotReb 		int 	`json:"rebonds"`
	Steals 		int 	`json:"steals"`
	Pfouls 		int 	`json:"fautes"`
	Assist 		int 	`json:"assists"`
}

type Stats []MyStat
//migrate table to db

func NewStats(s *MyStat) {
	
	res, err := config.GetDB().Exec("INSERT INTO `my_stats`(  `firstname`, `last_name`, `points`, `min`, `tot_reb`, `steals`, `pfouls`, `assist`) VALUES (?,?,?,?,?,?,?,?)",
	 s.Firstname, s.LastName, s.Points, s.Min, s.TotReb,s.Steals, s.Pfouls, s.Assist)
	
	fmt.Println(res)
	if err!= nil {
	  log.Println(err)
	}
  }


  func GetStatById(id int) *MyStat {
	var stat MyStat
	
	row:= config.GetDB().QueryRow("SELECT * FROM `my_stats` WHERE id = ?", id)
	err := row.Scan(&stat.Id, &stat.Firstname, &stat.LastName, &stat.Points, &stat.Min, &stat.TotReb, &stat.Steals, &stat.Pfouls, &stat.Assist)
	
	if err != nil {
	  log.Println("erreur dans le modele", err)
	}
  
	return &stat
	
  }

  func GetStats() *Stats {
	var stats Stats
  
	rows, err := config.GetDB().Query("SELECT * FROM `my_stats`")
  
	if err != nil {
	  log.Fatal(err)
	}
  
	// Close rows after all readed
	defer rows.Close()
  
	for rows.Next() {
	  var s MyStat
  
	  err := rows.Scan(&s.Id, &s.Firstname, &s.LastName, &s.Points, &s.Min, &s.TotReb, &s.Steals, &s.Pfouls, &s.Assist)
  
	  if err != nil {
		log.Fatal(err)
	  }
  
	  stats = append(stats, s)
	}
  
	return &stats
  }
