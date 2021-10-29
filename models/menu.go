package models

type Menu struct {
	Name       string  `json:"name"`
	Menu_items int     `json:"menu_items"`
	Menu       []Dish  `json:"menu"`
	Rating     float64 `json:"rating"`
}
