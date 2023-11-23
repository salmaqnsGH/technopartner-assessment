package web

import "time"

type TransactionResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Nominal     float64   `json:"nominal"`
	Date        time.Time `json:"date"`
	CategoryID  int       `json:"category_id"`
}
