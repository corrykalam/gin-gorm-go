package models

type Orders struct {
	OrderID      uint   `json:"order_id" gorm:"not null; primarykey; column:order_id"`
	CostumerName string `json:"costumer_name" gorm:"not null; type:varchar(300); column:customer_name"`
	OrderAt      string `json:"order_at" gorm:"not null; column:order_at"`
}
