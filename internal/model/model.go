package model

type OrderProduct struct {
	ProductId      int64            `json:"product_id"`
	ProductCount   int64            `json:"product_count"`
	ProductType    string           `json:"product_type"`
	OrderId        int64            `json:"order_id"`
	SecondaryRacks []*SecondaryRack `json:"secondary_racks"`
}

type Rack struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type MainRack struct {
	Title         string          `json:"title"`
	OrderProducts []*OrderProduct `json:"order_products"`
}

type SecondaryRack struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}
