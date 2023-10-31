package models

type ScoreCard struct {
	//gorm.Model
	Name     string `gorm:"primary_key"`
	Address  string
	City     string
	Country  string
	Pincode  string
	SATScore float64
	Passed   string
}
