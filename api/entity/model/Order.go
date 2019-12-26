package model

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type Order struct {
	ID					uint32		`gorm:"primary_key;auto_increment" json:"id"`
	CustomerId 			uint32		`gorm:"not null" json:"customer_id"`
	Customer 			*Customer	`json:"customer, omitempty"`
	EmployeeId 			uint32		`gorm:"not null" json:"employee_id"`
	Employee			*Employee	`json:"employee, omitempty"`
	OrderDate 			time.Time	`gorm:"not null" json:"order_date"`
	PurchaseOrderNumber	string		`gorm:"size:30" json:"purchase_order_number"`
	ShipDate 			mysql.NullTime	`json:"ship_date"`
	ShippingMethodId	uint32		`gorm:"not null" json:"shipping_method_id"`
	ShippingMethod 		*ShippingMethod	`json:"shipping_method, omitempty"`
	FreightCharge 		float64		`json:"freight_charge"`
	Taxes 				float64		`json:"taxes"`
	PaymentReceived		bool		`json:"payment_received"`
	Comment 			string		`gorm:"size:150" json:"comment"`
}
