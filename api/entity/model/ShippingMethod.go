package model

type ShippingMethod struct {
	ID				uint32	`gorm:"primary_key;auto_increment" json:"id"`
	ShippingMethod	string	`gorm:"size:20;not null" json:"shipping_method"`
}
