package models

import "time"

// DeliveryMethod represents how an order will be fulfilled.
type DeliveryMethod string

const (
	DeliveryPickup        DeliveryMethod = "pickup"
	DeliveryFarmer        DeliveryMethod = "farmer_delivery"
	DeliveryPartnerMethod DeliveryMethod = "partner_delivery"
)

// DeliveryPartner represents a third-party or local delivery person/service.
type DeliveryPartner struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name"`
	Contact   string  `gorm:"type:varchar(20);not null" json:"contact"`
	VehicleNo string  `gorm:"type:varchar(50)" json:"vehicle_no"`
	Rating    float64 `gorm:"type:decimal(2,1);default:0" json:"rating"`
	IsActive  bool    `gorm:"default:true" json:"is_active"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// DeliveryOption defines delivery preferences/capabilities for a shop (farmer).
type DeliveryOption struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ShopID uint `gorm:"not null;index" json:"shop_id"`

	SupportsPickup     bool    `gorm:"default:true" json:"supports_pickup"`
	SupportsDelivery   bool    `gorm:"default:false" json:"supports_delivery"`
	DeliveryRadiusKm   float64 `gorm:"type:decimal(5,2);default:5.0" json:"delivery_radius_km"`
	BaseDeliveryCharge float64 `gorm:"type:decimal(10,2);default:0" json:"base_delivery_charge"`
	PerKmCharge        float64 `gorm:"type:decimal(10,2);default:5" json:"per_km_charge"`
	HasPartnerDelivery bool    `gorm:"default:false" json:"has_partner_delivery"`

	PreferredPartnerID *uint            `json:"preferred_partner_id"`
	Partner            *DeliveryPartner `gorm:"foreignKey:PreferredPartnerID" json:"partner"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// DeliveryRecord captures fulfillment details for a specific order.
type DeliveryRecord struct {
	ID      uint `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID uint `gorm:"uniqueIndex;not null" json:"order_id"`
	ShopID  uint `gorm:"index;not null" json:"shop_id"`
	UserID  uint `gorm:"index;not null" json:"user_id"`

	Method DeliveryMethod `gorm:"type:varchar(50);not null" json:"method"`

	DeliveryAddress string  `gorm:"type:varchar(255)" json:"delivery_address"`
	Latitude        float64 `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude       float64 `gorm:"type:decimal(10,8)" json:"longitude"`

	DeliveryCharge  float64          `gorm:"type:decimal(10,2);default:0" json:"delivery_charge"`
	AssignedPartner *uint            `json:"assigned_partner_id"`
	Partner         *DeliveryPartner `gorm:"foreignKey:AssignedPartner" json:"partner"`

	EstimatedTime time.Time `json:"estimated_time"`
	DeliveredAt   time.Time `json:"delivered_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
