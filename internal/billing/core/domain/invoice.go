package domain

import "time"

type Invoice struct {
	Id         string
	OrderId    string
	Total      float64
	AmountPaid float64
	CreatedAt  time.Time
}
