package models

type Items struct {
	ItemID      uint   `gorm:"not null; primarykey; column:item_id"`
	ItemCode    uint   `json:"itemCode" gorm:"not null; column:item_code"`
	Description string `json:"description" gorm:"not null; type:varchar(300); column:description"`
	Quantity    uint   `json:"quantity" gorm:"not null; column:quantity"`
	OrderID     uint   `gorm:"not null; column:order_id"`
}

type AddOrder struct {
	OrderAt      string  `json:"orderedAt"`
	CustomerName string  `json:"customerName"`
	Items        []Items `json:"items"`
}
