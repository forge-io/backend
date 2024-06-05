package models

type Product struct {
	ID       string  `json:"id"`
	Model    string  `json:"model"`
	Category string  `json:"category"`
	Year     int     `json:"year"`
	Brand    string  `json:"brand"`
	Km       string  `json:"km"`
	Collor   string  `json:"collor"`
	Motor    string  `json:"motor"`
	Price    float64 `json:"price"`
}