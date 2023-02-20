package models

import (
	"twitchbot/nbabot/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type MyStat struct{
	gorm.Model
	Points int `gorm:"" json:"points"`
	Min string `gorm:"" json:"minutes"`
	TotReb int `gorm:"" json:"rebonds"`
	Steals int `gorm:"" json:"steals"`
	Pfouls int `gorm:"" json:"fautes"`
	Assist int `gorm:"" json:"assists"`
}

//migrate table to db
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&MyStat{})
}

//func create
func (a *MyStat) CreateStats() *MyStat{
	db.NewRecord(a)
	db.Create(&a)
	return a
}

//func get All
func GetStats() []MyStat{
	var Ruches []MyStat
	db.Find(&Ruches)
	return []MyStat{}
}

func GetStatsById(Id int64) (*MyStat, *gorm.DB){
	var getStats MyStat
	db:= db.Where("ID=?", Id).Find(&getStats)
	return &getStats, db
}