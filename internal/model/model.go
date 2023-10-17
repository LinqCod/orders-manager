package model

type OrderInfo struct {
	OrderId     int64  `json:"order_id"`
	ProductId   int64  `json:"product_id"`
	Count       int64  `json:"count"`
	ProductType string `json:"product_type"`
	RackId      int64  `json:"rack_id"`
	RackType    string `json:"rack_type"`
	RackTitle   string `json:"rack_title"`
}
