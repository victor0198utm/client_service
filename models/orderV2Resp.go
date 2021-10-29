package models

type OrderV2Resp struct {
	Restaurant_id          int `json:"restaurant_id"`
	Order_id               int `json:"order_id"`
	Estimated_waiting_time int `json:"estimated_waiting_time"`
	Created_time           int `json:"created_time"`
	Registered_time        int `json:"registered_time"`
}
