package models

type Loja struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Address    string `json:"address"`
	Categories string `json:"categories"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Name       string `json:"name"`
	Postalcode string `json:"postalcode"`
	Province   string `json:"province"`
	Websites   string `json:"websites"`
}
