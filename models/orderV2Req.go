package models

type OrderV2Req struct {
	Items        []int `json:"items"`
	Priority     int   `json:"priority"`
	Max_wait     int   `json:"max_wait"`
	Created_time int   `json:"created_time"`
}
