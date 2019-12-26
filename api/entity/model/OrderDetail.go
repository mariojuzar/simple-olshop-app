package model

type OrderDetail struct {
	ID			uint32	`gorm:"primary_key;auto_increment" json:"id"`
	OrderId 	uint32	`gorm:"not null" json:"order_id"`
	Order		*Order	`json:"order, omitempty"`
	ProductId 	uint32	`gorm:"not null" json:"product_id"`
	Product		*Product`json:"product, omitempty"`
	Quantity 	int		`gorm:"not null" json:"quantity"`
	UnitPrice 	float64	`gorm:"not null" json:"unit_price"`
	Discount	float64	`json:"discount"`
}
	