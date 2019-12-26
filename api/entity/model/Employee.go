package model

type Employee struct {
	ID			uint32	`gorm:"primary_key;auto_increment" json:"id"`
	FirstName	string	`gorm:"size:50;not null" json:"first_name"`
	LastName	string	`gorm:"size:50" json:"last_name"`
	Title 		string 	`gorm:"size:50" json:"title"`
	WorkPhone	string	`gorm:"size:30" json:"work_phone"`
}
