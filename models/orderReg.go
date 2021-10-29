package models

type OrderRegistered struct {
	Order_id               int                    `json:"order_id"`
	Is_ready               bool                   `json:"is_ready"`
	Estimated_waiting_time int                    `json:"estimated_waiting_time"`
	Priority               int                    `json:"priority"`
	Max_wait               int                    `json:"max_wait"`
	Created_time           int                    `json:"created_time"`
	Registered_time        int                    `json:"registered_time"`
	Prepared_time          int                    `json:"prepared_time"`
	Cooking_time           int                    `json:"cooking_time"`
	Cooking_details        []Cooking_details_type `json:"cooking_details"`
}
