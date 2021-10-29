package models

type RestaurantDescription struct {
	Restaurant_id int     `json:"restaurant_id"`
	Name          string  `json:"name"`
	Address       string  `json:"address"`
	Menu_items    int     `json:"menu_items"`
	Menu          []Dish  `json:"menu"`
	Rating        float64 `json:"rating"`
}
