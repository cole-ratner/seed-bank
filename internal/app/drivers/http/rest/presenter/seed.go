package presenter

import (
	"seed-bank/entity"
)

//Seed data
type Seed struct {
	ID       	  entity.ID `json:"id"`
	Category      string  	`json:"category"`
	Name          string 	`json:"name"`
	Vendor        string 	`json:"vendor"`
	Quantity      int 		`json:"quantity"`
	UnitOfMeasure string 	`json:"unitOfMeasure"`
}
