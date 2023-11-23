package web

import "time"

type TransactionCreateRequest struct {
	Name        string    `validate:"required,max=100,min=1" json:"name"`
	Description string    `json:"description"`
	Nominal     float64   `json:"nominal"`
	Date        time.Time `json:"date"`
	CategoryID  int       `json:"category_id"`
}
