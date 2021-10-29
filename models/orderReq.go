package models

type OrderReq struct {
	Restaurant_id int   `json:"restaurant_id"`
	Items         []int `json:"items"`
	Priority      int   `json:"priority"`
	Max_wait      int   `json:"max_wait"`
	Created_time  int   `json:"created_time"`
}
