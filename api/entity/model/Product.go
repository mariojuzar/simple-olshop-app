package model

type Product struct {
	ID			uint32	`gorm:"primary_key;auto_increment" json:"id"`
	ProductName	string	`gorm:"size:50;not null" json:"product_name"`
	UnitPrice	float64	`gorm:"not null" json:"unit_price"`
	InStock 	bool	`gorm:"not null" json:"in_stock"`
}
