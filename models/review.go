package models

type Review struct {
	Restaurant_id int `json:"restaurant_id"`
	Stars         int `json:"stars"`
}
