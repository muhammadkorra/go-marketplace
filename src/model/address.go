package model

type AddressType string

const (
	Billing  AddressType = "Billing"
	Delivery AddressType = "Delivery"
)

type Address struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Number      string      `json:"number"`
	Building    string      `json:"building"`
	Street      string      `json:"street"`
	Area        string      `json:"area"`
	City        string      `json:"city"`
	AddressType AddressType `json:"address_type"`

	UserID uint `json:"-"`
}
