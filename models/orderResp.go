package models

type OrderResp struct {
	Restaurant_id          int    `json:"restaurant_id"`
	Restaurant_address     string `json:"restaurant_address"`
	Order_id               int    `json:"order_id"`
	Estimated_waiting_time int    `json:"estimated_waiting_time"`
	Created_time           int    `json:"created_time"`
	Registered_time        int    `json:"registered_time"`
}
