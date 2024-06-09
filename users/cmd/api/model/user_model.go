package model

type User struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Email   string `json:"email"`
		Password string `json:"password"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
}