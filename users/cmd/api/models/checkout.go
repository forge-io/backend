package models

type Checkout struct {
	ID      string `json:"id"`
	IDCar   string `json:"idCar"`
	IDUser  string `json:"idUser"`
}
