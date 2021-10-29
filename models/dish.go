package models

type Dish struct {
	Dish_id          int    `json:"id"`
	Name             string `json:"name"`
	Preparation_time int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	Cooking_aparatus string `json:"cooking-apparatus"`
}
