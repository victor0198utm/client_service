package models

type ClientOrderReq struct {
	Client_id int        `json:"client_id"`
	Orders    []OrderReq `json:"orders"`
}
