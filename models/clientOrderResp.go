package models

type ClientOrderResp struct {
	Order_id int         `json:"order_id"`
	Orders   []OrderResp `json:"orders"`
}
