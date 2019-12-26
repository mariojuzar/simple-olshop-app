package model

type Customer struct {
	ID					uint32	`gorm:"primary_key;auto_increment" json:"id"`
	CompanyName			string	`gorm:"size:50" json:"company_name"`
	FirstName			string	`gorm:"size:30;not null" json:"first_name"`
	LastName			string	`gorm:"size:50" json:"last_name"`
	BillingAddress		string	`gorm:"size:255" json:"billing_address"`
	City 				string	`gorm:"size:50" json:"city"`
	StateOrProvince		string	`gorm:"size:20" json:"state_or_province"`
	ZipCode 			string	`gorm:"size:20" json:"zip_code"`
	Email 				string	`gorm:"size:75" json:"email"`
	CompanyWebsite		string	`gorm:"size:200" json:"company_website"`
	PhoneNumber			string	`gorm:"size:30" json:"phone_number"`
	FaxNumber			string	`gorm:"size:30" json:"fax_number"`
	ShipAddress			string	`gorm:"size:255" json:"ship_address"`
	ShipCity			string 	`gorm:"size:50" json:"ship_city"`
	ShipStateOrProvince	string	`gorm:"size:50" json:"ship_state_or_province"`
	ShipZipCode			string	`gorm:"size:50" json:"ship_zip_code"`
	ShipPhoneNumber		string	`gorm:"size:45" json:"ship_phone_number"`
}