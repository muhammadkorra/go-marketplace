package user

import "go-marketplace/src/address"

type User struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Addresses []address.Address `json:"addresses,omitempty" gorm:"foreignKey:UserID;references:ID"`
}
