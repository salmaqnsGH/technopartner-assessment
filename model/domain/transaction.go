package domain

import "time"

type Transaction struct {
	ID          int
	Name        string
	Description string
	Nominal     float64
	Date        time.Time
	CategoryID  int
}
