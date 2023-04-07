package entity

import (
	"time"
)

type Deal struct {
	Id         int
	Header     string
	Descripton string
	Price      float64
}

type Sale struct {
	Id     int
	Volume float64
	Date   time.Time
}
