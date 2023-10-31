package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Name   string
	Side   bool
	Price  int64
	Amount float64
}
