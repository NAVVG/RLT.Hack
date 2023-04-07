package entity

import (
	"time"
)

type Deal struct {
	Id         int
	UserId     int
	Header     string
	Descripton string
	Price      float64
}

type Sale struct {
	Id     int
	UserId int
	DealId int
	Volume float64
	Date   time.Time
}
