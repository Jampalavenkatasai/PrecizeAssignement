package models

import "gorm.io/gorm"

type ScoreCard struct {
	gorm.Model
	Name     string
	Address  string
	City     string
	Country  string
	Pincode  string
	SATScore float64
	Passed   bool
}
