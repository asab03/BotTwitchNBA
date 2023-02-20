package config

import (
	"database/sql"
	"log"
)

var (
	db * sql.DB
)

func Connect() {
	
	d, err := sql.Open("mysql", "rootroot:rootroot@tcp(127.0.0.1:6034)/appnba_db?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err.Error())
	}
	db =  d

	createStatsTable()
}

func createStatsTable(){
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS my_stats(id serial AUTO_INCREMENT, firstname varchar(20), last_name varchar(20), points int, min varchar(20), tot_reb int, steals int, pfouls int, assist int)")

	if err != nil {
	  log.Fatal(err)
	}
}

func GetDB() *sql.DB{
	return db
}