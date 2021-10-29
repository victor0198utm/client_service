package models

type Menus struct {
	Restaurants      int    `json:"restaurants"`
	Restaurants_data []Menu `json:"restaurants_data"`
}
