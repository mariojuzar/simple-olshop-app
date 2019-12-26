package model

type DisplayOrderDetail struct {
	OrderId			uint32				`json:"order_id"`
	CustomerName	string				`json:"customer_name"`
	EmployeeName	string				`json:"employee_name"`
	ShippingMethod	string				`json:"shipping_method"`
	ProductList		[]DisplayProduct	`json:"product_list"`
	TotalPayment	float64				`json:"total_payment"`
}

type DisplayProduct struct {
	ProductId 		uint32	`json:"product_id"`
	ProductName		string	`json:"product_name"`
	Quantity		int		`json:"quantity"`
	UnitPrice		float64	`json:"unit_price"`
	Discount		float64	`json:"discount"`
	SubTotalPrice	float64	`json:"sub_total_price"`
}