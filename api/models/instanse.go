package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Inst struct {
	// gorm.Model
	Customer string `json:"lic.customer"`
	// Expiration      time.Time `json:"lic.expiration"`
	// Features        string    `json:"lic.features"`
	// Max_active_user int       `json:"lic.max-active-users"`
	// // CreatedAt       time.Time
	// // UpdatedAt       time.Time
}

type Instance struct {
	gorm.Model
	Customer        string    `json:"lic.customer" gorm:"primary_key"`
	Expiration      time.Time `json:"lic.expiration"`
	Features        string    `json:"lic.features"`
	Max_active_user int       `json:"lic.max-active-users"`
	// CreatedAt       time.Time
	// UpdatedAt       time.Time
}

// type Instance struct {
// 	gorm.Model
// 	InstanceLicense string    `gorm:"primary_key"`
// 	Licenses        []License `gorm:"foreignKey:Customer;references:InstanceLicense"`
// }

// type License struct {
// 	gorm.Model
// 	Customer        string    `json:"lic.customer"`
// 	Expiration      time.Time `json:"lic.expiration"`
// 	Features        string    `json:"lic.features"`
// 	Max_active_user int       `json:"lic.max-active-users"`
// }
type Order struct {
	// gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
	// gorm.Model
	LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}
