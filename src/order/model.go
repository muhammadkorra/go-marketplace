package order

import (
	"go-marketplace/src/product"
	"time"
)

type Order struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeliveredAt time.Time `json:"delivered_at"`

	OrderItems []OrderItem `json:"items,omitempty" gorm:"foreignKey:OrderID"`
	UserID     uint        `json:"-"`
}

type OrderItem struct {
	ID       uint `gorm:"primaryKey" json:"id"`
	Quantity uint `json:"quantity"`

	Product   product.Product `json:"product" gorm:"foreignKey:ProductID"`
	ProductID uint            `json:"-"`
	OrderID   uint            `json:"-"`
}
